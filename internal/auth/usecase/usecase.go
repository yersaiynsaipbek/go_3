package usecase

import (
	auth2 "go-clean-architecture/internal/auth"
	"go-clean-architecture/models"
	"go-clean-architecture/pkg/utils/hashing"
)

type AuthUseCase struct {
	userRepo auth2.UserRepository
	roleRepo auth2.RoleRepository
}

func NewAuthUseCase(
	userRepo auth2.UserRepository,
	roleRepo auth2.RoleRepository) *AuthUseCase {

	return &AuthUseCase{
		userRepo: userRepo,
		roleRepo: roleRepo,
	}
}

func (a *AuthUseCase) Register(user *models.User) error {
	hashedPassword, err := hashing.HashPassword(user.Password)
	if err != nil {
		return err
	}
	user.Password = hashedPassword

	roleID, err := a.roleRepo.GetIDByName(models.STUDENT_ROLE)
	if err != nil {
		return err
	}

	return a.userRepo.CreateUser(user, roleID)
}

func (a *AuthUseCase) Login(username, password string) (*models.User, error) {
	user, err := a.userRepo.GetUserByUsername(username)
	if err != nil {
		return nil, auth2.ErrUserNotFound
	}

	if !hashing.CheckPasswordHash(password, user.Password) {
		return nil, auth2.ErrIncorrectCredentials
	}
	return user, nil
}
