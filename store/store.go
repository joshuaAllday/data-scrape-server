package store

import "database/sql"

type HandlerFunctions interface {
	GetUser()
	GetHealth() error
	CreateUser(email string, password string) (*bool, error)
	LoginUser(email string) (*UserDetails, error)
	AddUserOauthToken(email string, token string) error
	DeleteAuthToken(id string) error
	FetchUserInfo(email string) (*UserInfo, error)
}

type DB struct {
	*sql.DB
}
