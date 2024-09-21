package mock

import (
	"context"
	"github.com/stretchr/testify/mock"
	"todo/models"
)

type TaskStorageMock struct {
	mock.Mock
}

func (s *TaskStorageMock) CreateTask(_ context.Context, task *models.Task) error {
	args := s.Called(task)

	return args.Error(0)
}
func (s *TaskStorageMock) GetTasks(_ context.Context, user *models.User) ([]*models.Task, error) {
	args := s.Called(user)

	return args.Get(0).([]*models.Task), args.Error(1)
}
func (s *TaskStorageMock) DeleteTask(_ context.Context, id string, user *models.User) error {
	args := s.Called(id, user)

	return args.Error(0)
}
func (s *TaskStorageMock) ChangeStatus(_ context.Context, id string, isComplete bool, user *models.User) error {
	args := s.Called(id, isComplete, user)

	return args.Error(0)
}
