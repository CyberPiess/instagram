package main

import (
	"log"
	"net/http"

	instagram "github.com/CyberPiess/Instagram-2.0/internal/app/Instagram"

	_ "github.com/lib/pq"
)

func main() {

	mux := http.NewServeMux()

	db, err := instagram.OpenDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	mux.HandleFunc("/createUser", func(w http.ResponseWriter, r *http.Request) {
		instagram.Create(w, r, db)
	})

	log.Println("Запуск веб-сервера на http://localhost:8080")
	err = http.ListenAndServe(":8080", mux)
	log.Fatal(err)

}