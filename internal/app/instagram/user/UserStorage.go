package instagram

import (
	"database/sql"

	sq "github.com/Masterminds/squirrel"
)

func ifUserExist(username string, db *sql.DB) error {
	psql := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)

	query, args, err := psql.Select("username").From("users").Where("username = ?", username).ToSql()
	if err != nil {
		return err
	}
	row := db.QueryRow(query, args...)
	err = row.Scan(&username)
	return err
}

func ifEmailExist(user_email string, db *sql.DB) error {
	psql := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)

	query, args, err := psql.Select("user_email").From("users").Where("username = ?", user_email).ToSql()
	if err != nil {
		return err
	}
	row := db.QueryRow(query, args...)
	err = row.Scan(&user_email)
	return err
}

func userCreate(new_user User, db *sql.DB) bool {

	query := sq.Insert("users").
		Columns("username", "user_email", "hashed_password", "create_time").
		Values(new_user.username, new_user.user_email, new_user.hashed_password, new_user.create_time).
		RunWith(db).
		PlaceholderFormat(sq.Dollar)
	_, err := query.Query()

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
