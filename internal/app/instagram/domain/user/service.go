package user

import database "github.com/CyberPiess/instagram/internal/app/instagram/infrastructure/database"

type userStorage interface {
	insert(db_user database.UserRepository) error
}

type UserService struct {
	store userStorage
}

func (u *UserService) CreateUser(new_user User) error {
	if new_user.Username == "" || new_user.User_email == "" || new_user.Password == "" {
		return nil
	}

	hashed_password := hashAndSalt([]byte(new_user.Password))

	db_user := database.UserRepository{
		Username:    new_user.Username,
		User_email:  new_user.User_email,
		Password:    hashed_password,
		Create_time: new_user.Create_time,
	}

	err := db_user.Insert(&new_user.DB)

	return err
}
