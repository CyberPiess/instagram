package instagram

import (
	"database/sql"

	_ "github.com/lib/pq"
)

var db *sql.DB

// OpenDB открывает соединение с базой данных и возвращает объект *sql.DB
func OpenDB() (*sql.DB, error) {
	if db != nil {
		return db, nil
	}

	db, err := sql.Open("postgres", "host=localhost port=5432 user=admin password=password dbname=Instagram sslmode=disable")
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
