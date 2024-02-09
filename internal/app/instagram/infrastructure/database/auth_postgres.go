package database

import (
	"database/sql"

	domain "github.com/CyberPiess/instagram/internal/app/instagram/domain/user"

	sq "github.com/Masterminds/squirrel"
)

type AuthPostgres struct {
	db *sql.DB
}

func NewAuthPostgres(db *sql.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (a *AuthPostgres) CreateUser(user domain.User) (int, error) {
	var id int
	query := sq.Insert("users").
		Columns("username", "user_email", "hashed_password", "create_time").
		Values(username, user_email, hashed_password, create_time).
		RunWith(a.db).
		PlaceholderFormat(sq.Dollar)
	row, err := query.Query()

	if err != nil {
		return 0, err
	}

	if err = row.Scan(&id); err != nil {
		return 0, err
	}
	return id, err
}
