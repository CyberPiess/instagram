package instagram

import (
	"database/sql"
	"fmt"
	"net/http"
	"time"

	user "github.com/CyberPiess/instagram/internal/app/instagram/user"
)

type Post struct {
	post_image       string
	post_description *string
	create_time      time.Time
}

func PostCreate(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	user_id := user.Login(w, r, db)

	post_description := r.FormValue("post_description")
	post_image := r.FormValue("post_image")
	post_create := time.Now()
	if post_image == "" {
		http.Error(w, http.StatusText(400), 400)
	}

	switch {
	case post_description != "":
		post_create_result := postCreateWithDescr(post_image, post_description, user_id, post_create, db)
		if post_create_result != 200 {
			http.Error(w, http.StatusText(post_create_result), post_create_result)
			return
		}
	case post_description == "":
		post_create_result := postCreateWithOutDescr(post_image, post_create, user_id, db)
		if post_create_result != 200 {
			http.Error(w, http.StatusText(post_create_result), post_create_result)
			return
		}
	}

	fmt.Fprintf(w, "Post created successfully\n")

}

func PostGet(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	user_id := user.Login(w, r, db)

	post_get_result, posts_found := postGetAllPosts(user_id, db)
	if post_get_result != 200 {
		http.Error(w, http.StatusText(post_get_result), post_get_result)
		return
	}

	if len(posts_found) > 0 {
		for _, value := range posts_found {
			fmt.Fprintf(w, "\npost image: %s\n", value.post_image)
			if value.post_description == nil {
				fmt.Fprintf(w, "post description: \n")
			} else {
				fmt.Fprintf(w, "post description: %v\n", *value.post_description)
			}

			fmt.Fprintf(w, "create time: %s\n", value.create_time.Format("2006-01-02 15:04:05"))
		}
	} else {
		fmt.Fprintf(w, "\nNo posts found\n")
	}
}
