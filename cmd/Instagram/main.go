package main

import (
	"log"
	"net/http"

	appPost "github.com/CyberPiess/instagram/internal/app/instagram/application/post"
	appUser "github.com/CyberPiess/instagram/internal/app/instagram/application/user"

	domainPost "github.com/CyberPiess/instagram/internal/app/instagram/domain/post"
	domainUser "github.com/CyberPiess/instagram/internal/app/instagram/domain/user"

	database "github.com/CyberPiess/instagram/internal/app/instagram/infrastructure/database"
	postRepo "github.com/CyberPiess/instagram/internal/app/instagram/infrastructure/database/post"
	userRepo "github.com/CyberPiess/instagram/internal/app/instagram/infrastructure/database/user"

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

	userStorage := userRepo.NewUserRepository(db)
	postStorage := postRepo.NewPostRepository(db)

	userService := domainUser.NewUserService(userStorage)
	postService := domainPost.NewPostService(postStorage)

	userHandler := appUser.NewUserHandler(userService)
	postHandler := appPost.NewPostHandler(postService)

	mux.HandleFunc("/createUser", userHandler.UserCreate())
	mux.HandleFunc("/loginUser", userHandler.UserLogin())
	mux.HandleFunc("/logoutUser", userHandler.UserLogout())
	mux.HandleFunc("/createPost", postHandler.PostCreate())

	log.Println("Запуск веб-сервера на http://localhost:8080")
	err = http.ListenAndServe(":8080", mux)
	log.Fatal(err)

}
