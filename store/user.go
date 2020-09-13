package store

import (
	"database/sql"
)

var (
	insertUser   = `INSERT INTO Users(Email, Password) VALUES($1, $2)`
	getUserLogin = `SELECT Users.ID, Users.Email, Users.Password FROM Users WHERE Email = $1`
	getUserInfo  = `Select Users.ID, Users.Email FROM Users WHERE Email = $1`
)

func (db *DB) GetUser() {

}

func (db *DB) CreateUser(email string, password string) (*bool, error) {
	stmt, _ := db.Prepare(insertUser)
	_, err := stmt.Exec(email, password)
	valid := true
	if err != nil {
		return nil, err
	}

	return &valid, nil
}

type UserDetails struct {
	ID       string
	Email    string
	Password string
}

func (db *DB) LoginUser(email string) (*UserDetails, error) {
	row := db.QueryRow(getUserLogin, email)
	userStruct := new(UserDetails)
	err := row.Scan(
		&userStruct.ID,
		&userStruct.Email,
		&userStruct.Password,
	)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}

	return userStruct, nil
}

type UserInfo struct {
	ID    string
	Email string
}

func (db *DB) FetchUserInfo(email string) (*UserInfo, error) {
	row := db.QueryRow(getUserInfo, email)
	userStruct := new(UserInfo)
	err := row.Scan(
		&userStruct.ID,
		&userStruct.Email,
	)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}

	return userStruct, nil
}
