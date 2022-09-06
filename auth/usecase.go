package auth

import (
	"context"

	"github.com/khuchuz/go-clean-architecture/models"
)

const CtxUserKey = "user"

type UseCase interface {
	SignUp(ctx context.Context, username, email, password string) error
	SignIn(ctx context.Context, username, password string) (string, error)
	ForgotPassword(ctx context.Context, email string) (string, error)
	ResetPassword(ctx context.Context, username, password string) (string, error)
	ParseToken(ctx context.Context, accessToken string) (*models.User, error)
}
