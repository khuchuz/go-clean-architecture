package mock

import (
	"context"

	"github.com/khuchuz/go-clean-architecture/models"
	"github.com/stretchr/testify/mock"
)

type UserStorageMock struct {
	mock.Mock
}

func (s *UserStorageMock) CreateUser(ctx context.Context, user *models.User) error {
	args := s.Called(user)

	return args.Error(0)
}

func (s *UserStorageMock) GetUser(ctx context.Context, username, password string) (*models.User, error) {
	args := s.Called(username, password)

	return args.Get(0).(*models.User), args.Error(1)
}

func (s *UserStorageMock) UpdatePassword(ctx context.Context, username, password string) error {
	args := s.Called(username, password)

	return args.Error(0)
}

func (s *UserStorageMock) IsUserExistByUsername(ctx context.Context, username string) bool {
	args := s.Called(username)

	return args.Bool(0)
}

func (s *UserStorageMock) IsUserExistByEmail(ctx context.Context, email string) bool {
	args := s.Called(email)

	return args.Bool(0)
}
