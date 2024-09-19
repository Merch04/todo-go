package todo

import (
	"context"
	"todo/models"
)

type UseCase interface {
	Get(ctx context.Context, user *models.User) ([]*models.Task, error)
	Create(ctx context.Context, user *models.User) error
	Delete(ctx context.Context, user *models.User) error
	ChangeStatus(ctx context.Context, user *models.User) error
}
