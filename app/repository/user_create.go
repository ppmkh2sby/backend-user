package repository

import (
	"fmt"
	"github.com/ppmkh2sby/backend-library/helpers/logformat"
	"github.com/ppmkh2sby/backend-library/models"
	"time"
)

// CreateUser is function for saving new user to database
func (p *postgresDB) CreateUser(user *models.Users) (*models.Users, error) {
	queryStr := fmt.Sprintf("INSERT INTO %s (id, username, password, email, role, card_id, santri_id, created_at, updated_at)", usersTable)
	queryStr += "VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9)"

	_, err := p.db.Exec(queryStr, user.ID, user.Username, user.Password, user.Email, user.Role, user.CardID, user.SantriID, time.Now(), time.Now())
	if err != nil {
		return nil, logformat.PostgresExecResponseError(err)
	}

	return user, nil
}
