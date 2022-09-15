package usecase

import (
	"context"
	"testing"

	"github.com/khuchuz/go-clean-architecture/auth/entities"
	"github.com/khuchuz/go-clean-architecture/auth/repository/mock"
	"github.com/khuchuz/go-clean-architecture/models"
	"github.com/stretchr/testify/assert"
)

func TestAuthFlow(t *testing.T) {
	repo := new(mock.UserStorageMock)

	uc := NewAuthUseCase(repo, "salt", []byte("secret"), 86400)

	var (
		username    = "usermock"
		email       = "usermock@gmail.com"
		oldpassword = "passold"
		password    = "pass"

		ctx = context.Background()

		user = &models.User{
			Username: username,
			Email:    email,
			Password: "11f5639f22525155cb0b43573ee4212838c78d87", // sha1 of pass+salt
		}
	)

	// Sign Up
	repo.On("CreateUser", user).Return(nil)
	err := uc.SignUp(ctx, entities.SignUpInput{Username: username, Email: email, Password: password})
	assert.NoError(t, err)

	// Sign In (Get Auth Token)
	repo.On("GetUser", user.Username, user.Password).Return(user, nil)
	token, err := uc.SignIn(ctx, entities.SignInput{Username: username, Password: password})
	assert.NoError(t, err)
	assert.NotEmpty(t, token)

	// Change Password
	repo.On("UpdatePassword", user.Username, user.Password).Return(user, nil)
	err = uc.ChangePassword(ctx, entities.ChangePasswordInput{Username: username, Password: password, OldPassword: oldpassword})
	assert.NoError(t, err)
	assert.NotEmpty(t, token)

	// Verify token
	parsedUser, err := uc.ParseToken(ctx, token)
	assert.NoError(t, err)
	assert.Equal(t, user, parsedUser)
}
