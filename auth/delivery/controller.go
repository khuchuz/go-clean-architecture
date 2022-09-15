package delivery

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/khuchuz/go-clean-architecture/auth"
	"github.com/khuchuz/go-clean-architecture/auth/entities"
	itface "github.com/khuchuz/go-clean-architecture/auth/itface"
)

type Handler struct {
	useCase itface.UseCase
}

func NewHandler(useCase itface.UseCase) *Handler {
	return &Handler{
		useCase: useCase,
	}
}

func (h *Handler) SignUp(c *gin.Context) {
	inp := new(entities.SignUpInput)

	if err := c.BindJSON(inp); err != nil {
		c.JSON(http.StatusBadRequest, signResponse{Message: auth.ErrBadRequest.Error()})
		return
	}

	if err := h.useCase.SignUp(c.Request.Context(), *inp); err != nil {
		c.JSON(http.StatusInternalServerError, signResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, signResponse{Message: "Sign Up Berhasil"})
}

func (h *Handler) SignIn(c *gin.Context) {
	inp := new(entities.SignInput)

	if err := c.BindJSON(inp); err != nil {
		c.JSON(http.StatusBadRequest, signResponse{Message: auth.ErrBadRequest.Error()})
		return
	}

	token, err := h.useCase.SignIn(c.Request.Context(), *inp)
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
	inp := new(entities.ChangePasswordInput)

	if err := c.BindJSON(inp); err != nil {
		c.JSON(http.StatusBadRequest, signResponse{Message: auth.ErrBadRequest.Error()})
		return
	}

	err := h.useCase.ChangePassword(c.Request.Context(), *inp)
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
