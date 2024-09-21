package usecase

import (
	"context"
	"github.com/stretchr/testify/mock"
	"todo/models"
)

type TodoUseCaseMock struct {
	mock.Mock
}

func (m *TodoUseCaseMock) Create(_ context.Context, title, description string, user *models.User) error {
	args := m.Called(title, description, user)

	return args.Error(0)
}
func (m *TodoUseCaseMock) Delete(_ context.Context, id string, user *models.User) error {
	args := m.Called(id, user)

	return args.Error(0)
}
func (m *TodoUseCaseMock) Get(_ context.Context, user *models.User) ([]*models.Task, error) {
	args := m.Called(user)

	return args.Get(0).([]*models.Task), args.Error(1)
}

func (m *TodoUseCaseMock) ChangeStatus(_ context.Context, id string, status bool, user *models.User) error {
	args := m.Called(id, status, user)

	return args.Error(0)
}
