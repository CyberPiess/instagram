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

func Home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello"))
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
	}

	user := new(User)
	ifUserExist := ifExist(username, db)
	err := ifUserExist.Scan(&user.username, &user.password)
	if err != sql.ErrNoRows {
		fmt.Fprintln(w, "User with this username already exists")
		return
	}

	//TODO: поменять отдельные поля на структуру?
	createResult := create(username, password, db)
	if createResult != 200 {
		http.Error(w, http.StatusText(createResult), createResult)
		return
	}

	fmt.Fprintf(w, "User %s created sucessfully\n", username)
}

// func Update(w http.ResponseWriter, r *http.Request, db *sql.DB) {
// }

func Delete(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	if r.Method != "POST" {
		http.Error(w, http.StatusText(405), 405)
		return
	}

	username := r.FormValue("username")
	if username == "" {
		http.Error(w, http.StatusText(400), 400)
	}

	ifUserExist := ifExist(username, db)
	err := ifUserExist.Scan(username)
	if err == sql.ErrNoRows {
		fmt.Fprintln(w, "User with this username does not exists")
		return
	}

	createResult := delete(username, db)
	if createResult != 200 {
		http.Error(w, http.StatusText(createResult), createResult)
		return
	}

	fmt.Fprintf(w, "User %s deleted sucessfully\n", username)
}
