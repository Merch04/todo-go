package localcache

import (
	"context"
	"sync"
	"todo/models"
	"todo/todo"
)

type TaskLocalStorage struct {
	tasks map[string]*models.Task
	mutex *sync.Mutex
}

func NewTaskLocalStorage() *TaskLocalStorage {
	return &TaskLocalStorage{
		tasks: make(map[string]*models.Task),
		mutex: new(sync.Mutex),
	}
}

func (s *TaskLocalStorage) CreateTask(_ context.Context, task *models.Task) error {
	s.mutex.Lock()
	s.tasks[task.ID] = task
	s.mutex.Unlock()
	return nil
}

func (s *TaskLocalStorage) GetTasks(_ context.Context, user *models.User) ([]*models.Task, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	out := make([]*models.Task, 0)
	for _, task := range s.tasks {
		if task.UserID == user.ID {
			out = append(out, task)
		}
	}
	return out, nil
}

func (s *TaskLocalStorage) DeleteTask(_ context.Context, id string, user *models.User) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	task, exists := s.tasks[id]

	if exists && task.UserID == user.ID {
		delete(s.tasks, id)
		return nil
	}

	return todo.ErrTaskNotFound
}
func (s *TaskLocalStorage) ChangeStatus(_ context.Context, id string, isComplete bool, user *models.User) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	task, exists := s.tasks[id]

	if exists && task.UserID == user.ID {
		task.IsComplete = isComplete
		return nil
	}
	return todo.ErrTaskNotFound

}
