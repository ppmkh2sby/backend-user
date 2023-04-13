package repository

import (
	"database/sql"
	"github.com/ppmkh2sby/backend-library/models"
)

const (
	usersTable = "users"
)

type postgresDB struct {
	db *sql.DB
}

type UserRepository interface {
	GetAllUsers() ([]models.Users, error)
	GetUserByID(id string) (*models.Users, error)
	GetUserByUsername(username string) (*models.Users, error)
	CreateUser(user *models.Users) (*models.Users, error)
	UpdateUser(user *models.Users) error
	DeleteUser(id string) error
	ChangePasswordUser(id, password string) error
}

// NewUserRepository is function to initial user repository
func NewUserRepository(db *sql.DB) *postgresDB {
	return &postgresDB{db: db}
}
