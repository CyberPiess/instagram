package instagram

import (
	"database/sql"
)

func ifUserExist(username string, db *sql.DB) *sql.Row {
	row := db.QueryRow("SELECT username, password FROM \"Users\" WHERE username = $1", username)
	return row
}

func userCreate(username string, password string, db *sql.DB) int {
	_, err := db.Query("INSERT INTO \"Users\" (username, password) VALUES ($1, $2)", username, password)

	if err != nil {
		return 500
	}
	return 200
}

func userDelete(username string, db *sql.DB) int {
	_, err := db.Query("DELETE FROM \"Users\" WHERE username = $1", username)
	if err != nil {
		return 500
	}
	return 200
}

func userUpdate(username string, password string, newUsername string, newPassword string, db *sql.DB) int {
	loginResult, _ := userLogin(username, password, db)
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

func userLogin(username string, password string, db *sql.DB) (int, int) {
	result, err := db.Query("SELECT user_id FROM \"Users\" WHERE username = $1 and password=$2", username, password)
	if err != nil {
		return 500, -1
	}

	var user_id int
	for result.Next() {
		err = result.Scan(&user_id)
		if err != nil {
			panic(err)
		}
	}

	return 200, user_id
}
