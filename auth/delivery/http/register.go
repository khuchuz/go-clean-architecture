package http

import (
	"github.com/gin-gonic/gin"
	"github.com/khuchuz/go-clean-architecture/auth"
)

func RegisterHTTPEndpoints(router *gin.Engine, uc auth.UseCase) {
	h := NewHandler(uc)

	authEndpoints := router.Group("/auth")
	{
		authEndpoints.POST("/sign-up", h.SignUp)
		authEndpoints.POST("/sign-in", h.SignIn)
		authEndpoints.POST("/change-pass", h.ChangePassword)
		authEndpoints.POST("/init-change-pass", h.InitChangePassword)
	}
}
