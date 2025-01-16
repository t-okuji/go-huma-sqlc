package controller

import (
	"context"

	"github.com/t-okuji/learn-huma/db/sqlc"
	"github.com/t-okuji/learn-huma/usecase"
)

type AuthorsOutput struct {
	Body struct {
		Authors []sqlc.Author `json:"authors"`
	}
}

type IAuthorController interface {
	ListAuthors(ctx context.Context, _ *struct {
	}) (*AuthorsOutput, error)
}

type authorController struct {
	as usecase.IAuthorUsecase
}

func NewAuthorController(as usecase.IAuthorUsecase) IAuthorController {
	return &authorController{as}
}

func (ah *authorController) ListAuthors(ctx context.Context, _ *struct {
}) (*AuthorsOutput, error) {
	result, err := ah.as.ListAuthors(ctx)
	if err != nil {
		return nil, err
	}

	resp := &AuthorsOutput{}

	resp.Body.Authors = result
	return resp, nil
}
