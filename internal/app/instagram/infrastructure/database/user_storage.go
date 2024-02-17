package database

import (
	"time"

	sq "github.com/Masterminds/squirrel"
)

type UserRepository struct {
	Username    string
	User_email  string
	Password    string
	Create_time time.Time
}

func (ur *UserRepository) Insert(env *Env) error {

	query := sq.Insert("user_table").
		Columns("username", "user_email", "hashed_password", "create_time").
		Values(ur.Username, ur.User_email, ur.Password, ur.Create_time).
		RunWith(env.db).
		PlaceholderFormat(sq.Dollar)
	_, err := query.Query()

	return err
}
