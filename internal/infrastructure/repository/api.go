package repository

import (
	"context"
	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5"
	"github.com/shamil/todo-app/internal/application/domain"
)

func (r *Repository) CreateTask(ctx context.Context, task domain.Task) (domain.Task, error) {
	query, args, err := r.sb.Insert("tasks").
		Columns("title", "description", "status").
		Values(task.Title, task.Description, task.Status).
		Suffix("RETURNING id, created_at, updated_at").
		ToSql()
	if err != nil {
		return task, err
	}

	err = r.db.QueryRow(ctx, query, args...).Scan(&task.ID, &task.CreatedAt, &task.UpdatedAt)
	return task, err
}

func (r *Repository) GetAllTasks(ctx context.Context) ([]domain.Task, error) {
	query, args, err := r.sb.Select("id", "title", "description", "status", "created_at", "updated_at").
		From("tasks").
		ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := r.db.Query(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []domain.Task
	for rows.Next() {
		var task domain.Task
		err := rows.Scan(&task.ID, &task.Title, &task.Description, &task.Status, &task.CreatedAt, &task.UpdatedAt)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}
	return tasks, nil
}

func (r *Repository) UpdateTask(ctx context.Context, task domain.Task) (domain.Task, error) {
	query, args, err := r.sb.Update("tasks").
		Set("title", task.Title).
		Set("description", task.Description).
		Set("status", task.Status).
		Set("updated_at", sq.Expr("now()")).
		Where(sq.Eq{"id": task.ID}).
		Suffix("RETURNING updated_at").
		ToSql()
	if err != nil {
		return task, err
	}

	err = r.db.QueryRow(ctx, query, args...).Scan(&task.UpdatedAt)
	return task, err
}

func (r *Repository) DeleteTask(ctx context.Context, id int) error {
	query, args, err := r.sb.Delete("tasks").
		Where(sq.Eq{"id": id}).
		ToSql()
	if err != nil {
		return err
	}

	res, err := r.db.Exec(ctx, query, args...)
	if err != nil {
		return err
	}
	if res.RowsAffected() == 0 {
		return pgx.ErrNoRows
	}
	return nil
}
