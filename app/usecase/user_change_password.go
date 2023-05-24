package usecase

import (
	"github.com/ppmkh2sby/backend-library/helpers/generate"
)

// ChangePasswordUser is function on usecase to handle change password on database
func (u *userUsecase) ChangePasswordUser(id, password string) error {
	password, err := generate.HashPassword(password)
	if err != nil {
		return err
	}

	err = u.ChangePasswordUser(id, password)
	if err != nil {
		return err
	}

	return nil
}
