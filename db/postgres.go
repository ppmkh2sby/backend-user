package db

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

const (
	// PostgresDriver is default name "postgres"
	PostgresDriver = "postgres"
)

// InitPostgreDB create PostgreDB instance
func InitPostgresDB(driver, dbserver, dbname, dbuser, dbpass string, dbport, timeout int) (*sql.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s connect_timeout=%d sslmode=disble",
		dbserver, dbport, dbuser, dbpass, dbname, timeout)
	db, err := sql.Open(driver, psqlInfo)
	if err != nil {
		return nil, err
	}

	return db, nil
}
