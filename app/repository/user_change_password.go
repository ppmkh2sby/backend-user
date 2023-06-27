package repository

import (
	"context"
	"fmt"
	"github.com/ppmkh2sby/backend-library/helpers/logformat"
)

// ChangePasswordUser is function on repository to handle change user's password
func (p *PostgresDB) ChangePasswordUser(ctx context.Context, id, password string) error {
	queryStr := fmt.Sprintf("UPDATE %s SET password = $1 WHERE id = $2", usersTable)

	_, err := p.DB.Exec(queryStr, password, id)
	if err != nil {
		return logformat.PostgresExecResponseError(err)
	}

	return nil
}
