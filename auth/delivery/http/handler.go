package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/khuchuz/go-clean-architecture/auth"
)

type Handler struct {
	useCase auth.UseCase
}

func NewHandler(useCase auth.UseCase) *Handler {
	return &Handler{
		useCase: useCase,
	}
}

type signInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type changePasswordInput struct {
	Username    string `json:"username"`
	OldPassword string `json:"oldpassword"`
	Password    string `json:"password"`
}

type signUpInput struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type signResponse struct {
	Message string `json:"message"`
}

func (h *Handler) SignUp(c *gin.Context) {
	inp := new(signUpInput)

	if err := c.BindJSON(inp); err != nil {
		c.JSON(http.StatusBadRequest, signResponse{Message: auth.ErrBadRequest.Error()})
		return
	}

	if err := h.useCase.SignUp(c.Request.Context(), inp.Username, inp.Email, inp.Password); err != nil {
		c.JSON(http.StatusInternalServerError, signResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, signResponse{Message: "Sign Up Berhasil"})
}

type signInResponse struct {
	Token string `json:"token"`
}

func (h *Handler) SignIn(c *gin.Context) {
	inp := new(signInput)

	if err := c.BindJSON(inp); err != nil {
		c.JSON(http.StatusBadRequest, signResponse{Message: auth.ErrBadRequest.Error()})
		return
	}

	token, err := h.useCase.SignIn(c.Request.Context(), inp.Username, inp.Password)
	if err != nil {
		if err == auth.ErrUserNotFound {
			c.JSON(http.StatusUnauthorized, signResponse{Message: auth.ErrUserNotFound.Error()})
			return
		}
		c.JSON(http.StatusUnauthorized, signResponse{Message: auth.ErrUnknown.Error()})
		return
	}

	c.JSON(http.StatusOK, signInResponse{Token: token})
}

func (h *Handler) ChangePassword(c *gin.Context) {
	inp := new(changePasswordInput)

	if err := c.BindJSON(inp); err != nil {
		c.JSON(http.StatusBadRequest, signResponse{Message: auth.ErrBadRequest.Error()})
		return
	}

	err := h.useCase.ChangePassword(c.Request.Context(), inp.Username, inp.OldPassword, inp.Password)
	if err != nil {
		if err == auth.ErrUserNotFound {
			c.JSON(http.StatusUnauthorized, signResponse{Message: auth.ErrUserNotFound.Error()})
			return
		}
		c.JSON(http.StatusUnauthorized, signResponse{Message: auth.ErrUnknown.Error()})
		return
	}

	c.JSON(http.StatusOK, signResponse{Message: "Password berhasil diubah"})
}
