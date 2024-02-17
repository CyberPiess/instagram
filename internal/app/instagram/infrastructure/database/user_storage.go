package database

import (
	"database/sql"

	"github.com/CyberPiess/instagram/internal/app/instagram/domain/user"
	sq "github.com/Masterminds/squirrel"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (ur *UserRepository) Insert(newUser user.User) error {

	query := sq.Insert("user_table").
		Columns("username", "user_email", "hashed_password", "create_time").
		Values(newUser.Username, newUser.User_email, newUser.Password, newUser.Create_time).
		RunWith(ur.db).
		PlaceholderFormat(sq.Dollar)
	_, err := query.Query()

	return err
}
