package usecase

import (
	"context"
	"github.com/ppmkh2sby/backend-library/models"
	"github.com/ppmkh2sby/backend-user/app/repository"
)

type userUsecase struct {
	userRepository repository.UserRepository
}

type UserUsecase interface {
	SignUpUser(ctx context.Context, user *models.Users) (*models.Users, error)
	SignInUser(ctx context.Context, user *models.Users) (*models.Users, error)
	GetAllUsers(ctx context.Context) ([]models.Users, error)
	GetUserByID(ctx context.Context, id string) (*models.Users, error)
	UpdateUser(ctx context.Context, user *models.Users) (*models.Users, error)
	DeleteUser(ctx context.Context, id string) error
	ChangePasswordUser(ctx context.Context, id, password string) error
}

// NewUserUsecase is funtion to initial User Usecase
func NewUserUsecase(u repository.UserRepository) *userUsecase {
	return &userUsecase{
		userRepository: u,
	}
}
