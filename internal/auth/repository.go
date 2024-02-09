package auth

import (
	"go-clean-architecture/models"
)

type UserRepository interface {
	CreateUser(user *models.User, roleID int) error
	GetUserByUsername(username string) (*models.User, error)
	GetUserByID(userID int) (*models.User, error)
	GetAllUsers() (*[]models.User, error)
	UpdateUserByID(userID int, user *models.User) error
	DeleteUserByID(userID int) error
}

type RoleRepository interface {
	CreateRole(role string) error
	GetNameByID(roleID int) (string, error)
	GetIDByName(roleName string) (int, error)
}
