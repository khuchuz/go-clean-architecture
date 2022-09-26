package mock

import (
	"context"

	"github.com/khuchuz/go-clean-architecture/auth/entities"
	"github.com/khuchuz/go-clean-architecture/models"
	"github.com/stretchr/testify/mock"
)

type AuthUseCaseMock struct {
	mock.Mock
}

func (m *AuthUseCaseMock) SignUp(ctx context.Context, inp entities.SignUpInput) error {
	args := m.Called(inp.Username, inp.Email, inp.Password)

	return args.Error(0)
}

func (m *AuthUseCaseMock) SignIn(ctx context.Context, inp entities.SignInput) (string, error) {
	args := m.Called(inp.Username, inp.Password)

	return args.Get(0).(string), args.Error(1)
}

func (m *AuthUseCaseMock) ChangePassword(ctx context.Context, inp entities.ChangePasswordInput) error {
	args := m.Called(inp.Username, inp.OldPassword, inp.Password)

	return args.Error(0)
}

func (m *AuthUseCaseMock) ParseToken(ctx context.Context, accessToken string) (*models.User, error) {
	args := m.Called(accessToken)

	return args.Get(0).(*models.User), args.Error(1)
}
