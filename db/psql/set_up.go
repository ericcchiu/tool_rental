package psql

import (
	"database/sql"
	"fmt"
)

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

	defer db.Close()
	return
}
