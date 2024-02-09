package usersession

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
	"github.com/no-src/gin-session-redis/redis"
	"go-clean-architecture/config"
	"strconv"
)

var (
	store       sessions.Store
	sessionName string
	sessionKey  string
)

var (
	ErrSessionUserNotFound = errors.New("logged user ID not found in session or is of incorrect type")
)

func NewUserSessionStore(config *config.Session) {
	connectionTimeoutSeconds, _ := strconv.Atoi(config.Redis.ConnectionTimeoutSeconds)
	store, _ = redis.NewStore(
		connectionTimeoutSeconds,
		config.Redis.NetworkType,
		config.Redis.Host+":"+config.Redis.Port,
		config.Redis.Password,
		[]byte(config.SessionKey),
	)
	sessionName = config.SessionName
	sessionKey = config.SessionKey
}

func IsAuthenticated(c *gin.Context) bool {
	userID, err := GetLoggedUserID(c)
	if err != nil || userID == -1 {
		return false
	}
	return true
}

func GetLoggedUserID(c *gin.Context) (int, error) {
	session, err := store.Get(c.Request, sessionName)
	if err != nil {
		return -1, err
	}

	userID, ok := session.Values[sessionKey].(int)
	if !ok {
		return -1, ErrSessionUserNotFound
	}

	return userID, nil
}

func SetSessionLoggedID(c *gin.Context, userID int) error {
	session, err := store.New(c.Request, sessionName)
	if err != nil {
		return err
	}
	session.Values[sessionKey] = userID
	err = session.Save(c.Request, c.Writer)
	if err != nil {
		return err
	}
	return nil
}
