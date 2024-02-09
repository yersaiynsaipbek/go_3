package usecase

import (
	"go-clean-architecture/internal/news"
	"go-clean-architecture/models"
)

type NewsUseCase struct {
	newsDB news.NewsRepository
}

func NewNewsUseCase(newsDB news.NewsRepository) *NewsUseCase {
	return &NewsUseCase{
		newsDB: newsDB,
	}
}

func (n *NewsUseCase) CreateNews(news *models.News) error {
	err := n.newsDB.CreateNews(news)
	if err != nil {
		return err
	}
	return nil
}

func (n *NewsUseCase) GetAllNews() (*[]models.News, error) {
	newsList, err := n.newsDB.GetAllNews()
	if err != nil {
		return nil, err
	}
	return newsList, nil
}

func (n *NewsUseCase) GetNewsByID(newsID int) (*models.News, error) {
	newsItem, err := n.newsDB.GetNewsByID(newsID)
	if err != nil {
		return nil, err
	}
	return newsItem, nil
}

func (n *NewsUseCase) DeleteNewsByID(newsID int) error {
	err := n.newsDB.DeleteNewsByID(newsID)
	if err != nil {
		return err
	}
	return nil
}
