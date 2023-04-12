package repository

import (
	"fmt"
	"github.com/ppmkh2sby/backend-library/helpers/logformat"
	"github.com/ppmkh2sby/backend-library/models"
)

// GetUserByID is function for get user specific by id user
func (p *postgresDB) GetUserByID(id string) (*models.Users, error) {
	var user *models.Users

	queryStr := "SELECT id, username, email, card_id, santri_id, created_at, updated_at "
	queryStr += fmt.Sprintf("FROM %s WHERE id = %s", usersTable, id)

	result, err := p.db.Query(queryStr)
	if err != nil {
		return nil, logformat.PostgresQueryResponseError(err)
	}

	for result.Next() {
		err = result.Scan(&user.ID, &user.Username, &user.Email, &user.CardID, &user.SantriID, &user.CreatedAt, &user.UpdatedAt)
	}

	return user, nil
}
