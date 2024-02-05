package user

import (
	"database/sql"
	"net/http"

	"github.com/CyberPiess/instagram/internal/app/instagram/infrastructure/user"
)

func (u *User) Create(w http.ResponseWriter, r *http.Request, db *sql.DB) {
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

	//для обращения к функциям в пакете storage.go делаю импорт: "github.com/CyberPiess/instagram/internal/app/instagram/infrastructure/user"
	//но суть в том, что в пакете есть функция Create, которая в качестве одного из аргументов принимает структуру User
	//соответственно в storage.go осуществляется импорт текущего пакета
	//так я получаю цикл
	//я понимаю, почему я получаю цикл, как его исправить?
	user_repository := &user.UserRepository{}
	if_user_exist := user_repository.IfUsernameExist(username, db)
	print(if_user_exist)
}
