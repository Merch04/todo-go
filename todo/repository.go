package todo

import (
	"context"
	"todo/models"
)

type TaskRepository interface {
	CreateTask(ctx context.Context, task *models.Task) error
	GetTasks(ctx context.Context, user *models.User) ([]*models.Task, error)
	DeleteTask(ctx context.Context, id string, user *models.User) error
	ChangeStatus(ctx context.Context, id string, isComplete bool, user *models.User) error
}
