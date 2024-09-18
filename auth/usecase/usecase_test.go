package usecase

import (
	"context"
	"testing"
	"todo/auth/repository/mock"
	"todo/models"

	"github.com/stretchr/testify/assert"
)

func TestAuthFlow(t *testing.T) {
	repo := new(mock.UserStorageMock)

	uc := NewAuthUseCase(
		repo,
		"salt",
		[]byte("secret"),
		86400,
	)

	var (
		username = "user"
		password = "pass"

		user = &models.User{
			Username: username,
			Password: "11f5639f22525155cb0b43573ee4212838c78d87", // sha1 of pass+salt
		}
	)

	// Sign Up
	repo.On("CreateUser", user).Return(nil)
	err := uc.SignUp(context.Background(), username, password)
	assert.NoError(t, err)

	// Sign In (Get Auth Token)
	repo.On("GetUser", user.Username, user.Password).Return(user, nil)
	token, err := uc.SignIn(context.Background(), username, password)
	assert.NoError(t, err)
	assert.NotEmpty(t, token)

	// Verify token
	parsedUser, err := uc.ParseToken(context.Background(), token)
	assert.NoError(t, err)
	assert.Equal(t, user, parsedUser)
}
