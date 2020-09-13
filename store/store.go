package store

import "database/sql"

type HandlerFunctions interface {
	GetUser()
	GetHealth() error
	CreateUser(email string, password string) (*bool, error)
	LoginUser(email string) (*UserDetails, error)
}

type DB struct {
	*sql.DB
}
