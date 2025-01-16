package main

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/adapters/humago"
	"github.com/danielgtaylor/huma/v2/humacli"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/t-okuji/learn-huma/controller"
	"github.com/t-okuji/learn-huma/db"
	"github.com/t-okuji/learn-huma/repository"
	"github.com/t-okuji/learn-huma/router"
	"github.com/t-okuji/learn-huma/usecase"

	_ "github.com/danielgtaylor/huma/v2/formats/cbor"
)

// Options for the CLI.
type Options struct {
	Port int `help:"Port to listen on" short:"p" default:"8888"`
}

func LogMiddleware(ctx huma.Context, next func(huma.Context)) {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	log.Info().Dict("api", zerolog.Dict().
		Str("method", ctx.Operation().Method).
		Str("path", ctx.Operation().Path),
	).Msg("")
	next(ctx)
}

func main() {
	ctx := context.Background()
	conn, err := db.NewDB(ctx)
	if err != nil {
		log.Err(err).Msg("")
	}
	defer db.CloseDB(ctx, conn)

	authorRepository := repository.NewAuthorRepository(conn)
	authorUsecase := usecase.NewAuthorUsecase(authorRepository)
	authorController := controller.NewAuthorController(authorUsecase)

	// Create a CLI app which takes a port option.
	cli := humacli.New(func(hooks humacli.Hooks, options *Options) {
		// Create a new router & API
		api_router := http.NewServeMux()
		api := humago.New(api_router, huma.DefaultConfig("My API", "1.0.0"))
		api.UseMiddleware(LogMiddleware)

		router.NewAuthorRouter(api, authorController)

		server := http.Server{
			Addr:    fmt.Sprintf(":%d", options.Port),
			Handler: api_router,
		}

		// Tell the CLI how to start your server.
		hooks.OnStart(func() {
			fmt.Printf("Starting server on port %d...\n", options.Port)
			// Start the server!
			server.ListenAndServe()
		})
	})

	// Run the CLI. When passed no commands, it starts the server.
	cli.Run()
}
