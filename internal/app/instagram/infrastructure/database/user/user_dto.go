package user

import "time"

type UserDTO struct {
	Username   string
	UserEmail  string
	Password   string
	CreateTime time.Time
}
