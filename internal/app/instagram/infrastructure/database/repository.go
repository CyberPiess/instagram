package database

import (
	"database/sql"
	"time"
)

type Authorization interface {
	CreateUser(username, user_email, hashed_password string, create_time time.Time) (int, error)
}

type Repository struct {
	Authorization
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
	}
}
