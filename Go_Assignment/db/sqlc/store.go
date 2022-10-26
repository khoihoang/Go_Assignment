package db

import (
	"database/sql"
)

type Store struct {
	Queries *Queries
	Db      *sql.DB
}

func InitalStore(db *sql.DB) *Store {
	return &Store{
		Db:      db,
		Queries: New(db),
	}
}
