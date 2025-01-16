package repository

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/t-okuji/learn-huma/db/sqlc"
)

type IAuthorRepository interface {
	ListAuthors(ctx context.Context) ([]sqlc.Author, error)
}

type authorRepository struct {
	queries *sqlc.Queries
}

func NewAuthorRepository(db *pgx.Conn) IAuthorRepository {
	return &authorRepository{queries: sqlc.New(db)}
}

func (ar *authorRepository) ListAuthors(ctx context.Context) ([]sqlc.Author, error) {
	result, err := ar.queries.ListAuthors(ctx)
	if err != nil {
		return nil, err
	}
	return result, nil
}
