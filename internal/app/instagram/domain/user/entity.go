package user

import (
	"time"
)

type User struct {
	Username    string
	User_email  string
	Password    string
	Create_time time.Time
}
