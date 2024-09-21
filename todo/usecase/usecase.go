package usecase

import (
	"context"
	"todo/models"
	"todo/todo"
)

type TaskUseCase struct {
	taskRepo todo.TaskRepository
}

func NewTodoUseCase(taskRepo todo.TaskRepository) *TaskUseCase {
	return &TaskUseCase{
		taskRepo: taskRepo,
	}
}

func (a *TaskUseCase) Create(ctx context.Context, title, description string, user *models.User) error {
	task := &models.Task{
		Title:       title,
		Description: description,
		UserID:      user.ID,
	}
	return a.taskRepo.CreateTask(ctx, task)
}
func (a *TaskUseCase) Delete(ctx context.Context, id string, user *models.User) error {
	return a.taskRepo.DeleteTask(ctx, id, user)
}
func (a *TaskUseCase) Get(ctx context.Context, user *models.User) ([]*models.Task, error) {
	return a.taskRepo.GetTasks(ctx, user)
}
func (a *TaskUseCase) ChangeStatus(ctx context.Context, id string, status bool, user *models.User) error {
	return a.taskRepo.ChangeStatus(ctx, id, status, user)
}
