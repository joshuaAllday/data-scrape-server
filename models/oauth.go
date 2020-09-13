package models

import (
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type Oauth struct {
	Token        string `json:"token"`
	RefreshToken string `json:"refresh_token`
	Expires      int64  `json:"expires_in"`
	Role         string `json:"user"`
}

func CreateJwt(id string) (*Oauth, error) {
	expires := time.Now().Add(time.Minute * 15).Unix()
	token, _ := createToken(id, expires, false)
	refreshToken, _ := createToken(id, time.Now().Add(time.Minute*1000).Unix(), true)
	oauth := Oauth{
		Token:        *token,
		RefreshToken: *refreshToken,
		Expires:      expires,
		Role:         "User",
	}
	return &oauth, nil
}

func createToken(id string, exp int64, refresh bool) (*string, error) {
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
