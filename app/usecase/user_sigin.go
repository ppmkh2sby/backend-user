package usecase

import (
	"github.com/ppmkh2sby/backend-library/helpers/generate"
	"github.com/ppmkh2sby/backend-library/models"
)

// SingInUser is funstion to sign in user
func (u *userUsecase) SignInUser(user *models.Users) (*models.Users, error) {
	userDB, err := u.userRepository.GetUserByUsername(user.Username)
	if err != nil {
		return nil, err
	}

	err = generate.CheckPasswordAndHash(userDB.Password, user.Password)
	if err != nil {
		return nil, err
	}

	userResponse, err := u.userRepository.GetUserByID(userDB.ID)
	if err != nil {
		return nil, err
	}

	return userResponse, nil
}
