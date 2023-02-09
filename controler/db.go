package controler

import "database/sql"

type DB struct {
	db *sql.DB
}

func (db *DB) Open() error {
	db, err := sql.Open()
}
