package instagram

import (
	"database/sql"
)

func ifExist(username string, db *sql.DB) *sql.Row {
	row := db.QueryRow("SELECT username, password FROM \"Users\" WHERE username = $1", username)
	return row
}

func create(username string, password string, db *sql.DB) int {
	_, err := db.Query("INSERT INTO \"Users\" (username, password) VALUES ($1, $2)", username, password)

	if err != nil {
		return 500
	}
	return 200
}

func delete(username string, db *sql.DB) int {
	_, err := db.Query("DELETE FROM \"Users\" WHERE username = $1", username)
	if err != nil {
		return 500
	}
	return 200
}

func update(username string, password string, newUsername string, newPassword string, db *sql.DB) int {
	loginResult := login(username, password, db)
	if loginResult != 200 {
		return 500
	}

	if username != "" && password == "" {
		_, err := db.Query("UPDATE \"Users\" SET username=$2 WHERE username = $1", username, newUsername)
		if err != nil {
			return 500
		}
		return 200
	}

	if username == "" && password != "" {
		_, err := db.Query("UPDATE \"Users\" SET password=$2 WHERE username = $1", username, newPassword)
		if err != nil {
			return 500
		}
		return 200
	}

	if username != "" && password != "" {
		_, err := db.Query("UPDATE \"Users\" SET username=$2, password=$3 WHERE username = $1", username, newUsername, newPassword)
		if err != nil {
			return 500
		}
		return 200
	}
	return 200
}

func login(username string, password string, db *sql.DB) int {
	_, err := db.Query("SELECT username, password FROM \"Users\" WHERE username = $1 and password=$2", username, password)
	if err != nil {
		return 500
	}
	return 200
}
