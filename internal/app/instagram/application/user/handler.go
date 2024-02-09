package application

import (
	"net/http"

	dto "github.com/CyberPiess/instagram/internal/app/instagram/domain/user"
)

type User struct {
}

func (u *User) Create(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, http.StatusText(400), 400)
		return
	}

	username := r.FormValue("username")
	user_email := r.FormValue("user_email")
	password := r.FormValue("password")
	new_user := &dto.User{
		Username:   username,
		User_email: user_email,
		Password:   password,
	}

	dto.Create(*new_user)
}
