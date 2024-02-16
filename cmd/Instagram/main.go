package main

import (
	"log"
	"net/http"

	userHandler "github.com/CyberPiess/instagram/internal/app/instagram/application/user"
	database "github.com/CyberPiess/instagram/internal/app/instagram/infrastructure/database"

	_ "github.com/lib/pq"
)

func main() {

	mux := http.NewServeMux()

	db, err := database.NewPostgresDb(database.Config{
		Host:     "localhost",
		Port:     "5432",
		Username: "postgres",
		DBName:   "Instagram",
		SSLMode:  "disable",
		Password: "password",
	})

	if err != nil {
		log.Fatal("failed to initialize db: %s", err.Error())
	}

	mux.HandleFunc("/createUser", userHandler.UserCreate(db))

	log.Println("Запуск веб-сервера на http://localhost:8080")
	err = http.ListenAndServe(":8080", mux)
	log.Fatal(err)

}
