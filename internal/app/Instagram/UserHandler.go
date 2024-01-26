package instagram

import (
	"database/sql"
	"fmt"
	"net/http"
)

type User struct {
	username string
	password string
}

func Create(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	if r.Method != "POST" {
		http.Error(w, http.StatusText(405), 405)
		return
	}

	username := r.FormValue("username")
	password := r.FormValue("password")
	if username == "" || password == "" {
		http.Error(w, http.StatusText(400), 400)
		return
	}

	user := new(User)
	ifUserExist := ifUserExist(username, db)
	err := ifUserExist.Scan(&user.username, &user.password)
	if err != sql.ErrNoRows {
		fmt.Fprintln(w, "User with this username already exists")
		return
	}

	//TODO: поменять отдельные поля на структуру?
	createResult := userCreate(username, password, db)
	if createResult != 200 {
		http.Error(w, http.StatusText(createResult), createResult)
		return
	}

	fmt.Fprintf(w, "User %s created sucessfully\n", username)
}

func Update(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	if r.Method != "POST" {
		http.Error(w, http.StatusText(405), 405)
		return
	}

	username := r.FormValue("username")
	password := r.FormValue("password")
	newUsername := r.FormValue("newUsername")
	newPassword := r.FormValue("newPassword")
	if username == "" || password == "" {
		http.Error(w, http.StatusText(400), 400)
	}

	userExist := ifUserExist(username, db)
	user := new(User)
	err := userExist.Scan(&user.username, &user.password)
	if err == sql.ErrNoRows {
		fmt.Fprintln(w, "User with current username does not exists")
		return
	}

	userExist = ifUserExist(newUsername, db)
	err = userExist.Scan(newUsername)
	if err != sql.ErrNoRows {
		fmt.Fprintln(w, "User with new username already exists")
		return
	}

	updateResult := userUpdate(username, password, newUsername, newPassword, db)
	if updateResult != 200 {
		http.Error(w, http.StatusText(updateResult), updateResult)
		return
	}

	fmt.Fprintf(w, "User %s updated sucessfully\n", username)
}

func Delete(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	if r.Method != "POST" {
		http.Error(w, http.StatusText(405), 405)
		return
	}

	username := r.FormValue("username")
	if username == "" {
		http.Error(w, http.StatusText(400), 400)
	}

	ifUserExist := ifUserExist(username, db)
	err := ifUserExist.Scan(username)
	if err == sql.ErrNoRows {
		fmt.Fprintln(w, "User with this username does not exists")
		return
	}

	createResult := userDelete(username, db)
	if createResult != 200 {
		http.Error(w, http.StatusText(createResult), createResult)
		return
	}

	fmt.Fprintf(w, "User %s deleted sucessfully\n", username)
}

func Login(w http.ResponseWriter, r *http.Request, db *sql.DB) int {
	if r.Method != "POST" {
		http.Error(w, http.StatusText(405), 405)
		return -1
	}

	username := r.FormValue("username")
	password := r.FormValue("password")
	if username == "" || password == "" {
		http.Error(w, http.StatusText(400), 400)
	}

	user := new(User)
	ifUserExist := ifUserExist(username, db)
	err := ifUserExist.Scan(&user.username, &user.password)
	if err == sql.ErrNoRows {
		fmt.Fprintln(w, "User with this username does not exists")
		return -1
	}

	loginResult, userId := userLogin(username, password, db)
	if loginResult != 200 {
		http.Error(w, http.StatusText(loginResult), loginResult)
		return -1
	}

	fmt.Fprintf(w, "User %s logged sucessfully\n", username)
	return userId
}
