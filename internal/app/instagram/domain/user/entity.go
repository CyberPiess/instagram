package user

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	Id              uuid.UUID
	Username        string
	User_email      string
	Hashed_password string
	Create_time     time.Time
}
