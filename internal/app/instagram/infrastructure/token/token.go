package token

import (
	"fmt"

	"github.com/golang-jwt/jwt/v5"
)

type Credentials struct {
	UserId string
	jwt.RegisteredClaims
}

type token struct {
}

func NewToken() *token {
	return &token{}
}

const (
	secretKey = "secret"
)

func (t *token) VerifyToken(tokenString string) (*Credentials, error) {
	var jwtClaims Credentials
	token, err := jwt.ParseWithClaims(tokenString, &jwtClaims, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})

	if err != nil {
		return &Credentials{}, err
	}

	if !token.Valid {
		return &Credentials{}, fmt.Errorf("invalid token")
	}

	return &jwtClaims, nil
}
