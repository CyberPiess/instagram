package user

import (
	"time"

	"github.com/google/uuid"

	sq "github.com/Masterminds/squirrel"
)

type User struct {
	Id              uuid.UUID
	Username        string
	User_email      string
	Hashed_password string
	Create_time     time.Time
}

type UserRepository struct {
}

func (ur *UserRepository) ifUsernameExist(username string) error {
	psql := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)

	query, args, err := psql.Select("username").From("users").Where("username = ?", username).ToSql()
	if err != nil {
		return err
	}
	row := db.QueryRow(query, args...)
	err = row.Scan(&username)
	return err
}

func (ur *UserRepository) ifEmailExist(user_email string) error {
	psql := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)

	query, args, err := psql.Select("user_email").From("users").Where("username = ?", user_email).ToSql()
	if err != nil {
		return err
	}
	row := db.QueryRow(query, args...)
	err = row.Scan(&user_email)
	return err
}

func (ur *UserRepository) Create(new_user *User) error {

	query := sq.Insert("users").
		Columns("username", "user_email", "hashed_password", "create_time").
		Values(new_user.Username, new_user.User_email, new_user.Hashed_password, new_user.Create_time).
		RunWith(db).
		PlaceholderFormat(sq.Dollar)
	_, err := query.Query()

	return err
}
