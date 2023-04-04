package application

import (
	"example.com/namespace/todo-app/go-services/internal/todo-backend/domain"
	"example.com/namespace/todo-app/go-services/internal/todo-backend/generated/core"
)

type TodoService struct {
	storage domain.TodoStorage
}

func NewTodoService(storage domain.TodoStorage) *TodoService {
	return &TodoService{
		storage: storage,
	}
}

func (s *TodoService) CreateTodo(ctx *core.MifyRequestContext, note domain.TodoNote) (domain.TodoNote, error) {
	return s.storage.InsertTodoNote(ctx, note)
}

func (s *TodoService) UpdateTodo(ctx *core.MifyRequestContext, note domain.TodoNote) (domain.TodoNote, error) {
	return s.storage.UpdateTodoNote(ctx, note)
}

func (s *TodoService) GetTodo(ctx *core.MifyRequestContext, id int64) (domain.TodoNote, error) {
	return s.storage.SelectTodoNote(ctx, id)
}

func (s *TodoService) ListTodos(ctx *core.MifyRequestContext) ([]domain.TodoNote, error) {
	return s.storage.SelectTodoNotes(ctx)
}

func (s *TodoService) DeleteTodo(ctx *core.MifyRequestContext, id int64) error {
	return s.storage.DeleteTodoNote(ctx, id)
}
