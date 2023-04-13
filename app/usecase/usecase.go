package usecase

import (
	"github.com/ppmkh2sby/backend-library/models"
	"github.com/ppmkh2sby/backend-user/app/repository"
)

type userUsecase struct {
	userRepository repository.UserRepository
}

type UserUsecase interface {
	SignUpUser(user *models.Users) (*models.Users, error)
	SignInUser(user *models.Users) (*models.Users, error)
	GetAllUsers() ([]models.Users, error)
	GetUserByID(id string) (*models.Users, error)
	UpdateUser(user *models.Users) (*models.Users, error)
	DeleteUser(id string) error
	ChangePasswordUser(id, password string) error
}

// NewUserUsecase is funtion to initial User Usecase
func NewUserUsecase(u repository.UserRepository) *userUsecase {
	return &userUsecase{
		userRepository: u,
	}
}
