package news

import "go-clean-architecture/models"

type NewsRepository interface {
	CreateNews(news *models.News) error
	GetAllNews() (*[]models.News, error)
	GetNewsByID(newsID int) (*models.News, error)
	DeleteNewsByID(newsID int) error
}
