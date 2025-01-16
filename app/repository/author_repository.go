package repository

import (
	"context"

	"github.com/danielgtaylor/huma/v2"
	"github.com/jackc/pgx/v5"
	"github.com/t-okuji/learn-huma/db/sqlc"
)

type IAuthorRepository interface {
	GetAuthor(ctx context.Context, id int64) (sqlc.Author, error)
	ListAuthors(ctx context.Context) ([]sqlc.Author, error)
	CreateAuthor(ctx context.Context, input sqlc.CreateAuthorParams) (sqlc.Author, error)
	UpdateAuthor(ctx context.Context, input sqlc.UpdateAuthorParams) (sqlc.Author, error)
	DeleteAuthor(ctx context.Context, id int64) error
}

type authorRepository struct {
	queries *sqlc.Queries
}

func NewAuthorRepository(db *pgx.Conn) IAuthorRepository {
	return &authorRepository{queries: sqlc.New(db)}
}

func (ar *authorRepository) GetAuthor(ctx context.Context, id int64) (sqlc.Author, error) {
	result, err := ar.queries.GetAuthor(ctx, id)

	switch {
	case err == pgx.ErrNoRows:
		return sqlc.Author{}, huma.Error404NotFound("", err)
	case err != nil:
		return sqlc.Author{}, err
	default:
		return result, nil
	}
}

func (ar *authorRepository) ListAuthors(ctx context.Context) ([]sqlc.Author, error) {
	result, err := ar.queries.ListAuthors(ctx)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (ar *authorRepository) CreateAuthor(ctx context.Context, input sqlc.CreateAuthorParams) (sqlc.Author, error) {
	result, err := ar.queries.CreateAuthor(ctx, input)
	if err != nil {
		return sqlc.Author{}, err
	}
	return result, nil
}

func (ar *authorRepository) UpdateAuthor(ctx context.Context, input sqlc.UpdateAuthorParams) (sqlc.Author, error) {
	result, err := ar.queries.UpdateAuthor(ctx, input)
	switch {
	case err == pgx.ErrNoRows:
		return sqlc.Author{}, huma.Error404NotFound("", err)
	case err != nil:
		return sqlc.Author{}, err
	default:
		return result, nil
	}
}

func (ar *authorRepository) DeleteAuthor(ctx context.Context, id int64) error {
	err := ar.queries.DeleteAuthor(ctx, id)
	if err != nil {
		return err
	}
	return nil
}
