package user

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type User struct {
	Username   string
	UserEmail  string
	Password   string
	CreateTime time.Time
}

type LoginUserReq struct {
	Username string
	Password string
}

type LoginUserRes struct {
	AccessToken string
	UserId      string
	Username    string
}

type MyJWTClaims struct {
	UserId string
	jwt.RegisteredClaims
}
