package usecase

import (
	"context"
	"github.com/ppmkh2sby/backend-library/models"
)

// GetAllUsers is function on usecase to handle get all users from database
func (u *userUsecase) GetAllUsers(ctx context.Context) ([]models.Users, error) {
	users, err := u.userRepository.GetAllUsers(ctx)
	if err != nil {
		return nil, err
	}

	return users, nil
}
