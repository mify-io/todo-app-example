// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.0
// source: queries.sql

package postgres

import (
	"context"
	"time"
)

const deleteTodoNote = `-- name: DeleteTodoNote :exec
DELETE FROM todos WHERE id = $1
`

func (q *Queries) DeleteTodoNote(ctx context.Context, id int64) error {
	_, err := q.db.Exec(ctx, deleteTodoNote, id)
	return err
}

const insertTodoNote = `-- name: InsertTodoNote :one
INSERT INTO todos (
  title, description
) VALUES (
  $1, $2
)
RETURNING id, title, description, is_completed, created_at, updated_at
`

type InsertTodoNoteParams struct {
	Title       string
	Description string
}

func (q *Queries) InsertTodoNote(ctx context.Context, arg InsertTodoNoteParams) (Todo, error) {
	row := q.db.QueryRow(ctx, insertTodoNote, arg.Title, arg.Description)
	var i Todo
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Description,
		&i.IsCompleted,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const selectTodoNote = `-- name: SelectTodoNote :one
SELECT id, title, description, is_completed, created_at, updated_at FROM todos
WHERE id = $1 LIMIT 1
`

func (q *Queries) SelectTodoNote(ctx context.Context, id int64) (Todo, error) {
	row := q.db.QueryRow(ctx, selectTodoNote, id)
	var i Todo
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Description,
		&i.IsCompleted,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const updateTodoNote = `-- name: UpdateTodoNote :one
UPDATE todos
SET title = $2, description = $3,
is_completed = $4, updated_at = $5
WHERE id = $1
RETURNING id, title, description, is_completed, created_at, updated_at
`

type UpdateTodoNoteParams struct {
	ID          int64
	Title       string
	Description string
	IsCompleted bool
	UpdatedAt   time.Time
}

func (q *Queries) UpdateTodoNote(ctx context.Context, arg UpdateTodoNoteParams) (Todo, error) {
	row := q.db.QueryRow(ctx, updateTodoNote,
		arg.ID,
		arg.Title,
		arg.Description,
		arg.IsCompleted,
		arg.UpdatedAt,
	)
	var i Todo
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Description,
		&i.IsCompleted,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}
