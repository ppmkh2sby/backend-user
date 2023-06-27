package repository

import (
	"context"
	"fmt"
	"github.com/ppmkh2sby/backend-library/helpers/logformat"
)

// DeleteUser is function on repository to handle delete user from database
func (p *PostgresDB) DeleteUser(ctx context.Context, id string) error {
	queryStr := fmt.Sprintf("DELETE FROM %s WHERE id = $1", usersTable)

	_, err := p.DB.Exec(queryStr, id)
	if err != nil {
		return logformat.PostgresExecResponseError(err)
	}

	return nil
}
