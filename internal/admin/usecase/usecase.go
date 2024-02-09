package usecase

import (
	userAndRoleRepo "go-clean-architecture/internal/auth"
	"go-clean-architecture/internal/news"
	"go-clean-architecture/models"
)

type AdminUseCase struct {
	userRepo userAndRoleRepo.UserRepository
	roleRepo userAndRoleRepo.RoleRepository
	newsRepo news.NewsRepository
}

func NewAdminUseCase(
	userRepo userAndRoleRepo.UserRepository,
	roleRepo userAndRoleRepo.RoleRepository,
	newsRepo news.NewsRepository) *AdminUseCase {
	return &AdminUseCase{
		userRepo: userRepo,
		roleRepo: roleRepo,
		newsRepo: newsRepo,
	}
}

func (a *AdminUseCase) GetAllUsers() (*[]models.User, error) {
	userList, err := a.userRepo.GetAllUsers()
	if err != nil {
		return nil, err
	}

	return userList, nil
}

func (a *AdminUseCase) GetUserByID(userID int) (*models.User, error) {
	user, err := a.userRepo.GetUserByID(userID)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (a *AdminUseCase) checkUserRoleByID(userID int, expectedRole string) (bool, error) {
	user, err := a.userRepo.GetUserByID(userID)
	if err != nil {
		return false, err
	}

	return user.Role == expectedRole, nil
}

func (a *AdminUseCase) IsAdmin(userID int) (bool, error) {
	return a.checkUserRoleByID(userID, models.ADMIN_ROLE)
}

func (a *AdminUseCase) UpdateUserByID(userID int, user *models.User) error {
	roleID, err := a.roleRepo.GetIDByName(user.Role)
	if err != nil {
		return err
	}
	user.RoleID = roleID

	return a.userRepo.UpdateUserByID(userID, user)
}

func (a *AdminUseCase) DeleteUserByID(userID int) error {
	return a.userRepo.DeleteUserByID(userID)
}
