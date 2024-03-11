//go:generate mockgen -source=handler.go -destination=mocks/post_service_mock.go
package application

import (
	"net/http"
	"time"

	"github.com/CyberPiess/instagram/internal/app/instagram/domain/post"
)

type postService interface {
	CreatePost(newPost post.Post, image post.Image) error
}

type Post struct {
	service postService
}

func NewPostHandler(service postService) *Post {
	return &Post{service: service}
}

func (p *Post) PostCreate() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, http.StatusText(400), 400)
			return
		}

		cookie, err := r.Cookie("jwt")
		if err != nil {
			http.Error(w, "No Token Found", 400)
			return
		}

		var newPost post.Post

		newPost.PostDescription = r.FormValue("PostDescription")
		newPost.CreateTime = time.Now()
		newPost.AccessToken = cookie.Value

		file, header, err := r.FormFile("myFile")
		if err != nil {
			http.Error(w, "Wrong data supplied", 400)
			return
		}
		defer file.Close()

		image := post.Image{
			ObjectName:  header.Filename,
			File:        file,
			ContentType: header.Header["Content-Type"][0],
			FileSize:    header.Size,
		}

		err = p.service.CreatePost(newPost, image)
		if err != nil {
			http.Error(w, err.Error(), 400)
			return
		}
		w.Write([]byte("Post created successfully"))
	}

}
