package usecase

import (
	"context"
)

// DeleteUser is function on usecase to handle delete user from database
func (u *userUsecase) DeleteUser(ctx context.Context, id string) error {
	err := u.userRepository.DeleteUser(ctx, id)
	if err != nil {
		return err
	}

	return nil
}
