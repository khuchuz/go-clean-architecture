package usecase

import (
	"context"
	"crypto/sha1"
	"fmt"
	"time"

	"github.com/khuchuz/go-clean-architecture/models"

	"github.com/dgrijalva/jwt-go/v4"
	"github.com/khuchuz/go-clean-architecture/auth"
	"github.com/khuchuz/go-clean-architecture/auth/entities"
	itface "github.com/khuchuz/go-clean-architecture/auth/itface"
)

type AuthClaims struct {
	jwt.StandardClaims
	User *models.User `json:"user"`
}

type AuthUseCase struct {
	userRepo       itface.UserRepository
	hashSalt       string
	signingKey     []byte
	expireDuration time.Duration
}

func NewAuthUseCase(
	userRepo itface.UserRepository,
	hashSalt string,
	signingKey []byte,
	tokenTTLSeconds time.Duration) *AuthUseCase {
	return &AuthUseCase{
		userRepo:       userRepo,
		hashSalt:       hashSalt,
		signingKey:     signingKey,
		expireDuration: time.Second * tokenTTLSeconds,
	}
}

func (a *AuthUseCase) SignUp(ctx context.Context, inp entities.SignUpInput) error {
	pwd := sha1.New()
	pwd.Write([]byte(inp.Password))
	pwd.Write([]byte(a.hashSalt))
	if inp.Username == "" || inp.Email == "" || inp.Password == "" {
		return auth.ErrDataTidakLengkap
	}

	//if cekUser, _ := a.userRepo.GetUserByUsername(ctx, inp.Username); cekUser != nil {
	//	return auth.ErrUserDuplicate
	//}

	//if cekUser, _ := a.userRepo.GetUserByEmail(ctx, inp.Email); cekUser != nil {
	//	return auth.ErrEmailDuplicate
	//}

	user := &models.User{
		Username: inp.Username,
		Email:    inp.Email,
		Password: fmt.Sprintf("%x", pwd.Sum(nil)),
	}

	return a.userRepo.CreateUser(ctx, user)
}

func (a *AuthUseCase) SignIn(ctx context.Context, inp entities.SignInput) (string, error) {
	pwd := sha1.New()
	pwd.Write([]byte(inp.Password))
	pwd.Write([]byte(a.hashSalt))
	password := fmt.Sprintf("%x", pwd.Sum(nil))

	user, err := a.userRepo.GetUser(ctx, inp.Username, password)
	if err != nil {
		return "", auth.ErrUserNotFound
	}

	claims := AuthClaims{
		User: user,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: jwt.At(time.Now().Add(a.expireDuration)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(a.signingKey)
}

func (a *AuthUseCase) ChangePassword(ctx context.Context, inp entities.ChangePasswordInput) error {
	if inp.Username == "" || inp.OldPassword == "" || inp.Password == "" {
		return auth.ErrDataTidakLengkap
	}
	if inp.OldPassword == inp.Password {
		return auth.ErrPasswordSame
	}
	pwd := sha1.New()
	pwd.Write([]byte(inp.OldPassword))
	pwd.Write([]byte(a.hashSalt))
	oldpassword := fmt.Sprintf("%x", pwd.Sum(nil))

	pwd2 := sha1.New()
	pwd2.Write([]byte(inp.Password))
	pwd2.Write([]byte(a.hashSalt))
	password := fmt.Sprintf("%x", pwd2.Sum(nil))

	_, err := a.userRepo.GetUser(ctx, inp.Username, oldpassword)
	if err != nil {
		return auth.ErrUserNotFound
	}
	return a.userRepo.UpdatePassword(ctx, inp.Username, password)
}

func (a *AuthUseCase) ParseToken(ctx context.Context, accessToken string) (*models.User, error) {
	token, err := jwt.ParseWithClaims(accessToken, &AuthClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return a.signingKey, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*AuthClaims); ok && token.Valid {
		return claims.User, nil
	}

	return nil, auth.ErrInvalidAccessToken
}
