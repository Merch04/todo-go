package localcache

import (
	"context"
	"sync"
	"todo/models"
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

func (s *TaskLocalStorage) CreateTask(ctx context.Context, task *models.Task) error {
	s.mutex.Lock()
	s.tasks[task.ID] = task
	s.mutex.Unlock()
	return nil
}

func (s *TaskLocalStorage) GetTasks(ctx context.Context, user *models.User) ([]*models.Task, error) {
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

func (s *TaskLocalStorage) DeleteTask(ctx context.Context, id string) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	delete(s.tasks, id)
	return nil
}
func (s *TaskLocalStorage) ChangeStatus(ctx context.Context, id string, isComplete bool) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	task, _ := s.tasks[id]
	task.IsComplete = isComplete
	return nil
}
