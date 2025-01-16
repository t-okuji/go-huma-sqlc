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

func NewDB(ctx context.Context) (*pgxpool.Pool, error) {
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
	conn, err := pgxpool.New(ctx, dsn)
	if err != nil {
		log.Err(err).Msg("")
		return nil, huma.Error500InternalServerError("", errors.New("failed to connect to DB"))
	}
	fmt.Println("Connected")
	return conn, nil
}

func CloseDB(ctx context.Context, conn *pgxpool.Pool) {
	conn.Close()
}
