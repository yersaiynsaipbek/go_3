package http

import (
	"errors"
	"github.com/gin-gonic/gin"
	"go-clean-architecture/internal/news"
	"go-clean-architecture/models"
	"go-clean-architecture/pkg/utils/helpers"
	"net/http"
	"strconv"
)

type Handler struct {
	useCase news.UseCase
}

func NewHandler(useCase news.UseCase) *Handler {
	return &Handler{
		useCase: useCase,
	}
}

func (h *Handler) CreateNewsForm(c *gin.Context) {
	c.HTML(http.StatusOK, "create-news.html", nil)
}

func (h *Handler) CreateNews(c *gin.Context) {
	var news models.News

	if err := c.BindJSON(&news); err != nil {
		helpers.RenderErrorPage(c, http.StatusInternalServerError, err)
		return
	}

	if err := h.useCase.CreateNews(&news); err != nil {
		helpers.RenderErrorPage(c, http.StatusInternalServerError, err)
		return
	}

	c.Status(http.StatusCreated)
}

func (h *Handler) GetAllNews(c *gin.Context) {
	newsList, err := h.useCase.GetAllNews()
	if err != nil {
		helpers.RenderErrorPage(c, http.StatusInternalServerError, err)
		return
	}

	c.HTML(http.StatusOK, "news-list.html", newsList)
}

func (h *Handler) GetNewsByID(c *gin.Context) {
	newsID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		helpers.RenderErrorPage(c, http.StatusBadRequest, err)
		return
	}

	newsData, err := h.useCase.GetNewsByID(newsID)
	if err != nil {
		if errors.Is(err, news.ErrNewsNotFound) {
			helpers.RenderErrorPage(c, http.StatusNotFound, err)
			return
		}

		helpers.RenderErrorPage(c, http.StatusInternalServerError, err)
		return
	}

	c.HTML(http.StatusOK, "news-data.html", newsData)
}

func (h *Handler) DeleteNewsByID(c *gin.Context) {
	newsID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		helpers.RenderErrorPage(c, http.StatusBadRequest, err)
		return
	}

	if err := h.useCase.DeleteNewsByID(newsID); err != nil {
		if errors.Is(err, news.ErrNewsNotFound) {
			helpers.RenderErrorPage(c, http.StatusNotFound, err)
			return
		}

		helpers.RenderErrorPage(c, http.StatusInternalServerError, err)
		return
	}

	c.Status(http.StatusNoContent)
	//c.Redirect(http.StatusNoContent, "/api/news/")
}
