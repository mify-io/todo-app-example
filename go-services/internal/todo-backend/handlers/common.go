package handlers

import (
	"example.com/namespace/todo-app/go-services/internal/todo-backend/domain"
	openapi "example.com/namespace/todo-app/go-services/internal/todo-backend/generated/api"
)

func MakeAPITodoNote(note domain.TodoNote) openapi.TodoNote {
	return openapi.TodoNote{
		Id:          note.ID,
		Title:       note.Title,
		Description: note.Description,
		IsCompleted: note.IsCompleted,
		CreatedAt:   note.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt:   note.UpdatedAt.Format("2006-01-02 15:04:05"),
	}
}
