package instagram

import (
	"database/sql"
	"time"
)

func ifUserExist(username string, db *sql.DB) bool {
	row := db.QueryRow("SELECT username FROM users WHERE username = $1", username)
	err := row.Scan(username)
	return err != sql.ErrNoRows
}

func userCreate(username string, user_email string, hashed_password string, create_time time.Time, db *sql.DB) bool {
	_, err := db.Query("INSERT INTO Users (username, user_email, hashed_password, create_time) VALUES ($1, $2, $3, $4)", username, user_email, hashed_password, create_time)

	return err == nil
}

func userDelete(username string, db *sql.DB) int {
	_, err := db.Query("DELETE FROM \"Users\" WHERE username = $1", username)
	if err != nil {
		return 500
	}
	return 200
}

func userUpdate(username string, password string, newUsername string, newPassword string, db *sql.DB) int {
	login_result, _ := userLogin(username, password, db)
	if login_result != 200 {
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
