// vim: set ft=go:

package app

import (
	"example.com/namespace/todo-app/go-services/internal/todo-backend/application"
	"example.com/namespace/todo-app/go-services/internal/todo-backend/generated/core"
	"example.com/namespace/todo-app/go-services/internal/todo-backend/storage"
)

type ServiceExtra struct {
	TodoService *application.TodoService
}

func NewServiceExtra(ctx *core.MifyServiceContext) (*ServiceExtra, error) {
	memStorage := storage.NewTodoMemoryStorage()
	todoService := application.NewTodoService(memStorage)
	extra := &ServiceExtra{
		TodoService: todoService,
	}
	return extra, nil
}
