package task

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

var ErrNotFound = errors.New("task not found")

type Repository struct{ db *pgxpool.Pool }

func NewRepository(db *pgxpool.Pool) *Repository { return &Repository{db: db} }

func (r *Repository) List(ctx context.Context) ([]Task, error) {
	rows, err := r.db.Query(ctx, `SELECT id, title, description, status, created_at, updated_at FROM tasks ORDER BY created_at DESC`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	tasks := []Task{}
	for rows.Next() {
		var item Task
		if err := rows.Scan(&item.ID, &item.Title, &item.Description, &item.Status, &item.CreatedAt, &item.UpdatedAt); err != nil {
			return nil, err
		}
		tasks = append(tasks, item)
	}
	return tasks, rows.Err()
}

func (r *Repository) Create(ctx context.Context, input Input) (Task, error) {
	return r.scanTask(r.db.QueryRow(ctx, `INSERT INTO tasks (title, description, status) VALUES ($1, $2, $3) RETURNING id, title, description, status, created_at, updated_at`, input.Title, input.Description, input.Status))
}

func (r *Repository) Get(ctx context.Context, id int64) (Task, error) {
	item, err := r.scanTask(r.db.QueryRow(ctx, `SELECT id, title, description, status, created_at, updated_at FROM tasks WHERE id=$1`, id))
	if errors.Is(err, pgx.ErrNoRows) {
		return Task{}, ErrNotFound
	}
	return item, err
}

func (r *Repository) Update(ctx context.Context, id int64, input Input) (Task, error) {
	item, err := r.scanTask(r.db.QueryRow(ctx, `UPDATE tasks SET title=$1, description=$2, status=$3, updated_at=NOW() WHERE id=$4 RETURNING id, title, description, status, created_at, updated_at`, input.Title, input.Description, input.Status, id))
	if errors.Is(err, pgx.ErrNoRows) {
		return Task{}, ErrNotFound
	}
	return item, err
}

func (r *Repository) Delete(ctx context.Context, id int64) error {
	result, err := r.db.Exec(ctx, "DELETE FROM tasks WHERE id=$1", id)
	if err != nil {
		return err
	}
	if result.RowsAffected() == 0 {
		return ErrNotFound
	}
	return nil
}

type rowScanner interface{ Scan(...any) error }

func (r *Repository) scanTask(row rowScanner) (Task, error) {
	var item Task
	err := row.Scan(&item.ID, &item.Title, &item.Description, &item.Status, &item.CreatedAt, &item.UpdatedAt)
	return item, err
}
