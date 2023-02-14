package domain

import (
	"time"

	"example.com/namespace/todo-app/go-services/internal/todo-backend/generated/core"
)

type TodoNote struct {
	ID          int64
	Title       string
	Description string
	IsCompleted bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

// Storage interface with it we can choose different implementations and write
// unit tests.
type TodoStorage interface {
	InsertTodoNote(ctx *core.MifyRequestContext, todoNote TodoNote) (TodoNote, error)
	UpdateTodoNote(ctx *core.MifyRequestContext, todoNote TodoNote) (TodoNote, error)
	SelectTodoNote(ctx *core.MifyRequestContext, id int64) (TodoNote, error)
	DeleteTodoNote(ctx *core.MifyRequestContext, id int64) error
}
