package main

import (
	"log"
	"net/http"

	app "github.com/CyberPiess/instagram/internal/app/instagram/application"
	user "github.com/CyberPiess/instagram/internal/app/instagram/domain/user"
	database "github.com/CyberPiess/instagram/internal/app/instagram/infrastructure/database"

	_ "github.com/lib/pq"
)

func main() {

	mux := http.NewServeMux()

	db, err := database.OpenDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	mux.HandleFunc("/createUser", func(w http.ResponseWriter, r *http.Request) {
		app.UserHandler.Create(&user.User{}, w, r, db)
	})

	log.Println("Запуск веб-сервера на http://localhost:8080")
	err = http.ListenAndServe(":8080", mux)
	log.Fatal(err)

}
