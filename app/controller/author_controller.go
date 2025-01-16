package controller

import (
	"context"

	"github.com/t-okuji/learn-huma/db/sqlc"
	"github.com/t-okuji/learn-huma/usecase"
)

type GetAuthorInput struct {
	Id int64 `path:"id" required:"true" example:"1" doc:"Author id"`
}
type AuthorOutput struct {
	Body sqlc.Author `json:"author"`
}
type AuthorsOutput struct {
	Body struct {
		Authors []sqlc.Author `json:"authors"`
	}
}

type IAuthorController interface {
	GetAuthor(ctx context.Context, input *GetAuthorInput) (*AuthorOutput, error)
	ListAuthors(ctx context.Context, _ *struct {
	}) (*AuthorsOutput, error)
}

type authorController struct {
	as usecase.IAuthorUsecase
}

func NewAuthorController(as usecase.IAuthorUsecase) IAuthorController {
	return &authorController{as}
}

func (ah *authorController) GetAuthor(ctx context.Context, input *GetAuthorInput) (*AuthorOutput, error) {
	result, err := ah.as.GetAuthor(ctx, input.Id)
	if err != nil {
		return nil, err
	}

	resp := &AuthorOutput{}

	resp.Body = result
	return resp, nil
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
