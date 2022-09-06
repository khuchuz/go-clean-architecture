package auth

import (
	"context"

	"github.com/khuchuz/go-clean-architecture/models"
)

const CtxUserKey = "user"

type UseCase interface {
	SignUp(ctx context.Context, username, email, password string) error
	SignIn(ctx context.Context, username, password string) (string, error)
	ChangePassword(ctx context.Context, username, oldpassword, password string) error
	InitChangePassword(ctx context.Context, email string) error
	ParseToken(ctx context.Context, accessToken string) (*models.User, error)
}
