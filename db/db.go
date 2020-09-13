package db

import (
	"database/sql"

	"github.com/data-scrape/data-scrape-server/model"
	_ "github.com/lib/pq"
)

func SetupDatabase(connectionString string) (*model.DB, error) {
	db, err := sql.Open("postgres", connectionString+"?sslmode=disable")
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return &model.DB{db}, nil
}
