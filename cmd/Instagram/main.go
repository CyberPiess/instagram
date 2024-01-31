package main

import (
	"log"
	"net/http"

	database "github.com/CyberPiess/instagram/internal/app/instagram/database"
	posts "github.com/CyberPiess/instagram/internal/app/instagram/posts"
	user "github.com/CyberPiess/instagram/internal/app/instagram/user"

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
		user.Create(w, r, db)
	})

	mux.HandleFunc("/deleteUser", func(w http.ResponseWriter, r *http.Request) {
		user.Delete(w, r, db)
	})

	mux.HandleFunc("/logUser", func(w http.ResponseWriter, r *http.Request) {
		user.Login(w, r, db)
	})

	mux.HandleFunc("/updateUser", func(w http.ResponseWriter, r *http.Request) {
		user.Update(w, r, db)
	})

	mux.HandleFunc("/createPost", func(w http.ResponseWriter, r *http.Request) {
		posts.PostCreate(w, r, db)
	})

	mux.HandleFunc("/getPost", func(w http.ResponseWriter, r *http.Request) {
		posts.PostGet(w, r, db)
	})

	log.Println("Запуск веб-сервера на http://localhost:8080")
	err = http.ListenAndServe(":8080", mux)
	log.Fatal(err)

}
