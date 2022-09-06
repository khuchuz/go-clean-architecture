package auth

import (
	"context"

	"github.com/khuchuz/go-clean-architecture/models"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user *models.User) error
	GetUser(ctx context.Context, username, password string) (*models.User, error)
	GetUserByUsername(ctx context.Context, username string) (*models.User, error)
}
