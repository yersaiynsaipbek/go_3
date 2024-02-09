package cookies

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

const (
	cookieName = "GSESSIONID"
)

func SetCookie(c *gin.Context, value string, maxAge time.Time) {
	cookie := &http.Cookie{
		Name:     cookieName,
		Value:    value,
		HttpOnly: true,
		Path:     "/",
		Expires:  maxAge,
	}
	http.SetCookie(c.Writer, cookie)
}

func GetCookie(c *gin.Context) (*http.Cookie, error) {
	cookie, err := c.Request.Cookie(cookieName)
	if err != nil {
		return nil, err
	}
	return cookie, nil
}

func DeleteCookie(c *gin.Context) {
	cookie := &http.Cookie{
		Name:     cookieName,
		Value:    "",
		HttpOnly: true,
		Path:     "/",
		MaxAge:   -1,
	}
	http.SetCookie(c.Writer, cookie)
}
