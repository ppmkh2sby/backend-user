package usecase

import "github.com/ppmkh2sby/backend-library/models"

// UpdateUser is function on usecase to handle update user's data
func (u *userUsecase) UpdateUser(user *models.Users) (*models.Users, error) {
	err := u.userRepository.UpdateUser(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}
