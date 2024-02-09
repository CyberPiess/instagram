package user

import (
	_ "database/sql"
	"time"

	"github.com/google/uuid"

	storage "github.com/CyberPiess/instagram/internal/app/instagram/infrastructure/user"
)

func Create(new_user User) error {

	if new_user.Username == "" || new_user.User_email == "" || new_user.Password == "" {
		return nil
	}

	hashed_password := hashAndSalt([]byte(new_user.Password))

	db_user := &storage.User{
		Id:              uuid.New(),
		Username:        new_user.Username,
		User_email:      new_user.User_email,
		Hashed_password: hashed_password,
		Create_time:     time.Now(),
	}

	user_reposityory := &storage.UserRepository{}
	err := user_reposityory.Create(db_user)
	return err
}
