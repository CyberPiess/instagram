//go:generate mockgen -source=handler.go -destination=mocks/mock.go
package application

import (
	"net/http"
	"time"

	"github.com/CyberPiess/instagram/internal/app/instagram/domain/user"
)

type userService interface {
	CreateUser(newUser user.User) error
	LoginUser(logUser *user.LoginUserReq) (*user.LoginUserRes, error)
	VerifyData(newUser user.User) error
}

type User struct {
	service userService
}

func NewUserHandler(service userService) *User {
	return &User{service: service}
}

func (u *User) UserCreate() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, http.StatusText(400), 400)
			return
		}

		username := r.FormValue("username")
		userEmail := r.FormValue("userEmail")
		password := r.FormValue("password")
		newUser := user.User{
			Username:   username,
			UserEmail:  userEmail,
			Password:   password,
			CreateTime: time.Now(),
		}

		err := u.service.CreateUser(newUser)
		if err != nil {
			http.Error(w, http.StatusText(500), 500)
			return
		}

		w.Write([]byte("User created!"))
	}
}

func (u *User) UserLogin() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, http.StatusText(400), 400)
			return
		}

		username := r.FormValue("username")
		password := r.FormValue("password")
		logUser := user.LoginUserReq{
			Username: username,
			Password: password,
		}

		logedUser, err := u.service.LoginUser(&logUser)
		if err != nil {
			if err.Error() == "incorrect username or password" {
				http.Error(w, http.StatusText(401), 401)
			} else {
				http.Error(w, http.StatusText(500), 500)
			}

			return
		}

		http.SetCookie(w, &http.Cookie{
			Name:   "jwt",
			Value:  logedUser.AccessToken,
			MaxAge: 3600,
			Path:   "/",

			Secure:   false,
			HttpOnly: true,
		})

		w.Write([]byte("User logged in!"))
	}
}

func (u *User) UserLogout() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		http.SetCookie(w, &http.Cookie{
			Name:   "jwt",
			Value:  "",
			MaxAge: 3600,

			Secure:   false,
			HttpOnly: true,
		})
	}
}
