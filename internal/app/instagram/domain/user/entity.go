package user

import (
	"time"

	database "github.com/CyberPiess/instagram/internal/app/instagram/infrastructure/database"
)

type User struct {
	Username    string
	User_email  string
	Password    string
	Create_time time.Time
	DB          database.Env
}
