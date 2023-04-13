package usecase

// DeleteUser is function to delete user from database
func (u *userUsecase) DeleteUser(id string) error {
	err := u.userRepository.DeleteUser(id)
	if err != nil {
		return err
	}

	return nil
}
