package todo

import (
	"context"
	"todo/models"
)

type UseCase interface {
	Create(ctx context.Context, title, description string, user *models.User) error
	Delete(ctx context.Context, id string, user *models.User) error
	Get(ctx context.Context, user *models.User) ([]*models.Task, error)
	ChangeStatus(ctx context.Context, id string, status bool, user *models.User) error
}
