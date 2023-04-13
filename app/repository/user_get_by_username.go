package repository

import (
	"fmt"
	"github.com/ppmkh2sby/backend-library/helpers/logformat"
	"github.com/ppmkh2sby/backend-library/models"
)

// GetUserByUsername is function to get user by username for sign in
func (p *postgresDB) GetUserByUsername(username string) (*models.Users, error) {
	var user models.Users

	queryStr := fmt.Sprintf("SELECT id, username, password FROM %s WHERE username = $1", usersTable)

	result, err := p.db.Query(queryStr, username)
	if err != nil {
		return nil, logformat.PostgresQueryResponseError(err)
	}

	for result.Next() {
		err = result.Scan(&user.ID, &user.Username, &user.Password)
		if err != nil {
			return nil, logformat.PostgresScanError(err)
		}
	}

	return &user, nil
}
