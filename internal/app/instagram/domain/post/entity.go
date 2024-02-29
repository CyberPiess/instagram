package post

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Post struct {
	PostImage       string `json:"image"`
	PostDescription string `json:"description"`
	CreateTime      time.Time
	UserId          int
	AccessToken     string
}

type CreatePost struct {
	PostImage       string
	PostDescription string
	CreateTime      time.Time
	UserId          int
}

type Credentials struct {
	UserId string
	jwt.RegisteredClaims
}
