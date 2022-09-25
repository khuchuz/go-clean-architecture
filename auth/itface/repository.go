package itface

import (
	"context"

	"github.com/khuchuz/go-clean-architecture/models"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user *models.User) error
	GetUser(ctx context.Context, username, password string) (*models.User, error)
	IsUserExistByUsername(ctx context.Context, username string) bool
	IsUserExistByEmail(ctx context.Context, email string) bool
	UpdatePassword(ctx context.Context, username, password string) error
}
