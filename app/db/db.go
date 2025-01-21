package db

import (
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/danielgtaylor/huma/v2"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
)

func NewPool(ctx context.Context) (*pgxpool.Pool, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_PORT"),
		os.Getenv("POSTGRES_DB"),
	)
	conn, err := pgxpool.New(ctx, dsn)
	if err != nil {
		log.Err(err).Msg("")
		return nil, huma.Error500InternalServerError("", errors.New("failed to connect to DB"))
	}
	if err := conn.Ping(ctx); err != nil {
		log.Err(err).Msg("")
	}
	fmt.Println("Connected")
	return conn, nil
}

func ClosePool(ctx context.Context, conn *pgxpool.Pool) {
	conn.Close()
}
