package application

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/CyberPiess/instagram/internal/app/instagram/domain/post"
)

type postService interface {
	CreatePost(newPost post.Post) error
	VerifyToken(tokenString string) (*post.MyJWTClaims, error)
}

type Post struct {
	service postService
}

func NewPostHandler(service postService) *Post {
	return &Post{service: service}
}

func (p *Post) PostCreate() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("jwt")
		if err != nil {
			http.Error(w, "No Token Found", 400)
			return
		}

		var post post.Post

		err = json.NewDecoder(r.Body).Decode(&post)
		if err != nil {
			http.Error(w, "Wrong data supplied", 400)
			return
		}
		if post.PostImage == "" {
			http.Error(w, "No post supplied", 400)
			return
		}
		post.CreateTime = time.Now()
		post.AccessToken = cookie.Value

		err = p.service.CreatePost(post)
		if err != nil {
			http.Error(w, err.Error(), 400)
			return
		}
		w.Write([]byte("Post created successfully"))
	}

}
