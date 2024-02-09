package http

import (
	"errors"
	"go-clean-architecture/pkg/utils/helpers"
	"go-clean-architecture/pkg/utils/usersession"
	"net/http"

	"github.com/gin-gonic/gin"
	"go-clean-architecture/internal/admin"
)

type AdminMiddleware struct {
	usecase admin.UseCase
}

func NewAdminMiddleware(usecase admin.UseCase) gin.HandlerFunc {
	return (&AdminMiddleware{
		usecase: usecase,
	}).Handle
}

func (m *AdminMiddleware) Handle(c *gin.Context) {
	userID, err := usersession.GetLoggedUserID(c)
	if err != nil {
		if errors.Is(err, usersession.ErrSessionUserNotFound) {
			c.Redirect(http.StatusUnauthorized, "/auth/login")
			c.Abort()
			return
		}

		helpers.RenderErrorPage(c, http.StatusForbidden, err)
		c.Abort()
		return
	}

	isAdmin, err := m.usecase.IsAdmin(userID)
	if err != nil {
		helpers.RenderErrorPage(c, http.StatusInternalServerError, err)
		c.Abort()
		return
	}

	if !isAdmin {
		helpers.RenderErrorPage(c, http.StatusForbidden, err)
		c.Abort()
		return
	}

	c.Next()
}
