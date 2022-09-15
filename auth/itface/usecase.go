package itface

import (
	"context"

	"github.com/khuchuz/go-clean-architecture/auth/entities"
	"github.com/khuchuz/go-clean-architecture/models"
)

const CtxUserKey = "user"

type UseCase interface {
	SignUp(ctx context.Context, inp entities.SignUpInput) error
	SignIn(ctx context.Context, inp entities.SignInput) (string, error)
	ChangePassword(ctx context.Context, inp entities.ChangePasswordInput) error
	ParseToken(ctx context.Context, accessToken string) (*models.User, error)
}
