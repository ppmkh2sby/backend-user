package usecase

import "github.com/ppmkh2sby/backend-library/models"

// GetAllUsers is fucntion to get all users from database
func (u *userUsecase) GetAllUsers() ([]models.Users, error) {
	users, err := u.userRepository.GetAllUsers()
	if err != nil {
		return nil, err
	}

	return users, nil
}
