package models

import (
	"encoding/json"
	"errors"
	"io"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Email         string
	Password      *string
	CheckPassword *string
}

func UserFromJson(data io.Reader) *User {
	var user *User
	json.NewDecoder(data).Decode(&user)

	return user
}

func (user *User) SanitizeUserRegister() error {

	if user.CheckPassword == nil {
		return errors.New("Passwords don't match")
	}
	err := validatePassword(*user.Password, *user.CheckPassword)
	if err != nil {
		return err
	}

	err = SantizeEmail(user.Email)
	if err != nil {
		return err
	}

	return nil
}

func HashPassword(password string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		panic(err)
	}

	return string(hash)
}

func (user *User) SanitizeUserLogin() error {
	err := SantizeEmail(user.Email)

	if err != nil {
		return err
	}

	return nil
}

func CheckHashPasswords(password string, dbPassword string) bool {

	encryptionErr := bcrypt.CompareHashAndPassword([]byte(dbPassword), []byte(password))

	if encryptionErr != nil {
		return false
	}

	return true
}
