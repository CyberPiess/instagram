package instagram

import (
	"database/sql"
	"fmt"
	"net/http"
	"time"
)

type Post struct {
	postImage       string
	postDescription *string
	createTime      time.Time
}

func PostCreate(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	userId := Login(w, r, db)

	postDescription := r.FormValue("postDescription")
	postImage := r.FormValue("postImage")
	postCreate := time.Now()
	if postImage == "" {
		http.Error(w, http.StatusText(400), 400)
	}

	switch {
	case postDescription != "":
		postCreateResult := postCreateWithDescr(postImage, postDescription, userId, postCreate, db)
		if postCreateResult != 200 {
			http.Error(w, http.StatusText(postCreateResult), postCreateResult)
			return
		}
	case postDescription == "":
		postCreateResult := postCreateWithOutDescr(postImage, postCreate, userId, db)
		if postCreateResult != 200 {
			http.Error(w, http.StatusText(postCreateResult), postCreateResult)
			return
		}
	}

	fmt.Fprintf(w, "Post created successfully\n")

}

func PostGet(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	userId := Login(w, r, db)

	postGetResult, postsFound := postGetAllPosts(userId, db)
	if postGetResult != 200 {
		http.Error(w, http.StatusText(postGetResult), postGetResult)
		return
	}

	if len(postsFound) > 0 {
		for _, value := range postsFound {
			fmt.Fprintf(w, "\npost image: %s\n", value.postImage)
			if value.postDescription == nil {
				fmt.Fprintf(w, "post description: \n")
			} else {
				fmt.Fprintf(w, "post description: %v\n", *value.postDescription)
			}

			fmt.Fprintf(w, "create time: %s\n", value.createTime.Format("2006-01-02 15:04:05"))
		}
	} else {
		fmt.Fprintf(w, "\nNo posts found\n")
	}
}
