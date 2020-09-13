package model

import "database/sql"

type HandlerFunctions interface {
	GetUser()
}

type DB struct {
	*sql.DB
}
