package user

import (
	"net/http"
	"time"

	dto "github.com/CyberPiess/instagram/internal/app/instagram/domain/user"
	database "github.com/CyberPiess/instagram/internal/app/instagram/infrastructure/database"
)

type UserHandler interface {
	UserCreate(env *database.Env) http.HandlerFunc
}

func UserCreate(env *database.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			http.Error(w, http.StatusText(400), 400)
			return
		}

		username := r.FormValue("username")
		user_email := r.FormValue("user_email")
		password := r.FormValue("password")
		new_user := dto.User{
			Username:    username,
			User_email:  user_email,
			Password:    password,
			Create_time: time.Now(),
			DB:          *env,
		}
		storage := &dto.UserService{}

		err := storage.CreateUser(new_user)
		if err != nil {
			http.Error(w, http.StatusText(500), 500)
			return
		}
	}
}
