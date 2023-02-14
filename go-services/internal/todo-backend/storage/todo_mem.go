package storage

import (
	"errors"
	"time"

	"example.com/namespace/todo-app/go-services/internal/todo-backend/domain"
	"example.com/namespace/todo-app/go-services/internal/todo-backend/generated/core"
)

var (
	ErrTodoNoteNotFound = errors.New("note not found")
)

type TodoMemoryStorage struct {
	todos  map[int64]domain.TodoNote
	nextID int64
}

func NewTodoMemoryStorage() *TodoMemoryStorage {
	return &TodoMemoryStorage{
		todos:  make(map[int64]domain.TodoNote),
		nextID: 1, // start from one to emulate db sequence
	}
}

func (s *TodoMemoryStorage) InsertTodoNote(
	ctx *core.MifyRequestContext, todoNote domain.TodoNote) (domain.TodoNote, error) {
	todoNote.ID = s.nextID
	todoNote.CreatedAt = time.Now()
	todoNote.UpdatedAt = time.Now()

	s.todos[s.nextID] = todoNote
	s.nextID += 1
	return todoNote, nil
}

func (s *TodoMemoryStorage) SelectTodoNote(ctx *core.MifyRequestContext, id int64) (domain.TodoNote, error) {
	note, ok := s.todos[id]
	if !ok {
		return domain.TodoNote{}, ErrTodoNoteNotFound
	}
	return note, nil
}

func (s *TodoMemoryStorage) UpdateTodoNote(
	ctx *core.MifyRequestContext, todoNote domain.TodoNote) (domain.TodoNote, error) {
	note, ok := s.todos[todoNote.ID]
	if !ok {
		return domain.TodoNote{}, ErrTodoNoteNotFound
	}
	note.Title = todoNote.Title
	note.Description = todoNote.Description
	note.UpdatedAt = time.Now()

	s.todos[todoNote.ID] = note
	return note, nil
}

func (s *TodoMemoryStorage) DeleteTodoNote(ctx *core.MifyRequestContext, id int64) error {
	delete(s.todos, id)
	return nil
}
