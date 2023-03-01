-- name: SelectTodoNote :one
SELECT * FROM todos
WHERE id = $1 LIMIT 1;

-- name: InsertTodoNote :one
INSERT INTO todos (
  title, description
) VALUES (
  $1, $2
)
RETURNING *;

-- name: UpdateTodoNote :one
UPDATE todos
SET title = $2, description = $3,
is_completed = $4, updated_at = $5
WHERE id = $1
RETURNING *;

-- name: DeleteTodoNote :exec
DELETE FROM todos WHERE id = $1;
