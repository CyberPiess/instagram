package user

import "time"

type UserDAO struct {
	Username   string
	UserEmail  string
	Password   string
	CreateTime time.Time
}
