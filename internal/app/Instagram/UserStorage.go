package instagram

import (
	"database/sql"
)

func ifExist(username string, db *sql.DB) *sql.Row {
	row := db.QueryRow("SELECT username, password FROM \"Users\" WHERE username = $1", username)
	return row
}

func create(username string, password string, db *sql.DB) int {
	_, err := db.Exec("INSERT INTO \"Users\" (username, password) ) VALUES ($1, $2)", username, password)
	if err != nil {
		return 500
	}
	return 200
}
