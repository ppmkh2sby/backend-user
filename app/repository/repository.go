package repository

import (
	"context"
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
	GetAllUsers(ctx context.Context) ([]models.Users, error)
	GetUserByID(ctx context.Context, id string) (*models.Users, error)
	GetUserByUsername(ctx context.Context, username string) (*models.Users, error)
	CreateUser(ctx context.Context, user *models.Users) (*models.Users, error)
	UpdateUser(ctx context.Context, user *models.Users) error
	DeleteUser(ctx context.Context, id string) error
	ChangePasswordUser(ctx context.Context, id, password string) error
}

// NewUserRepository is function to initial user repository
func NewUserRepository(db *sql.DB) *postgresDB {
	return &postgresDB{db: db}
}
