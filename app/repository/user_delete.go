package repository

import (
	"fmt"
	"github.com/ppmkh2sby/backend-library/helpers/logformat"
)

// DeleteUser is funtion to delete user from database
func (p *postgresDB) DeleteUser(id string) error {
	queryStr := fmt.Sprintf("DELETE FROM %s WHERE id = $1", usersTable)

	_, err := p.db.Exec(queryStr, id)
	if err != nil {
		return logformat.PostgresExecResponseError(err)
	}

	return nil
}