package db

import (
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/danielgtaylor/huma/v2"
	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
)

func ConnectDB(ctx context.Context) (*pgx.Conn, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_PORT"),
		os.Getenv("POSTGRES_DATABASE"),
	)
	conn, err := pgx.Connect(ctx, dsn)
	if err != nil {
		log.Err(err).Msg("")
		return nil, huma.Error500InternalServerError("", errors.New("failed to connect to DB"))
	}
	fmt.Println("Connected")
	return conn, nil
}

func CloseDB(ctx context.Context, conn *pgx.Conn) {
	if err := conn.Close(ctx); err != nil {
		log.Err(err).Msg("")
	}
	fmt.Println("Closed")
}
