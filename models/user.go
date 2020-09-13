package models

import (
	"encoding/json"
	"io"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Email         string
	Password      string
	CheckPassword string
}

func UserFromJson(data io.Reader) (*User, error) {
	var user *User
	json.NewDecoder(data).Decode(&user)

	err := validatePassword(user.Password, user.CheckPassword)
	if err != nil {
		return nil, err
	}

	err = santizeEmail(user.Email)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func HashPassword(password string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		panic(err)
	}

	return string(hash)
}
