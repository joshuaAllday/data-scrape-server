package models

import (
	"errors"
	"net/mail"
	"strings"
)

func validatePassword(password string, checkPassword string) error {
	if password == checkPassword {
		return nil
	}

	return errors.New("Passwords don't match")
}

func IsLower(s string) bool {
	return strings.ToLower(s) == s
}

func validEmail(email string) bool {
	if !IsLower(email) {
		return false
	}

	if addr, err := mail.ParseAddress(email); err != nil {
		return false
	} else if addr.Name != "" {
		// mail.ParseAddress accepts input of the form "Billy Bob <billy@example.com>" which we don't allow
		return false
	}

	return true

}

func SantizeEmail(email string) error {

	if validEmail(email) {
		return nil
	}
	return errors.New("Invalid email")
}
