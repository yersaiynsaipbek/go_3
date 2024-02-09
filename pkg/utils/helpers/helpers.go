package helpers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func RenderErrorPage(c *gin.Context, status int, err error) {
	c.HTML(status, "error.html", gin.H{
		"statusCode": status,
		"statusName": http.StatusText(status),
		"message":    err.Error(),
	})
	c.AbortWithStatus(status)
}
