package repository

import (
	"context"
	"errors"

	"github.com/danielgtaylor/huma/v2"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/t-okuji/go-huma-sqlc/db/sqlc"
)

type IAuthorRepository interface {
	GetAuthor(ctx context.Context, id int64) (sqlc.Author, error)
	ListAuthors(ctx context.Context) ([]sqlc.Author, error)
	CreateAuthor(ctx context.Context, input sqlc.CreateAuthorParams) (sqlc.Author, error)
	UpdateAuthor(ctx context.Context, input sqlc.UpdateAuthorParams) (sqlc.Author, error)
	DeleteAuthor(ctx context.Context, id int64) error
}

type authorRepository struct {
	pool    *pgxpool.Pool
	queries *sqlc.Queries
}

func NewAuthorRepository(db *pgxpool.Pool) IAuthorRepository {
	return &authorRepository{pool: db, queries: sqlc.New(db)}
}

func (ar *authorRepository) GetAuthor(ctx context.Context, id int64) (sqlc.Author, error) {
	result, err := ar.queries.GetAuthor(ctx, id)

	switch {
	case err == pgx.ErrNoRows:
		return sqlc.Author{}, huma.Error404NotFound("", err)
	case err != nil:
		var pgErr *pgconn.ConnectError
		if errors.As(err, &pgErr) {
			return sqlc.Author{}, errors.New("failed to connect to database")
		}
		return sqlc.Author{}, err
	default:
		return result, nil
	}
}

func (ar *authorRepository) ListAuthors(ctx context.Context) ([]sqlc.Author, error) {
	result, err := ar.queries.ListAuthors(ctx)

	if err != nil {
		var pgErr *pgconn.ConnectError
		if errors.As(err, &pgErr) {
			return nil, errors.New("failed to connect to database")
		}
		return nil, err
	}
	return result, nil
}

func (ar *authorRepository) CreateAuthor(ctx context.Context, input sqlc.CreateAuthorParams) (sqlc.Author, error) {
	tx, err := ar.pool.Begin(ctx)
	if err != nil {
		var pgErr *pgconn.ConnectError
		if errors.As(err, &pgErr) {
			return sqlc.Author{}, errors.New("failed to connect to database")
		}
		return sqlc.Author{}, err
	}
	defer tx.Rollback(ctx)
	result, err := ar.queries.WithTx(tx).CreateAuthor(ctx, input)
	if err != nil {
		return sqlc.Author{}, err
	}
	tx.Commit(ctx)
	return result, nil
}

func (ar *authorRepository) UpdateAuthor(ctx context.Context, input sqlc.UpdateAuthorParams) (sqlc.Author, error) {
	tx, err := ar.pool.Begin(ctx)

	if err != nil {
		var pgErr *pgconn.ConnectError
		if errors.As(err, &pgErr) {
			return sqlc.Author{}, errors.New("failed to connect to database")
		}
		return sqlc.Author{}, err
	}
	defer tx.Rollback(ctx)

	result, err := ar.queries.WithTx(tx).UpdateAuthor(ctx, input)
	switch {
	case err == pgx.ErrNoRows:
		return sqlc.Author{}, huma.Error404NotFound("", err)
	case err != nil:
		return sqlc.Author{}, err
	default:
		tx.Commit(ctx)
		return result, nil
	}
}

func (ar *authorRepository) DeleteAuthor(ctx context.Context, id int64) error {
	err := ar.queries.DeleteAuthor(ctx, id)
	if err != nil {
		var pgErr *pgconn.ConnectError
		if errors.As(err, &pgErr) {
			return errors.New("failed to connect to database")
		}
		return err
	}
	return nil
}
