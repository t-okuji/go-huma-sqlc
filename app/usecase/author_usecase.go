package usecase

import (
	"context"

	"github.com/t-okuji/learn-huma/db/sqlc"
	"github.com/t-okuji/learn-huma/repository"
)

type IAuthorUsecase interface {
	GetAuthor(ctx context.Context, id int64) (sqlc.Author, error)
	ListAuthors(ctx context.Context) ([]sqlc.Author, error)
	CreateAuthor(ctx context.Context, input sqlc.CreateAuthorParams) (sqlc.Author, error)
	UpdateAuthor(ctx context.Context, input sqlc.UpdateAuthorParams) (sqlc.Author, error)
}

type authorUsecase struct {
	ar repository.IAuthorRepository
}

func NewAuthorUsecase(ar repository.IAuthorRepository) IAuthorUsecase {
	return &authorUsecase{ar}
}

func (as *authorUsecase) GetAuthor(ctx context.Context, id int64) (sqlc.Author, error) {
	result, err := as.ar.GetAuthor(ctx, id)
	if err != nil {
		return sqlc.Author{}, err
	}
	return result, nil
}

func (as *authorUsecase) ListAuthors(ctx context.Context) ([]sqlc.Author, error) {
	result, err := as.ar.ListAuthors(ctx)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (as *authorUsecase) CreateAuthor(ctx context.Context, input sqlc.CreateAuthorParams) (sqlc.Author, error) {
	result, err := as.ar.CreateAuthor(ctx, input)
	if err != nil {
		return sqlc.Author{}, err
	}
	return result, nil
}

func (as *authorUsecase) UpdateAuthor(ctx context.Context, input sqlc.UpdateAuthorParams) (sqlc.Author, error) {
	result, err := as.ar.UpdateAuthor(ctx, input)
	if err != nil {
		return sqlc.Author{}, err
	}
	return result, nil
}
