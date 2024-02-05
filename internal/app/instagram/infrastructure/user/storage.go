package user

import (
	"database/sql"

	"github.com/CyberPiess/instagram/internal/app/instagram/domain/user"

	sq "github.com/Masterminds/squirrel"
)

type UserRepository struct {
}

func (ur *UserRepository) IfUsernameExist(username string, db *sql.DB) error {
	psql := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)

	query, args, err := psql.Select("username").From("users").Where("username = ?", username).ToSql()
	if err != nil {
		return err
	}
	row := db.QueryRow(query, args...)
	err = row.Scan(&username)
	return err
}

func (ur *UserRepository) IfEmailExist(user_email string, db *sql.DB) error {
	psql := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)

	query, args, err := psql.Select("user_email").From("users").Where("username = ?", user_email).ToSql()
	if err != nil {
		return err
	}
	row := db.QueryRow(query, args...)
	err = row.Scan(&user_email)
	return err
}

func (ur *UserRepository) Create(user user.User, db *sql.DB) error {

	query := sq.Insert("users").
		Columns("username", "user_email", "hashed_password", "create_time").
		Values(user.Username, user.User_email, user.Hashed_password, user.Create_time).
		RunWith(db).
		PlaceholderFormat(sq.Dollar)
	_, err := query.Query()

	return err
}
