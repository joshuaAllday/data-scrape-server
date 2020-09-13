package db

import (
	"database/sql"

	store "github.com/data-scrape/data-scrape-server/store"
	_ "github.com/lib/pq"
)

func SetupDatabase(connectionString string) (*store.DB, error) {
	db, err := sql.Open("postgres", connectionString+"?sslmode=disable")
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return &store.DB{db}, nil
}
