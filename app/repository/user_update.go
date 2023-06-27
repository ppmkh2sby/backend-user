package repository

import (
	"context"
	"fmt"
	"github.com/ppmkh2sby/backend-library/helpers/logformat"
	"github.com/ppmkh2sby/backend-library/models"
	"time"
)

// UpdateUser is function on repository to handle update user by id
func (p *PostgresDB) UpdateUser(ctx context.Context, user *models.Users) error {
	queryStr := fmt.Sprintf("UPDATE %s SET username = $1, email = $2, updated_at = $3 WHERE id = $4", usersTable)

	_, err := p.DB.Exec(queryStr, user.Username, user.Email, time.Now(), user.ID)
	if err != nil {
		return logformat.PostgresExecResponseError(err)
	}

	return nil
}
