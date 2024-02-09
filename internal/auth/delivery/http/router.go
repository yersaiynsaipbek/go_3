package http

import (
	"github.com/gin-gonic/gin"
	"go-clean-architecture/internal/auth"
)

func AuthHTTPEndpoints(router *gin.Engine, uc auth.UseCase) {
	handler := NewHandler(uc)

	authEndpoints := router.Group("/auth")
	{
		authEndpoints.GET("/login", handler.LoginForm)
		authEndpoints.POST("/login", handler.Login)

		authEndpoints.GET("/register", handler.RegisterForm)
		authEndpoints.POST("/register", handler.Register)

		authEndpoints.POST("/logout", handler.Logout)
	}
}
