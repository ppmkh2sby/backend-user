package usecase

import "github.com/ppmkh2sby/backend-library/models"

// GetUserByID is function to get specific user by id
func (u *userUsecase) GetUserByID(id string) (*models.Users, error) {
	user, err := u.userRepository.GetUserByID(id)
	if err != nil {
		return nil, err
	}

	return user, nil
}
