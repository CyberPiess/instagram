package instagram

import (
	"database/sql"
	"time"
)

func postCreateWithDescr(post_image string, post_description string, user_id int, create_time time.Time, db *sql.DB) int {
	_, err := db.Query("INSERT INTO \"Posts\" (\"post_image\", \"post_description\", \"create_time\", \"user_id\") VALUES ($1, $2, $3, $4)", post_image, post_description, create_time, user_id)

	if err != nil {
		return 500
	}
	return 200
}

func postCreateWithOutDescr(post_image string, create_time time.Time, user_id int, db *sql.DB) int {
	_, err := db.Query("INSERT INTO \"Posts\" (\"post_image\", \"create_time\", \"user_id\") VALUES ($1, $2, $3)", post_image, create_time, user_id)

	if err != nil {
		return 500
	}
	return 200
}

func postGetAllPosts(user_id int, db *sql.DB) (int, []Post) {
	result, err := db.Query("SELECT \"post_image\", \"post_description\", \"create_time\" FROM \"Posts\" where \"user_id\" = $1", user_id)

	if err != nil {
		error_post := []Post{}
		return 500, error_post
	}

	posts := []Post{}

	for result.Next() {
		p := Post{}
		err := result.Scan(&p.post_image, &p.post_description, &p.create_time)
		if err != nil {
			error_post := []Post{}
			return 500, error_post
		}
		posts = append(posts, p)
	}

	return 200, posts
}
