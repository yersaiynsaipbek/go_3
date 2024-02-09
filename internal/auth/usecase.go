package auth

import "go-clean-architecture/models"

type UseCase interface {
	Register(user *models.User) error
	Login(username, password string) (*models.User, error)
}
