package postgres

import (
	"context"
	"gorm.io/gorm"
	"strconv"
	"todo/models"
)

type User struct {
	gorm.Model
	Username string `gorm:"unique;not null"`
	Password string `gorm:"not null"`
}

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (r UserRepository) CreateUser(ctx context.Context, user *models.User) error {
	model := toPostgresUser(user)
	err := r.db.Create(&model).Error
	if err != nil {
		return err
	}

	user.ID = strconv.Itoa(int(model.ID))
	return nil
}

func (r UserRepository) GetUser(ctx context.Context, username, password string) (*models.User, error) {
	user := new(User)
	err := r.db.First(&user, "username = ? AND password = ?", username, password).Error

	if err != nil {
		return nil, err
	}

	return toModel(user), nil
}

func toPostgresUser(u *models.User) *User {
	return &User{
		Username: u.Username,
		Password: u.Password,
	}
}

func toModel(u *User) *models.User {

	return &models.User{
		ID:       strconv.Itoa(int(u.ID)),
		Username: u.Username,
		Password: u.Password,
	}
}
