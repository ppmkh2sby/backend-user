package repository

import (
	"context"
	"fmt"
	"github.com/ppmkh2sby/backend-library/helpers/logformat"
	"github.com/ppmkh2sby/backend-library/models"
)

// GetUserByID is function on repository for get user specific by id user
func (p *PostgresDB) GetUserByID(ctx context.Context, id string) (*models.Users, error) {
	var user models.Users

	queryStr := "SELECT id, username, email, role, card_id, santri_id, created_at, updated_at "
	queryStr += fmt.Sprintf("FROM %s WHERE id = $1", usersTable)

	result, err := p.DB.Query(queryStr, id)
	if err != nil {
		return nil, logformat.PostgresQueryResponseError(err)
	}

	for result.Next() {
		err = result.Scan(&user.ID, &user.Username, &user.Email, &user.Role, &user.CardID, &user.SantriID, &user.CreatedAt, &user.UpdatedAt)
		if err != nil {
			return nil, logformat.PostgresScanError(err)
		}
	}

	return &user, nil
}
