package psql

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type toolDataStore struct {
	db *sql.DB
}

func NewPostgresToolDataStore(db *sql.DB) *toolDataStore {
	return &toolDataStore{
		db,
	}
}
