package usecase

import (
	"context"
	"github.com/ppmkh2sby/backend-library/helpers/generate"
	"github.com/ppmkh2sby/backend-library/models"
)

// SingInUser is function on usecase to handle sign in user
func (u *userUsecase) SignInUser(ctx context.Context, user *models.Users) (*models.Users, error) {
	userDB, err := u.userRepository.GetUserByUsername(ctx, user.Username)
	if err != nil {
		return nil, err
	}

	err = generate.CheckPasswordAndHash(userDB.Password, user.Password)
	if err != nil {
		return nil, err
	}

	userResponse, err := u.userRepository.GetUserByID(ctx, userDB.ID)
	if err != nil {
		return nil, err
	}

	return userResponse, nil
}
