package http

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-clean-architecture/internal/admin"
	"go-clean-architecture/models"
	"go-clean-architecture/pkg/utils/helpers"
	"net/http"
	"strconv"
)

type Handler struct {
	useCase admin.UseCase
}

func NewHandler(useCase admin.UseCase) *Handler {
	return &Handler{
		useCase: useCase,
	}
}

func (h *Handler) GetAllUsers(c *gin.Context) {
	users, err := h.useCase.GetAllUsers()
	if err != nil {
		helpers.RenderErrorPage(c, http.StatusInternalServerError, err)
		return
	}

	c.HTML(http.StatusOK, "user-list.html", users)
}

func (h *Handler) GetUserDataByID(c *gin.Context) {
	newsID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	user, err := h.useCase.GetUserByID(newsID)
	if err != nil {
		helpers.RenderErrorPage(c, http.StatusInternalServerError, err)
		return
	}

	c.HTML(http.StatusOK, "update-user.html", user)
}

func (h *Handler) ChangeUserDataByID(c *gin.Context) {
	userID, err := strconv.Atoi(c.Param("id"))
	if err != nil {

		helpers.RenderErrorPage(c, http.StatusInternalServerError, err)
		return
	}

	var user models.User
	if err := c.BindJSON(&user); err != nil {
		fmt.Println(err.Error())
		helpers.RenderErrorPage(c, http.StatusBadRequest, err)
		return
	}

	if err := h.useCase.UpdateUserByID(userID, &user); err != nil {
		helpers.RenderErrorPage(c, http.StatusInternalServerError, err)
		return
	}

	c.Status(http.StatusOK)
}

func (h *Handler) DeleteUserByID(c *gin.Context) {
	userID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		helpers.RenderErrorPage(c, http.StatusInternalServerError, err)
		return
	}

	if err := h.useCase.DeleteUserByID(userID); err != nil {
		helpers.RenderErrorPage(c, http.StatusNotFound, err)
		return
	}

	c.Status(http.StatusNoContent)
}
