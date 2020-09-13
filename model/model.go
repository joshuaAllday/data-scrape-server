package model

import "database/sql"

type HandlerFunctions interface {
	GetUser()
	GetHealth() error
}

type DB struct {
	*sql.DB
}
