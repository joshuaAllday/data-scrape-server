package models

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type Oauth struct {
	Token        string `json:"token"`
	RefreshToken string `json:"refresh_token"`
	Expires      int64  `json:"expires_in"`
	Role         string `json:"role"`
}

func CreateJwt(id string) (*Oauth, error) {
	expires := time.Now().Add(time.Minute * 15).Unix()
	token, _ := encodeToken(id, expires, false)
	refreshToken, _ := encodeToken(id, time.Now().Add(time.Minute*1000).Unix(), true)
	oauth := Oauth{
		Token:        *token,
		RefreshToken: *refreshToken,
		Expires:      expires,
		Role:         "User",
	}
	return &oauth, nil
}

func encodeToken(id string, exp int64, refresh bool) (*string, error) {
	signing := os.Getenv("AUTH_SIGNING_KEY")
	mySigningKeyAuthor := []byte(signing)
	token := jwt.New(jwt.SigningMethodHS512)
	claims := token.Claims.(jwt.MapClaims)

	claims["id"] = id
	claims["exp"] = exp
	claims["refresh"] = refresh

	tokenString, err := token.SignedString(mySigningKeyAuthor)
	if err != nil {
		return nil, err
	}
	return &tokenString, nil
}

func decodeToken(token string) (*string, error) {
	decodedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("There is an error")
		}
		signing := os.Getenv("AUTH_SIGNING_KEY")
		key := []byte(signing)
		return key, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := decodedToken.Claims.(jwt.MapClaims); ok && decodedToken.Valid {
		id := claims["id"].(string)
		return &id, nil
	}

	return nil, errors.New("Error decoding token")
}

func CreateTokenRefresh(token string) (*Oauth, *string, error) {
	id, err := decodeToken(token)

	if err != nil {
		return nil, nil, err
	}
	expires := time.Now().Add(time.Minute * 15).Unix()
	newToken, _ := encodeToken(*id, expires, false)

	return &Oauth{
		Token:        *newToken,
		RefreshToken: token,
		Expires:      expires,
		Role:         "User",
	}, id, nil
}

func GetTokenId(token string) (*string, error) {
	id, err := decodeToken(token)

	if err != nil {
		return nil, err
	}

	return id, nil
}

func ValidToken(token string) bool {
	decodedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("There is an error")
		}
		signing := os.Getenv("AUTH_SIGNING_KEY")
		key := []byte(signing)
		return key, nil
	})

	if err != nil {
		return false
	}

	if decodedToken.Valid {
		return true
	}

	return false
}
