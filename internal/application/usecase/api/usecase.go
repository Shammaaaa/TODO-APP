package api

import (
	"context"
	"github.com/shamil/todo-app/internal/application/domain"
)

type TaskRepository interface {
	CreateTask(ctx context.Context, task domain.Task) (domain.Task, error)
	DeleteTask(ctx context.Context, id int) error
	UpdateTask(ctx context.Context, task domain.Task) (domain.Task, error)
	GetAllTasks(ctx context.Context) ([]domain.Task, error)
}

type UseCase struct {
	TaskRepository TaskRepository
}

func NewApiUseCase(taskRepository TaskRepository) *UseCase {
	return &UseCase{
		TaskRepository: taskRepository,
	}
}

func (u *UseCase) CreateTask(ctx context.Context, task domain.Task) (domain.Task, error) {
	return u.TaskRepository.CreateTask(ctx, task)
}
func (u *UseCase) GetAllTasks(ctx context.Context) ([]domain.Task, error) {
	return u.TaskRepository.GetAllTasks(ctx)
}

func (u *UseCase) UpdateTask(ctx context.Context, task domain.Task) (domain.Task, error) {
	return u.TaskRepository.UpdateTask(ctx, task)

}

func (u *UseCase) DeleteTask(ctx context.Context, id int) error {
	return u.TaskRepository.DeleteTask(ctx, id)
}
