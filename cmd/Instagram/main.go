package main

import (
	"log"
	"net/http"

	app "github.com/CyberPiess/instagram/internal/app/instagram/application/user"
	database "github.com/CyberPiess/instagram/internal/app/instagram/infrastructure/database"

	_ "github.com/lib/pq"
)

func main() {

	mux := http.NewServeMux()

	db, err := database.NewPostgresDb(database.Config{
		Host:     "localhost",
		Port:     "5432",
		Username: "admin",
		DBName:   "Instagram",
		SSLMode:  "disabled",
		Password: "password",
	})

	if err != nil {
		log.Fatal("failed to initialize db: %s", err.Error())
	}
	defer db.Close()

	var user app.User

	mux.HandleFunc("/createUser", user.Create)

	log.Println("Запуск веб-сервера на http://localhost:8080")
	err = http.ListenAndServe(":8080", mux)
	log.Fatal(err)

}
