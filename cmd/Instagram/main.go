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

	mux.HandleFunc("/deleteUser", func(w http.ResponseWriter, r *http.Request) {
		instagram.Delete(w, r, db)
	})

	mux.HandleFunc("/logUser", func(w http.ResponseWriter, r *http.Request) {
		instagram.Login(w, r, db)
	})

	mux.HandleFunc("/updateUser", func(w http.ResponseWriter, r *http.Request) {
		instagram.Update(w, r, db)
	})

	mux.HandleFunc("/createPost", func(w http.ResponseWriter, r *http.Request) {
		instagram.PostCreate(w, r, db)
	})

	mux.HandleFunc("/getPost", func(w http.ResponseWriter, r *http.Request) {
		instagram.PostGet(w, r, db)
	})

	log.Println("Запуск веб-сервера на http://localhost:8080")
	err = http.ListenAndServe(":8080", mux)
	log.Fatal(err)

}
