//go:generate mockgen -source=user_storage.go -destination=mocks/mock_storage.go
package user

import (
	"database/sql"

	sq "github.com/Masterminds/squirrel"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (ur *UserRepository) Insert(newUser UserDTO) error {

	query := sq.Insert("public.user").
		Columns("username", "user_email", "hashed_password", "create_time").
		Values(newUser.Username, newUser.UserEmail, newUser.Password, newUser.CreateTime).
		RunWith(ur.db).
		PlaceholderFormat(sq.Dollar)
	_, err := query.Query()

	return err
}

func (ur *UserRepository) IfUserExist(newUser UserDTO) (bool, error) {
	var exists bool
	psql := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)

	query, args, err := psql.Select("username").Prefix("SELECT EXISTS (").From("public.user").Where("username = ?", newUser.Username).Suffix(")").ToSql()

	if err != nil {
		return false, err
	}

	row := ur.db.QueryRow(query, args...)
	err = row.Scan(&exists)
	return exists, err
}

func (ur *UserRepository) IfEmailExist(newUser UserDTO) (bool, error) {
	var exists bool
	psql := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)

	query, args, err := psql.Select("user_email").Prefix("SELECT EXISTS (").From("public.user").Where("user_email = ?", newUser.UserEmail).Suffix(")").ToSql()

	if err != nil {
		return false, err
	}

	row := ur.db.QueryRow(query, args...)
	err = row.Scan(&exists)
	return exists, err
}

func (ur *UserRepository) SelectUser(username string) (int, string, error) {
	var hashed_password string
	var user_id int
	psql := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)

	query, args, err := psql.Select("user_id, hashed_password").From("public.user").Where("username = ?", username).ToSql()
	if err != nil {
		return -1, "", err
	}

	row := ur.db.QueryRow(query, args...)
	err = row.Scan(&user_id, &hashed_password)
	if err != nil {
		return -1, "", err
	}

	return user_id, hashed_password, nil
}
