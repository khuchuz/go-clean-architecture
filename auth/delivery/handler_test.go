package delivery

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/khuchuz/go-clean-architecture/auth"
	"github.com/khuchuz/go-clean-architecture/auth/entities"
	"github.com/khuchuz/go-clean-architecture/auth/usecase/mock"
	"github.com/stretchr/testify/assert"
)

func TestSignUp_Success_200(t *testing.T) {
	r := gin.Default()
	uc := new(mock.AuthUseCaseMock)

	RegisterHTTPEndpoints(r, uc)

	signUpBody := &entities.SignUpInput{
		Username: "testuser",
		Email:    "testuser@gmail.com",
		Password: "testpass",
	}

	body, err := json.Marshal(signUpBody)
	assert.NoError(t, err)

	uc.On("SignUp", signUpBody.Username, signUpBody.Email, signUpBody.Password).Return(nil)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/auth/sign-up", bytes.NewBuffer(body))
	r.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
}
func TestSignUp_Failed_500(t *testing.T) {
	r := gin.Default()
	uc := new(mock.AuthUseCaseMock)

	RegisterHTTPEndpoints(r, uc)

	signUpBody := &entities.SignUpInput{
		Username: "testuser",
		Email:    "testuser@gmail.com",
		Password: "testpass",
	}

	body, err := json.Marshal(signUpBody)
	assert.NoError(t, err)

	uc.On("SignUp", signUpBody.Username, signUpBody.Email, signUpBody.Password).Return(errors.New("err"))

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/auth/sign-up", bytes.NewBuffer(body))
	r.ServeHTTP(w, req)

	assert.Equal(t, 500, w.Code)
}
func TestSignIn_Sucess_200(t *testing.T) {
	r := gin.Default()
	uc := new(mock.AuthUseCaseMock)

	RegisterHTTPEndpoints(r, uc)

	signInBody := &entities.SignInput{
		Username: "testuser",
		Password: "testpass",
	}

	body, err := json.Marshal(signInBody)
	assert.NoError(t, err)

	uc.On("SignIn", signInBody.Username, signInBody.Password).Return("jwt", nil)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/auth/sign-in", bytes.NewBuffer(body))
	r.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "{\"token\":\"jwt\"}", w.Body.String())
}
func TestSignIn_ErrUserNotFound(t *testing.T) {
	r := gin.Default()
	uc := new(mock.AuthUseCaseMock)

	RegisterHTTPEndpoints(r, uc)

	signInBody := &entities.SignInput{
		Username: "testuser",
		Password: "testpass",
	}

	body, err := json.Marshal(signInBody)
	assert.NoError(t, err)

	uc.On("SignIn", signInBody.Username, signInBody.Password).Return("", auth.ErrUserNotFound)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/auth/sign-in", bytes.NewBuffer(body))
	r.ServeHTTP(w, req)

	assert.Equal(t, 401, w.Code)
	assert.Equal(t, "{\"message\":\"user not found\"}", w.Body.String())
}

func TestSignIn_ErrUnknown(t *testing.T) {
	r := gin.Default()
	uc := new(mock.AuthUseCaseMock)

	RegisterHTTPEndpoints(r, uc)

	signInBody := &entities.SignInput{
		Username: "testuser",
		Password: "testpass",
	}

	body, err := json.Marshal(signInBody)
	assert.NoError(t, err)

	uc.On("SignIn", signInBody.Username, signInBody.Password).Return("", auth.ErrUnauthorized)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/auth/sign-in", bytes.NewBuffer(body))
	r.ServeHTTP(w, req)

	assert.Equal(t, 401, w.Code)
	assert.Equal(t, "{\"message\":\"unknown error\"}", w.Body.String())
}

func TestChangePassword_Success(t *testing.T) {
	r := gin.Default()
	uc := new(mock.AuthUseCaseMock)

	RegisterHTTPEndpoints(r, uc)

	changePassBody := &entities.ChangePasswordInput{
		Username:    "testuser",
		Password:    "newpass",
		OldPassword: "testpass",
	}

	body, err := json.Marshal(changePassBody)
	assert.NoError(t, err)

	uc.On("ChangePassword", changePassBody.Username, changePassBody.OldPassword, changePassBody.Password).Return("", auth.ErrUnauthorized)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/auth/change-pass", bytes.NewBuffer(body))
	r.ServeHTTP(w, req)

	assert.Equal(t, 401, w.Code)
	assert.Equal(t, "{\"message\":\"unknown error\"}", w.Body.String())
}
