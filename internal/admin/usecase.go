package admin

import "go-clean-architecture/models"

type UseCase interface {
	IsAdmin(userID int) (bool, error)
	GetAllUsers() (*[]models.User, error)
	GetUserByID(userID int) (*models.User, error)
	UpdateUserByID(userID int, user *models.User) error
	DeleteUserByID(userID int) error
}
