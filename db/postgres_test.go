package db_test

import (
	"github.com/ppmkh2sby/backend-user/db"
	"testing"
)

func TestInitPostgresDB(t *testing.T) {
	driver := db.PostgresDriver
	dbserver := "localhost"
	dbname := "rfid_db"
	dbuser := "ppmkh2sby"
	dbpass := "ppmkh2sby"
	dbport := 5432
	timeout := 5

	t.Run("Create postgres OK", func(t *testing.T) {
		_, err := db.InitPostgresDB(driver, dbserver, dbname, dbuser, dbpass, dbport, timeout)
		if err != nil {
			t.Fatalf("Create connection postgres shouldn't be error")
		}
	})

	t.Run("Create postgres NOK", func(t *testing.T) {
		driver = "postgre"
		_, err := db.InitPostgresDB(driver, dbserver, dbname, dbuser, dbpass, dbport, timeout)
		if err == nil {
			t.Fatalf("Create connection postgres should be error")
		}
	})
}
