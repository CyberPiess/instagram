package instagram

import (
	"database/sql"
	"fmt"
	"net/http"
	"time"
)

type User struct {
	username        string
	user_email      string
	hashed_password string
	create_time     time.Time
}

func Create(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	if r.Method != "POST" {
		http.Error(w, http.StatusText(400), 400)
		return
	}

	username := r.FormValue("username")
	password := r.FormValue("password")
	user_email := r.FormValue("user_email")
	if username == "" || password == "" || user_email == "" {
		http.Error(w, http.StatusText(400), 400)
		return
	}

	if_user_exist := ifUserExist(username, db)
	if_email_exist := ifEmailExist(user_email, db)
	switch {
	case if_user_exist != nil && if_user_exist != sql.ErrNoRows:
		fmt.Fprintln(w, "Bad request")
		return
	case if_user_exist == nil:
		fmt.Fprintln(w, "User with this username already exists")
		return
	case if_email_exist != nil && if_email_exist != sql.ErrNoRows:
		fmt.Fprintln(w, "Bad request")
		return
	case if_email_exist == nil:
		fmt.Fprintln(w, "User with this email already exists")
		return
	}

	hashed_password := hashAndSalt([]byte(password))
	new_user := User{
		username:        username,
		user_email:      user_email,
		hashed_password: hashed_password,
		create_time:     time.Now(),
	}
	create_result := userCreate(new_user, db)
	if !create_result {
		http.Error(w, http.StatusText(500), 500)
		return
	}

	fmt.Fprintf(w, "User %s created sucessfully\n", username)
}

func Update(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	if r.Method != "POST" {
		http.Error(w, http.StatusText(400), 400)
		return
	}

	username := r.FormValue("username")
	password := r.FormValue("password")
	new_username := r.FormValue("new_username")
	new_password := r.FormValue("new_password")
	if username == "" || password == "" {
		http.Error(w, http.StatusText(400), 400)
	}

	user_exist := ifUserExist(username, db)
	if user_exist != nil {
		fmt.Fprintln(w, "User with current username does not exists")
		return
	}

	user_exist = ifUserExist(new_username, db)
	if user_exist != nil {
		fmt.Fprintln(w, "User with new username already exists")
		return
	}

	hashed_new_password := hashAndSalt([]byte(new_password))
	update_result := userUpdate(username, password, new_username, hashed_new_password, db)
	if update_result != 200 {
		http.Error(w, http.StatusText(update_result), update_result)
		return
	}

	fmt.Fprintf(w, "User %s updated sucessfully\n", username)
}

func Delete(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	if r.Method != "POST" {
		http.Error(w, http.StatusText(400), 400)
		return
	}

	username := r.FormValue("username")
	if username == "" {
		http.Error(w, http.StatusText(400), 400)
	}

	if_user_exist := ifUserExist(username, db)
	if if_user_exist != nil {
		fmt.Fprintln(w, "User with this username does not exists")
		return
	}

	create_result := userDelete(username, db)
	if create_result != 200 {
		http.Error(w, http.StatusText(create_result), create_result)
		return
	}

	fmt.Fprintf(w, "User %s deleted sucessfully\n", username)
}

func Login(w http.ResponseWriter, r *http.Request, db *sql.DB) int {
	if r.Method != "POST" {
		http.Error(w, http.StatusText(400), 400)
		return -1
	}

	username := r.FormValue("username")
	password := r.FormValue("password")
	if username == "" || password == "" {
		http.Error(w, http.StatusText(400), 400)
	}

	if_user_exist := ifUserExist(username, db)
	if if_user_exist != nil {
		fmt.Fprintln(w, "User with this username does not exists")
		return -1
	}

	login_result, user_id := userLogin(username, password, db)
	if login_result != 200 {
		http.Error(w, http.StatusText(login_result), login_result)
		return -1
	}

	fmt.Fprintf(w, "User %s logged sucessfully\n", username)
	return user_id
}
