package http

import (
	"github.com/gin-gonic/gin"
	"go-clean-architecture/internal/admin"
)

func AdminHTTPEndpoints(router *gin.RouterGroup, uc admin.UseCase) {
	handler := NewHandler(uc)

	router.GET("/users/", handler.GetAllUsers)
	router.GET("/users/:id", handler.GetUserDataByID)
	router.PUT("/users/:id/update", handler.ChangeUserDataByID)
	router.DELETE("/users/:id/delete", handler.DeleteUserByID)
}
