package user

import (
	"net/http"
	"time"

	"github.com/CyberPiess/instagram/internal/app/instagram/domain/user"
)

type userService interface {
	CreateUser(newUser user.User) error
}

type User struct {
	service userService
}

func NewUserHandler(service userService) *User {
	return &User{service: service}
}

func (u *User) UserCreate() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			http.Error(w, http.StatusText(400), 400)
			return
		}

		username := r.FormValue("username")
		user_email := r.FormValue("user_email")
		password := r.FormValue("password")
		newUser := user.User{
			Username:    username,
			User_email:  user_email,
			Password:    password,
			Create_time: time.Now(),
		}

		err := u.service.CreateUser(newUser)
		if err != nil {
			http.Error(w, http.StatusText(500), 500)
			return
		}
	}
}
