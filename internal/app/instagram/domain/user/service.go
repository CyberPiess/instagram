package user

import storage "github.com/CyberPiess/instagram/internal/app/instagram/infrastructure/user"

type Repository interface {
	Create(new_user *storage.User) error
	ifUsernameExist(username string) error
	ifUserEmailExist(user_email string) error
}
