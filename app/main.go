package main

import (
	"context"
	"fmt"
	"net/http"

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

// GreetingOutput represents the greeting operation response.
type GreetingOutput struct {
	Body struct {
		Message string `json:"message" example:"Hello, world!" doc:"Greeting message"`
	}
}

// ReviewInput represents the review operation request.
type ReviewInput struct {
	Body struct {
		Author  string `json:"author" maxLength:"10" doc:"Author of the review"`
		Rating  int    `json:"rating" minimum:"1" maximum:"5" doc:"Rating from 1 to 5"`
		Message string `json:"message,omitempty" maxLength:"100" doc:"Review message"`
	}
}

func sampleRoutes(api huma.API) {
	// Register GET /greeting/{name} handler.
	huma.Register(api, huma.Operation{
		OperationID: "get-greeting",
		Method:      http.MethodGet,
		Path:        "/greeting/{name}",
		Summary:     "Get a greeting",
		Description: "Get a greeting for a person by name.",
		Tags:        []string{"Greetings"},
	}, func(ctx context.Context, input *struct {
		Name string `path:"name" maxLength:"30" example:"world" doc:"Name to greet"`
	}) (*GreetingOutput, error) {
		resp := &GreetingOutput{}
		resp.Body.Message = fmt.Sprintf("Hello, %s!", input.Name)
		return resp, nil
	})

	// Register POST /reviews
	huma.Register(api, huma.Operation{
		OperationID:   "post-review",
		Method:        http.MethodPost,
		Path:          "/reviews",
		Summary:       "Post a review",
		Tags:          []string{"Reviews"},
		DefaultStatus: http.StatusCreated,
	}, func(ctx context.Context, i *ReviewInput) (*struct{}, error) {
		// TODO: save review in data store.
		return nil, nil
	})
}

func LogMiddleware(ctx huma.Context, next func(huma.Context)) {
	log.Info().Dict("api", zerolog.Dict().
		Str("method", ctx.Operation().Method).
		Str("path", ctx.Operation().Path),
	).Msg("")
	next(ctx)
}

func main() {
	ctx := context.Background()
	conn, err := db.ConnectDB(ctx)
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

		sampleRoutes(api)

		router.NewAuthorRouter(api, authorController)

		// Tell the CLI how to start your server.
		hooks.OnStart(func() {
			fmt.Printf("Starting server on port %d...\n", options.Port)
			// Start the server!
			http.ListenAndServe(fmt.Sprintf(":%d", options.Port), api_router)
		})
	})

	// Run the CLI. When passed no commands, it starts the server.
	cli.Run()
}
