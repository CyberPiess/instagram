package user

import (
	"database/sql"
)

type Repository interface {
	Create(user *User, db *sql.DB) error
	IfUsernameExist(username string, db *sql.DB) error
	IfEmailExist(user_email string, db *sql.DB) error
}
