package delivery

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/khuchuz/go-clean-architecture/auth"
)

type AuthMiddleware struct {
	usecase auth.UseCase
}

func NewAuthMiddleware(usecase auth.UseCase) gin.HandlerFunc {
	return (&AuthMiddleware{
		usecase: usecase,
	}).Handle
}

func (m *AuthMiddleware) Handle(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		c.JSON(http.StatusUnauthorized, signResponse{Message: auth.ErrUnauthorized.Error()})
		return
	}

	headerParts := strings.Split(authHeader, " ")
	if len(headerParts) != 2 {
		c.JSON(http.StatusUnauthorized, signResponse{Message: auth.ErrUnauthorized.Error()})
		return
	}

	if headerParts[0] != "Bearer" {
		c.JSON(http.StatusUnauthorized, signResponse{Message: auth.ErrUnauthorized.Error()})
		return
	}

	user, err := m.usecase.ParseToken(c.Request.Context(), headerParts[1])
	if err != nil {
		status := http.StatusInternalServerError
		if err == auth.ErrInvalidAccessToken {
			status = http.StatusUnauthorized
		}

		c.JSON(status, signResponse{Message: auth.ErrUnknown.Error()})
		return
	}

	c.Set(auth.CtxUserKey, user)
}
