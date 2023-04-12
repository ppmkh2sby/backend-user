package repository

import (
	"fmt"
	"github.com/ppmkh2sby/backend-library/helpers/logformat"
	"github.com/ppmkh2sby/backend-library/models"
)

func (p *postgresDB) ChangePasswordUser(user models.Users) (models.Users, error) {
	queryStr := fmt.Sprintf("UPDATE %s SET password = $1 WHERE id = $2", usersTable)

	_, err := p.db.Exec(queryStr, user.Password, user.ID)
	if err != nil {
		return models.Users{}, logformat.PostgresExecResponseError(err)
	}

	return models.Users{ID: user.ID}, err
}
