package repository

import (
	"context"
	"database/sql"
	"errors"
	"github.com/ppmkh2sby/backend-library/helpers/logformat"
	"github.com/ppmkh2sby/backend-library/models"
)

// GetAllUsers is function on repository for get all user
func (p *PostgresDB) GetAllUsers(ctx context.Context) ([]models.Users, error) {
	var users []models.Users

	queryStr := "SELECT id, username, email, role, card_id, santri_id, created_at, updated_at "
	queryStr += "FROM users GROUP BY id, santri_id ORDER BY santri_id ASC"

	result, err := p.DB.Query(queryStr)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return users, nil
		}
		return nil, logformat.PostgresQueryResponseError(err)
	}

	for result.Next() {
		user := models.Users{}
		err = result.Scan(&user.ID, &user.Username, &user.Email, &user.Role, &user.CardID, &user.SantriID, &user.CreatedAt, &user.UpdatedAt)
		if err != nil {
			return nil, logformat.PostgresScanError(err)
		}
		users = append(users, user)
	}

	return users, nil
}
