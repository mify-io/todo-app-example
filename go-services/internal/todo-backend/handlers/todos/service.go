package api_todos

import (
	"net/http"
	"strconv"

	"example.com/namespace/todo-app/go-services/internal/todo-backend/domain"
	"example.com/namespace/todo-app/go-services/internal/todo-backend/generated/api"
	"example.com/namespace/todo-app/go-services/internal/todo-backend/generated/apputil"
	"example.com/namespace/todo-app/go-services/internal/todo-backend/generated/core"
	"example.com/namespace/todo-app/go-services/internal/todo-backend/handlers"
)

type TodosApiService struct{}

// NewTodosApiService creates a default api service
func NewTodosApiService() openapi.TodosApiServicer {
	return &TodosApiService{}
}

// TodosPost - Add new todo note
func (s *TodosApiService) TodosPost(
	ctx *core.MifyRequestContext, todoNoteCreateRequest openapi.TodoNoteCreateRequest) (openapi.ServiceResponse, error) {
	todoSvc := apputil.GetServiceExtra(ctx.ServiceContext()).TodoService

	note, err := todoSvc.CreateTodo(ctx, domain.TodoNote{
		Title:       todoNoteCreateRequest.Title,
		Description: todoNoteCreateRequest.Description,
	})
	if err != nil {
		// Pass some user friendly error to user instead of internal one,
		// which will be available in service logs. Since we don't expect
		// any errors here, we return Internal Server Error, which is catch-all
		// for all unknown bad things. These kind of errors will be reported in
		// metrics and it's easy to set alerts for them.
		return openapi.Response(http.StatusInternalServerError, openapi.Error{
			Code:    strconv.Itoa(http.StatusInternalServerError),
			Message: "Failed to create todo note",
		}), err
	}

	return openapi.Response(http.StatusOK, handlers.MakeAPITodoNote(note)), nil
}
