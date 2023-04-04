package api_todos_id

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"example.com/namespace/todo-app/go-services/internal/todo-backend/domain"
	"example.com/namespace/todo-app/go-services/internal/todo-backend/generated/api"
	"example.com/namespace/todo-app/go-services/internal/todo-backend/generated/apputil"
	"example.com/namespace/todo-app/go-services/internal/todo-backend/generated/core"
	"example.com/namespace/todo-app/go-services/internal/todo-backend/handlers"
	"example.com/namespace/todo-app/go-services/internal/todo-backend/storage"
)

type TodosIdApiService struct{}

// NewTodosIdApiService creates a default api service
func NewTodosIdApiService() openapi.TodosIdApiServicer {
	return &TodosIdApiService{}
}

// TodosIdDelete - Delete todo note by id
func (s *TodosIdApiService) TodosIdDelete(ctx *core.MifyRequestContext, id int64) (openapi.ServiceResponse, error) {
	todoSvc := apputil.GetServiceExtra(ctx.ServiceContext()).TodoService
	if err := todoSvc.DeleteTodo(ctx, id); err != nil {
		return openapi.Response(http.StatusInternalServerError, openapi.Error{
			Code:    strconv.Itoa(http.StatusInternalServerError),
			Message: "Failed to delete todo note",
		}), err
	}
	// Return empty response
	return openapi.Response(200, map[string]interface{}{}), nil
}

// TodosPut - Update todo note
func (s *TodosIdApiService) TodosIdPut(
	ctx *core.MifyRequestContext, id int64, todoNoteUpdateRequest openapi.TodoNoteUpdateRequest) (openapi.ServiceResponse, error) {
	todoSvc := apputil.GetServiceExtra(ctx.ServiceContext()).TodoService
	note, err := todoSvc.UpdateTodo(ctx, domain.TodoNote{
		ID:          id,
		Title:       todoNoteUpdateRequest.Title,
		Description: todoNoteUpdateRequest.Description,
		IsCompleted: todoNoteUpdateRequest.IsCompleted,
		UpdatedAt:   time.Now(),
	})
	if err != nil && errors.Is(err, storage.ErrTodoNoteNotFound) {
		// Return not found error
		return openapi.Response(http.StatusNotFound, openapi.Error{
			Code:    strconv.Itoa(http.StatusNotFound),
			Message: fmt.Sprintf("Unable to find todo note by id: %d", id),
		}), err
	}
	if err != nil {
		return openapi.Response(http.StatusInternalServerError, openapi.Error{
			Code:    strconv.Itoa(http.StatusInternalServerError),
			Message: "Failed to update todo note",
		}), err
	}

	return openapi.Response(http.StatusOK, handlers.MakeAPITodoNote(note)), nil
}

// TodosIdGet - Get todo note by id
func (s *TodosIdApiService) TodosIdGet(ctx *core.MifyRequestContext, id int64) (openapi.ServiceResponse, error) {
	todoSvc := apputil.GetServiceExtra(ctx.ServiceContext()).TodoService

	note, err := todoSvc.GetTodo(ctx, id)
	if err != nil && errors.Is(err, storage.ErrTodoNoteNotFound) {
		// Return not found error
		return openapi.Response(http.StatusNotFound, openapi.Error{
			Code:    strconv.Itoa(http.StatusNotFound),
			Message: fmt.Sprintf("Unable to find todo note by id: %d", id),
		}), err
	}
	if err != nil {
		return openapi.Response(http.StatusInternalServerError, openapi.Error{
			Code:    strconv.Itoa(http.StatusInternalServerError),
			Message: "Failed to get todo note",
		}), err
	}

	return openapi.Response(http.StatusOK, handlers.MakeAPITodoNote(note)), nil
}
