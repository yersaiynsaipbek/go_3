package postgres

import (
	"database/sql"
	"errors"
	newserror "go-clean-architecture/internal/news"
	"go-clean-architecture/models"
	"log"
)

type NewsRepository struct {
	newsDB *sql.DB
}

func NewNewsRepository(newsDB *sql.DB) *NewsRepository {
	return &NewsRepository{
		newsDB: newsDB,
	}
}

func (r *NewsRepository) CreateNews(news *models.News) error {
	query := "INSERT INTO news(category, title, content, estimation) VALUES ($1, $2, $3, $4)"

	stmt, err := r.newsDB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(news.Title, news.Content, news.Estimation, news.Category)
	if err != nil {
		return err
	}

	return nil
}

func (r *NewsRepository) GetAllNews() (*[]models.News, error) {
	query := "SELECT * FROM news"

	rows, err := r.newsDB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	aituNews := []models.News{}

	for rows.Next() {
		p := models.News{}
		err := rows.Scan(&p.ID, &p.Category, &p.Title, &p.Content, &p.Estimation)
		if err != nil {
			log.Println(err)
			continue
		}
		aituNews = append(aituNews, p)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return &aituNews, nil
}

func (r *NewsRepository) GetNewsByID(newsID int) (*models.News, error) {
	query := "SELECT * FROM news WHERE id = $1"

	row := r.newsDB.QueryRow(query, newsID)

	news := &models.News{}
	err := row.Scan(&news.ID, &news.Category, &news.Title, &news.Content, &news.Estimation)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, newserror.ErrNewsNotFound
		}
		return nil, err
	}

	return news, nil
}

func (r *NewsRepository) DeleteNewsByID(newsID int) error {
	query := "DELETE FROM news WHERE id = $1"

	_, err := r.newsDB.Exec(query, newsID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return newserror.ErrNewsNotFound
		}

		return err
	}

	return nil
}
