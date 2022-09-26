package usecase

import (
	"context"
	"testing"

	"github.com/khuchuz/go-clean-architecture/auth"
	"github.com/khuchuz/go-clean-architecture/auth/entities"
	"github.com/khuchuz/go-clean-architecture/auth/repository/mock"
	"github.com/khuchuz/go-clean-architecture/models"
	"github.com/stretchr/testify/assert"
)

func Test_SignUp_Success(t *testing.T) {
	repo := new(mock.UserStorageMock)
	uc := NewAuthUseCase(repo, "salt", []byte("secret"), 86400)
	var (
		username = "usermock"
		email    = "usermock@gmail.com"
		password = "pass"

		ctx = context.Background()

		user = &models.User{
			Username: username,
			Email:    email,
			Password: "11f5639f22525155cb0b43573ee4212838c78d87", // sha1 of pass+salt
		}
	)

	// Sign Up
	repo.On("IsUserExistByUsername", username).Return(false)
	repo.On("IsUserExistByEmail", email).Return(false)
	repo.On("CreateUser", user).Return(nil)
	err := uc.SignUp(ctx, entities.SignUpInput{Username: username, Email: email, Password: password})
	assert.NoError(t, err)
}

func Test_SignUp_Failed_DupUsername(t *testing.T) {
	repo := new(mock.UserStorageMock)
	uc := NewAuthUseCase(repo, "salt", []byte("secret"), 86400)
	var (
		username = "usermock"
		email    = "usermock@gmail.com"
		password = "pass"

		ctx = context.Background()

		user = &models.User{
			Username: username,
			Email:    email,
			Password: "11f5639f22525155cb0b43573ee4212838c78d87", // sha1 of pass+salt
		}
	)

	// Sign Up
	repo.On("IsUserExistByUsername", username).Return(true)
	repo.On("CreateUser", user).Return(nil)
	err := uc.SignUp(ctx, entities.SignUpInput{Username: username, Email: email, Password: password})
	assert.Error(t, err, auth.ErrUserDuplicate)
}

func Test_SignUp_Failed_DupEmail(t *testing.T) {
	repo := new(mock.UserStorageMock)
	uc := NewAuthUseCase(repo, "salt", []byte("secret"), 86400)
	var (
		username = "usermock"
		email    = "usermock@gmail.com"
		password = "pass"

		ctx = context.Background()

		user = &models.User{
			Username: username,
			Email:    email,
			Password: "11f5639f22525155cb0b43573ee4212838c78d87", // sha1 of pass+salt
		}
	)

	// Sign Up
	repo.On("IsUserExistByUsername", username).Return(false)
	repo.On("IsUserExistByEmail", email).Return(true)
	repo.On("CreateUser", user).Return(nil)
	err := uc.SignUp(ctx, entities.SignUpInput{Username: username, Email: email, Password: password})
	assert.Error(t, err, auth.ErrEmailDuplicate)
}
func Test_SignUp_Failed_EmptyUsername(t *testing.T) {
	repo := new(mock.UserStorageMock)
	uc := NewAuthUseCase(repo, "salt", []byte("secret"), 86400)
	var (
		username = ""
		email    = "usermock@gmail.com"
		password = "pass"

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
	assert.Error(t, err, auth.ErrDataTidakLengkap)
}

func Test_SignUp_Failed_EmptyEmail(t *testing.T) {
	repo := new(mock.UserStorageMock)
	uc := NewAuthUseCase(repo, "salt", []byte("secret"), 86400)
	var (
		username = "usermock"
		email    = ""
		password = "pass"

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
	assert.Error(t, err, auth.ErrDataTidakLengkap)
}

func Test_SignUp_Failed_Password(t *testing.T) {
	repo := new(mock.UserStorageMock)
	uc := NewAuthUseCase(repo, "salt", []byte("secret"), 86400)
	var (
		username = "usermock"
		email    = "usermock@gmail.com"
		password = ""

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
	assert.Error(t, err, auth.ErrDataTidakLengkap)
}

func Test_SignIn_Success(t *testing.T) {
	repo := new(mock.UserStorageMock)
	uc := NewAuthUseCase(repo, "salt", []byte("secret"), 86400)
	var (
		username = "usermock"
		email    = "usermock@gmail.com"
		password = "pass"

		ctx = context.Background()

		user = &models.User{
			Username: username,
			Email:    email,
			Password: "11f5639f22525155cb0b43573ee4212838c78d87", // sha1 of pass+salt
		}
	)

	// Sign In (Get Auth Token)
	repo.On("GetUser", user.Username, user.Password).Return(user, nil)
	token, err := uc.SignIn(ctx, entities.SignInput{Username: username, Password: password})
	assert.NoError(t, err)
	assert.NotEmpty(t, token)
}

func Test_SignIn_Failed(t *testing.T) {
	repo := new(mock.UserStorageMock)
	uc := NewAuthUseCase(repo, "salt", []byte("secret"), 86400)
	var (
		username = "usermock"
		email    = "usermock@gmail.com"
		password = "pass"

		ctx = context.Background()

		user = &models.User{
			Username: username,
			Email:    email,
			Password: "11f5639f22525155cb0b43573ee4212838c78d87", // sha1 of pass+salt
		}
	)

	// Sign In (Get Auth Token)
	repo.On("GetUser", user.Username, user.Password).Return(user, auth.ErrUnknown)
	token, err := uc.SignIn(ctx, entities.SignInput{Username: username, Password: password})
	assert.Error(t, err, auth.ErrUserNotFound)
	assert.Empty(t, token)
}
func Test_ParseToken_Success(t *testing.T) {
	repo := new(mock.UserStorageMock)
	uc := NewAuthUseCase(repo, "salt", []byte("secret"), 86400)
	var (
		username = "usermock"
		email    = "usermock@gmail.com"
		password = "pass"

		ctx = context.Background()

		user = &models.User{
			Username: username,
			Email:    email,
			Password: "11f5639f22525155cb0b43573ee4212838c78d87", // sha1 of pass+salt
		}
	)

	repo.On("GetUser", user.Username, user.Password).Return(user, nil)
	token, err := uc.SignIn(ctx, entities.SignInput{Username: username, Password: password})
	assert.NoError(t, err)
	assert.NotEmpty(t, token)

	// Verify token
	parsedUser, err := uc.ParseToken(ctx, token)
	assert.NoError(t, err)
	assert.Equal(t, user, parsedUser)
}

func Test_ParseToken_Failed(t *testing.T) {
	repo := new(mock.UserStorageMock)
	uc := NewAuthUseCase(repo, "salt", []byte("secret"), 86400)
	var (
		username = "usermock"
		email    = "usermock@gmail.com"

		ctx = context.Background()

		user = &models.User{
			Username: username,
			Email:    email,
			Password: "11f5639f22525155cb0b43573ee4212838c78d87", // sha1 of pass+salt
		}
	)
	var token = "mboh"

	// Verify token
	parsedUser, err := uc.ParseToken(ctx, token)
	assert.Error(t, err)
	assert.NotEqual(t, user, parsedUser)
}

func Test_ChangePassword_Sucess(t *testing.T) {
	repo := new(mock.UserStorageMock)
	uc := NewAuthUseCase(repo, "salt", []byte("secret"), 86400)
	var (
		username     = "usermock"
		email        = "usermock@gmail.com"
		password     = "pass"
		newpass      = "newpass"
		newpasscrypt = "128bef5c354a597f9160a61f267a7f06dab7d042"
		ctx          = context.Background()

		user = &models.User{
			Username: username,
			Email:    email,
			Password: "11f5639f22525155cb0b43573ee4212838c78d87", // sha1 of pass+salt
		}
	)

	// Change Password
	repo.On("GetUser", user.Username, user.Password).Return(user, nil)
	repo.On("UpdatePassword", user.Username, newpasscrypt).Return(nil)
	err := uc.ChangePassword(ctx, entities.ChangePasswordInput{Username: username, OldPassword: password, Password: newpass})
	assert.NoError(t, err)
}

func Test_ChangePassword_Failed_WrongOldPass(t *testing.T) {
	repo := new(mock.UserStorageMock)
	uc := NewAuthUseCase(repo, "salt", []byte("secret"), 86400)
	var (
		username     = "usermock"
		email        = "usermock@gmail.com"
		password     = "pass"
		newpass      = "newpass"
		newpasscrypt = "128bef5c354a597f9160a61f267a7f06dab7d042"
		ctx          = context.Background()

		user = &models.User{
			Username: username,
			Email:    email,
			Password: "11f5639f22525155cb0b43573ee4212838c78d87", // sha1 of pass+salt
		}
	)

	// Change Password
	repo.On("GetUser", user.Username, user.Password).Return(user, auth.ErrUnknown)
	repo.On("UpdatePassword", user.Username, newpasscrypt).Return(nil)
	err := uc.ChangePassword(ctx, entities.ChangePasswordInput{Username: username, OldPassword: password, Password: newpass})
	assert.Error(t, err, auth.ErrUserNotFound)
}

func Test_ChangePassword_Failed_EmptyField(t *testing.T) {
	repo := new(mock.UserStorageMock)
	uc := NewAuthUseCase(repo, "salt", []byte("secret"), 86400)
	var (
		username = "usermock"
		password = "pass"
		newpass  = "newpass"
		ctx      = context.Background()
	)

	// Empty Username
	err := uc.ChangePassword(ctx, entities.ChangePasswordInput{Username: "", OldPassword: password, Password: newpass})
	assert.EqualError(t, err, "data tidak lengkap")

	// Empty Password
	err = uc.ChangePassword(ctx, entities.ChangePasswordInput{Username: username, OldPassword: password, Password: ""})
	assert.EqualError(t, err, "data tidak lengkap")

	// Empty OldPassword
	err = uc.ChangePassword(ctx, entities.ChangePasswordInput{Username: username, OldPassword: "", Password: newpass})
	assert.EqualError(t, err, "data tidak lengkap")

}

func Test_ChangePassword_Failed_EqualNewOld(t *testing.T) {
	repo := new(mock.UserStorageMock)
	uc := NewAuthUseCase(repo, "salt", []byte("secret"), 86400)
	var (
		username = "usermock"
		password = "pass"
		ctx      = context.Background()
	)

	// Empty OldPassword
	err := uc.ChangePassword(ctx, entities.ChangePasswordInput{Username: username, OldPassword: password, Password: password})
	assert.EqualError(t, err, "password baru tidak boleh sama dengan password lama")
}
