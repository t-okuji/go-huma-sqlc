package usecase

import (
	"context"

	"github.com/t-okuji/learn-huma/db/sqlc"
	"github.com/t-okuji/learn-huma/repository"
)

type IAuthorUsecase interface {
	ListAuthors(ctx context.Context) ([]sqlc.Author, error)
}

type authorUsecase struct {
	ar repository.IAuthorRepository
}

func NewAuthorUsecase(ar repository.IAuthorRepository) IAuthorUsecase {
	return &authorUsecase{ar}
}

func (as *authorUsecase) ListAuthors(ctx context.Context) ([]sqlc.Author, error) {
	result, err := as.ar.ListAuthors(ctx)
	if err != nil {
		return nil, err
	}
	return result, nil
}
