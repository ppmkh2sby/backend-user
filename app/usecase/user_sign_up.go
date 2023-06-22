package usecase

import (
	"context"
	"github.com/ppmkh2sby/backend-library/helpers/generate"
	"github.com/ppmkh2sby/backend-library/models"
)

// SignUpUser is function on usecase to handle create user
func (u *userUsecase) SignUpUser(ctx context.Context, user *models.Users) (*models.Users, error) {
	var err error
	user.ID, err = generate.CreateUUID()
	if err != nil {
		return nil, err
	}

	user.Password, err = generate.HashPassword(user.Password)
	if err != nil {
		return nil, err
	}

	user, err = u.userRepository.CreateUser(ctx, user)
	if err != nil {
		return nil, err
	}

	userResponse := &models.Users{
		ID:        user.ID,
		Username:  user.Username,
		Email:     user.Email,
		CardID:    user.CardID,
		SantriID:  user.SantriID,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}

	return userResponse, err
}
