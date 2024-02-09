package http

import (
	"errors"
	"github.com/gin-gonic/gin"
	auth2 "go-clean-architecture/internal/auth"
	"go-clean-architecture/models"
	"go-clean-architecture/pkg/utils/cookies"
	"go-clean-architecture/pkg/utils/helpers"
	"go-clean-architecture/pkg/utils/usersession"
	"net/http"
)

type Handler struct {
	useCase auth2.UseCase
}

func NewHandler(useCase auth2.UseCase) *Handler {
	return &Handler{useCase: useCase}
}

func (h *Handler) LoginForm(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", nil)
}

func (h *Handler) RegisterForm(c *gin.Context) {
	c.HTML(http.StatusOK, "register.html", nil)
}

func (h *Handler) Register(c *gin.Context) {
	var registerUser models.User

	name := c.PostForm("name")
	surname := c.PostForm("surname")
	username := c.PostForm("username")
	password := c.PostForm("password")

	if name == "" || surname == "" || username == "" || password == "" {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	registerUser.Name = name
	registerUser.Surname = surname
	registerUser.Username = username
	registerUser.Password = password

	if err := h.useCase.Register(&registerUser); err != nil {
		helpers.RenderErrorPage(c, http.StatusInternalServerError, err)
		return
	}

	c.Redirect(http.StatusSeeOther, "/auth/login/")
}

func (h *Handler) Login(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")

	if username == "" || password == "" {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	user, err := h.useCase.Login(username, password)
	if err != nil {
		if errors.Is(err, auth2.ErrUserNotFound) {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		c.Status(http.StatusInternalServerError)
		helpers.RenderErrorPage(c, http.StatusInternalServerError, err)
		return
	}

	if err := usersession.SetSessionLoggedID(c, user.ID); err != nil {
		helpers.RenderErrorPage(c, http.StatusInternalServerError, err)
		return
	}

	c.Redirect(http.StatusSeeOther, "/api/news/")
}

func (h *Handler) Logout(c *gin.Context) {
	cookies.DeleteCookie(c)
	c.Redirect(http.StatusSeeOther, "/auth/login/")
}
