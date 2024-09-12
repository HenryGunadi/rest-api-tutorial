package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func NewPostgreSQL(conn string) (*sql.DB, error) {
	db, err := sql.Open("postgres", conn)
	if err != nil {
		return nil, err
	}

	return db, nil
}