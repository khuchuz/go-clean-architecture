package delivery

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/khuchuz/go-clean-architecture/auth"
	"github.com/khuchuz/go-clean-architecture/auth/usecase/mock"
	"github.com/khuchuz/go-clean-architecture/models"
	"github.com/stretchr/testify/assert"
)

func Test_Middleware_Success(t *testing.T) {
	r := gin.Default()
	uc := new(mock.AuthUseCaseMock)

	r.POST("/api/endpoint", NewAuthMiddleware(uc), func(c *gin.Context) {
		c.Status(http.StatusOK)
	})

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/api/endpoint", nil)

	// Valid Auth Header
	uc.On("ParseToken", "token").Return(&models.User{}, nil)
	req.Header.Set("Authorization", "Bearer token")
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
}

func Test_Middleware_Unauthorized1(t *testing.T) {
	r := gin.Default()
	uc := new(mock.AuthUseCaseMock)

	r.POST("/api/endpoint", NewAuthMiddleware(uc), func(c *gin.Context) {
		c.Status(http.StatusOK)
	})

	w := httptest.NewRecorder()

	// No Auth Header request
	req, _ := http.NewRequest("POST", "/api/endpoint", nil)
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusUnauthorized, w.Code)
}

func Test_Middleware_Unauthorized2(t *testing.T) {
	r := gin.Default()
	uc := new(mock.AuthUseCaseMock)

	r.POST("/api/endpoint", NewAuthMiddleware(uc), func(c *gin.Context) {
		c.Status(http.StatusOK)
	})

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/api/endpoint", nil)
	// Empty Auth Header request

	req.Header.Set("Authorization", "")
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusUnauthorized, w.Code)
}

func Test_Middleware_Unauthorized3(t *testing.T) {
	r := gin.Default()
	uc := new(mock.AuthUseCaseMock)

	r.POST("/api/endpoint", NewAuthMiddleware(uc), func(c *gin.Context) {
		c.Status(http.StatusOK)
	})

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/api/endpoint", nil)
	// Bearer Auth Header with no token request
	uc.On("ParseToken", "").Return(&models.User{}, auth.ErrInvalidAccessToken)

	req.Header.Set("Authorization", "Bearer ")
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusUnauthorized, w.Code)
}

func Test_Middleware_Unauthorized4(t *testing.T) {
	r := gin.Default()
	uc := new(mock.AuthUseCaseMock)

	r.POST("/api/endpoint", NewAuthMiddleware(uc), func(c *gin.Context) {
		c.Status(http.StatusOK)
	})

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/api/endpoint", nil)

	// Valid Auth Header
	uc.On("ParseToken", "token").Return(&models.User{}, nil)
	req.Header.Set("Authorization", "Bearer token ")
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusUnauthorized, w.Code)
}

func Test_Middleware_Unauthorized5(t *testing.T) {
	r := gin.Default()
	uc := new(mock.AuthUseCaseMock)

	r.POST("/api/endpoint", NewAuthMiddleware(uc), func(c *gin.Context) {
		c.Status(http.StatusOK)
	})

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/api/endpoint", nil)

	// Valid Auth Header
	uc.On("ParseToken", "token").Return(&models.User{}, nil)
	req.Header.Set("Authorization", "Bukan token")
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusUnauthorized, w.Code)
}
