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
	"github.com/CyberPiess/instagram/internal/app/instagram/infrastructure/minio"
	minioRepo "github.com/CyberPiess/instagram/internal/app/instagram/infrastructure/minio/post"
	"github.com/CyberPiess/instagram/internal/app/instagram/infrastructure/token"

	_ "github.com/lib/pq"
)

func main() {

	mux := http.NewServeMux()

	db, err := database.NewPostgresDb(database.Config{
		Host:     "localhost",
		Port:     "5432",
		Username: "admin",
		DBName:   "Instagram",
		SSLMode:  "disable",
		Password: "password",
	})

	minio, err := minio.NewMinioConnection(minio.MinioCred{
		Endpoint:        "localhost:9000",
		AccessKeyId:     "user",
		SecretAccessKey: "password",
		UseSSL:          false,
	})

	if err != nil {
		log.Fatal("failed to initialize db: %s", err.Error())
	}

	userStorage := userRepo.NewUserRepository(db)
	postStorage := postRepo.NewPostRepository(db)
	minioPostStorage := minioRepo.NewMinioPostStorage(minio)
	tokenInteraction := token.NewToken()

	userService := domainUser.NewUserService(userStorage, tokenInteraction)
	postService := domainPost.NewPostService(postStorage, tokenInteraction, minioPostStorage)

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
