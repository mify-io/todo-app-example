package storage

import (
	"errors"

	"example.com/namespace/todo-app/go-services/internal/todo-backend/domain"
	"example.com/namespace/todo-app/go-services/internal/todo-backend/generated/core"
	"example.com/namespace/todo-app/go-services/internal/todo-backend/generated/postgres"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

func makeDomainTodoNode(res postgres.Todo) domain.TodoNote {
	return domain.TodoNote{
		ID: res.ID,
		Title: res.Title,
		Description: res.Description,
		IsCompleted: res.IsCompleted,
		CreatedAt: res.CreatedAt,
		UpdatedAt: res.UpdatedAt,
	}
}

type TodoDBStorage struct {
	pool *pgxpool.Pool
	querier *postgres.Queries
}

func NewTodoDBStorage(ctx *core.MifyServiceContext) *TodoDBStorage {
	return &TodoDBStorage{
		pool: ctx.Postgres(),
		querier: postgres.New(ctx.Postgres()),
	}
}

func (s *TodoDBStorage) InsertTodoNote(
	ctx *core.MifyRequestContext, todoNote domain.TodoNote) (domain.TodoNote, error) {
	// Create transaction
	tx, err := s.pool.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return domain.TodoNote{}, err
	}
	// Don't forget to rollback in case of error
	defer tx.Rollback(ctx)
	res, err := s.querier.WithTx(tx).InsertTodoNote(ctx, postgres.InsertTodoNoteParams{
		Title: todoNote.Title,
		Description: todoNote.Description,
	})
	if err != nil {
		return domain.TodoNote{}, err
	}
	if err := tx.Commit(ctx); err != nil {
		return domain.TodoNote{}, err
	}
	return makeDomainTodoNode(res), nil
}

func (s *TodoDBStorage) SelectTodoNote(ctx *core.MifyRequestContext, id int64) (domain.TodoNote, error) {
	res, err := s.querier.SelectTodoNote(ctx, id)
	// Handle that no rows can be returned, so we can correctly return 404
	// to user instead of obscure 500 error.
	if err != nil && errors.Is(err, pgx.ErrNoRows) {
		return domain.TodoNote{}, ErrTodoNoteNotFound
	}
	if err != nil {
		return domain.TodoNote{}, err
	}
	return makeDomainTodoNode(res), nil
}

func (s *TodoDBStorage) SelectTodoNotes(ctx *core.MifyRequestContext) ([]domain.TodoNote, error) {
	res, err := s.querier.SelectTodoNotes(ctx)
	if err != nil && errors.Is(err, pgx.ErrNoRows) {
		return []domain.TodoNote{}, nil
	}
	if err != nil {
		return []domain.TodoNote{}, err
	}
	outList := make([]domain.TodoNote, 0, len(res))
	for _, dbNote := range res {
		outList = append(outList, makeDomainTodoNode(dbNote))
	}
	return outList, nil
}

func (s *TodoDBStorage) UpdateTodoNote(
	ctx *core.MifyRequestContext, todoNote domain.TodoNote) (domain.TodoNote, error) {
	tx, err := s.pool.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return domain.TodoNote{}, err
	}
	defer tx.Rollback(ctx)
	res, err := s.querier.WithTx(tx).UpdateTodoNote(ctx, postgres.UpdateTodoNoteParams{
		ID: todoNote.ID,
		Title: todoNote.Title,
		Description: todoNote.Description,
		IsCompleted: todoNote.IsCompleted,
		UpdatedAt: todoNote.UpdatedAt,
	})
	if err != nil {
		return domain.TodoNote{}, err
	}
	if err := tx.Commit(ctx); err != nil {
		return domain.TodoNote{}, err
	}
	return makeDomainTodoNode(res), nil
}

func (s *TodoDBStorage) DeleteTodoNote(ctx *core.MifyRequestContext, id int64) error {
	tx, err := s.pool.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)
	if err := s.querier.WithTx(tx).DeleteTodoNote(ctx, id); err != nil {
		return err
	}
	return tx.Commit(ctx)
}
