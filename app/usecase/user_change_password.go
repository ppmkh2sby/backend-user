package usecase

import (
	"context"
	"github.com/ppmkh2sby/backend-library/helpers/generate"
)

// ChangePasswordUser is function on usecase to handle change password on database
func (u *userUsecase) ChangePasswordUser(ctx context.Context, id, password string) error {
	password, err := generate.HashPassword(password)
	if err != nil {
		return err
	}

	err = u.userRepository.ChangePasswordUser(ctx, id, password)
	if err != nil {
		return err
	}

	return nil
}
