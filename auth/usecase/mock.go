package usecase

import (
	"context"

	"github.com/khuchuz/go-clean-architecture/models"
	"github.com/stretchr/testify/mock"
)

type AuthUseCaseMock struct {
	mock.Mock
}

func (m *AuthUseCaseMock) SignUp(ctx context.Context, username, email, password string) error {
	args := m.Called(username, email, password)

	return args.Error(0)
}

func (m *AuthUseCaseMock) SignIn(ctx context.Context, username, password string) (string, error) {
	args := m.Called(username, password)

	return args.Get(0).(string), args.Error(1)
}

func (m *AuthUseCaseMock) ChangePassword(ctx context.Context, username, oldpassword, password string) error {
	args := m.Called(username, oldpassword, password)

	return args.Error(1)
}

func (m *AuthUseCaseMock) ParseToken(ctx context.Context, accessToken string) (*models.User, error) {
	args := m.Called(accessToken)

	return args.Get(0).(*models.User), args.Error(1)
}
