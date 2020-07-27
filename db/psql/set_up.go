package psql

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

// NewPostgresConnection takes in a connection string for PostgreSQL database and returns a connection to the database
func NewPostgresConnection(connection string) (db *sql.DB, err error) {
	// Open a database. Prepares the database abstraction for later use.
	db, err = sql.Open("postgres", connection)
	if err != nil {
		return nil, err
	}

	// Check connection: checks if database is available and accessible
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	fmt.Println("Ping successful")

	return
}
