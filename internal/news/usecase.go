package news

import "go-clean-architecture/models"

type UseCase interface {
	CreateNews(news *models.News) error
	GetAllNews() (*[]models.News, error)
	GetNewsByID(newsID int) (*models.News, error)
	// TODO: Update method
	DeleteNewsByID(newsID int) error
}
