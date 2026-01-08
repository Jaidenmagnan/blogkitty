// This package initializes the db.
package db

import (
	"database/sql"
	"github.com/charmbracelet/log"

	_ "github.com/mattn/go-sqlite3"
)

// The global DB connection.
var DB *sql.DB

// Connects to the database.
func Connect() error {
	var err error
	DB, err = sql.Open("sqlite3", "./app.db")
	if err != nil {
		return err
	}

	if err = DB.Ping(); err != nil {
		return err
	}

	log.Info("Database connected successfully")
	return nil
}
