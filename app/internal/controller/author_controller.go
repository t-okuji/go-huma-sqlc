package controller

import (
	"context"

	"github.com/t-okuji/go-huma-sqlc/db/sqlc"
	"github.com/t-okuji/go-huma-sqlc/internal/usecase"
)

type GetAuthorInput struct {
	Id int64 `path:"id" required:"true" example:"1" doc:"Author id"`
}
type CreateAuthorInput struct {
	Body struct {
		Name string  `json:"name" required:"true" doc:"Author name"`
		Bio  *string `json:"bio,omitempty" doc:"Author bio"`
	}
}
type UpdateAuthorInput struct {
	Body struct {
		Id   int64   `path:"id" required:"true" example:"1" doc:"Author id"`
		Name string  `json:"name" required:"true" doc:"Author name"`
		Bio  *string `json:"bio,omitempty" doc:"Author bio"`
	}
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
	CreateAuthor(ctx context.Context, input *CreateAuthorInput) (*AuthorOutput, error)
	UpdateAuthor(ctx context.Context, input *UpdateAuthorInput) (*AuthorOutput, error)
	DeleteAuthor(ctx context.Context, input *GetAuthorInput) (*struct{}, error)
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

func (ah *authorController) CreateAuthor(ctx context.Context, input *CreateAuthorInput) (*AuthorOutput, error) {
	body := sqlc.CreateAuthorParams{Name: input.Body.Name, Bio: input.Body.Bio}
	result, err := ah.as.CreateAuthor(ctx, body)
	if err != nil {
		return nil, err
	}

	resp := &AuthorOutput{}

	resp.Body = result
	return resp, nil
}

func (ah *authorController) UpdateAuthor(ctx context.Context, input *UpdateAuthorInput) (*AuthorOutput, error) {
	body := sqlc.UpdateAuthorParams{ID: input.Body.Id, Name: input.Body.Name, Bio: input.Body.Bio}
	result, err := ah.as.UpdateAuthor(ctx, body)
	if err != nil {
		return nil, err
	}

	resp := &AuthorOutput{}

	resp.Body = result
	return resp, nil
}

func (ah *authorController) DeleteAuthor(ctx context.Context, input *GetAuthorInput) (*struct{}, error) {
	err := ah.as.DeleteAuthor(ctx, input.Id)
	if err != nil {
		return nil, err
	}

	return nil, nil
}
