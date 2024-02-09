package http

import (
	"go-clean-architecture/internal/auth"
	"net/http"

	"github.com/gin-gonic/gin"
	"go-clean-architecture/pkg/utils/usersession"
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
	if !usersession.IsAuthenticated(c) {
		c.Redirect(http.StatusUnauthorized, "/auth/login")
		c.Abort()
		return
	}

	c.Next()
}
