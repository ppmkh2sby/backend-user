package usecase

import (
	"context"
	"github.com/ppmkh2sby/backend-library/models"
)

// UpdateUser is function on usecase to handle update user's data
func (u *userUsecase) UpdateUser(ctx context.Context, user *models.Users) (*models.Users, error) {
	err := u.userRepository.UpdateUser(ctx, user)
	if err != nil {
		return nil, err
	}

	return user, nil
}
