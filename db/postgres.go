package db

import (
	"database/sql"
	"fmt"
)

// PostgreDB structure
type PostgresDB struct {
	DB *sql.DB
}

// InitPostgreDB create PostgreDB instance
func InitPostgresDB(driver, dbserver, dbname, dbuser, dbpass string, dbport, timeout int) (*PostgresDB, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s connect_timeout=%d sslmode=disble",
		dbserver, dbport, dbuser, dbpass, dbname, timeout)
	db, err := sql.Open(driver, psqlInfo)
	if err != nil {
		return nil, err
	}

	postgresDB := &PostgresDB{
		DB: db,
	}

	return postgresDB, nil
}
