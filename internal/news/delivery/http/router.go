package http

import (
	"github.com/gin-gonic/gin"
	"go-clean-architecture/internal/news"
)

func NewsHTTPEndpoints(router *gin.RouterGroup, uc news.UseCase) {
	handler := NewHandler(uc)

	newsEndpoints := router.Group("/news")
	{
		newsEndpoints.GET("/", handler.GetAllNews)
		newsEndpoints.GET("/:id", handler.GetNewsByID)

		newsEndpoints.GET("/add", handler.CreateNewsForm)
		newsEndpoints.POST("/add", handler.CreateNews)

		newsEndpoints.DELETE("/:id/delete", handler.DeleteNewsByID)
	}
}
