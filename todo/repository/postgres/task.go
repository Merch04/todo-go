package postgres

import (
	"context"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
	"todo/models"
)

type Task struct {
	ID          uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primarykey"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Title       string    `gorm:"not null"`
	Description string    `gorm:"not null"`
	IsComplete  bool      `gorm:"default:false"`
	UserID      uuid.UUID `gorm:"not null"`
}

type TaskRepository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) *TaskRepository {
	return &TaskRepository{db: db}
}

func (r TaskRepository) GetTasks(ctx context.Context, user *models.User) ([]*models.Task, error) {
	userId := user.ID

	out := make([]*Task, 0)
	err := r.db.Find(&out, "userID = ?", userId).Error
	if err != nil {
		return nil, err
	}
	return toTasks(out), nil
}

func (r TaskRepository) CreateTask(ctx context.Context, task *models.Task) error {
	model := toPostgresUser(task)
	err := r.db.Create(&model).Error
	if err != nil {
		return err
	}

	task.ID = model.ID.String()
	return nil
}

func (r TaskRepository) DeleteTask(ctx context.Context, id string) error {
	task := new(Task)
	if err := r.db.Delete(&task, id).Error; err != nil {
		return err
	}
	return nil
}

func (r TaskRepository) ChangeStatus(ctx context.Context, id string, isComplete bool) error {
	if err := r.db.Model(&Task{}).Where("id = ?", id).Update("is_complete", isComplete).Error; err != nil {
		return err
	}
	return nil
}

func toPostgresUser(t *models.Task) *Task {
	uuidStrUserId, _ := uuid.Parse(t.UserID)

	return &Task{
		Title:       t.Title,
		Description: t.Description,
		IsComplete:  t.IsComplete,
		UserID:      uuidStrUserId,
	}
}

func toTask(t *Task) *models.Task {
	return &models.Task{
		ID:          t.ID.String(),
		Title:       t.Title,
		Description: t.Description,
		IsComplete:  t.IsComplete,
	}
}
func toTasks(ts []*Task) []*models.Task {
	out := make([]*models.Task, len(ts))
	for i, t := range ts {
		out[i] = toTask(t)
	}

	return out
}
