package router

import (
	"net/http"

	"github.com/danielgtaylor/huma/v2"
	"github.com/t-okuji/go-huma-sqlc/internal/controller"
)

func NewAuthorRouter(api huma.API, ac controller.IAuthorController) {
	huma.Register(api, huma.Operation{
		OperationID: "get-author",
		Method:      http.MethodGet,
		Path:        "/author/{id}",
		Summary:     "Get an author",
		Description: "Get an author by id.",
		Tags:        []string{"Authors"},
	}, ac.GetAuthor)

	huma.Register(api, huma.Operation{
		OperationID: "get-authors",
		Method:      http.MethodGet,
		Path:        "/authors",
		Summary:     "Get authors",
		Description: "Get authors.",
		Tags:        []string{"Authors"},
	}, ac.ListAuthors)

	huma.Register(api, huma.Operation{
		OperationID:   "post-author",
		Method:        http.MethodPost,
		Path:          "/author",
		Summary:       "Post a author",
		Tags:          []string{"Authors"},
		DefaultStatus: http.StatusCreated,
	}, ac.CreateAuthor)

	huma.Register(api, huma.Operation{
		OperationID:   "put-author",
		Method:        http.MethodPut,
		Path:          "/author",
		Summary:       "Update a author",
		Tags:          []string{"Authors"},
		DefaultStatus: http.StatusOK,
	}, ac.UpdateAuthor)

	huma.Register(api, huma.Operation{
		OperationID:   "delete-author",
		Method:        http.MethodDelete,
		Path:          "/author/{id}",
		Summary:       "Delete a author",
		Tags:          []string{"Authors"},
		DefaultStatus: http.StatusNoContent,
	}, ac.DeleteAuthor)
}
