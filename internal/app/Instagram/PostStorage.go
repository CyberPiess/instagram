package instagram

import (
	"database/sql"
	"time"
)

func postCreateWithDescr(postImage string, postDescription string, userId int, createTime time.Time, db *sql.DB) int {
	_, err := db.Query("INSERT INTO \"Posts\" (\"postImage\", \"postDescription\", \"createTime\", \"userId\") VALUES ($1, $2, $3, $4)", postImage, postDescription, createTime, userId)

	if err != nil {
		return 500
	}
	return 200
}

func postCreateWithOutDescr(postImage string, createTime time.Time, userId int, db *sql.DB) int {
	_, err := db.Query("INSERT INTO \"Posts\" (\"postImage\", \"createTime\", \"userId\") VALUES ($1, $2, $3)", postImage, createTime, userId)

	if err != nil {
		return 500
	}
	return 200
}

func postGetAllPosts(userId int, db *sql.DB) (int, []Post) {
	result, err := db.Query("SELECT \"postImage\", \"postDescription\", \"createTime\" FROM \"Posts\" where \"userId\" = $1", userId)

	if err != nil {
		errorPost := []Post{}
		return 500, errorPost
	}

	posts := []Post{}

	for result.Next() {
		p := Post{}
		err := result.Scan(&p.postImage, &p.postDescription, &p.createTime)
		if err != nil {
			errorPost := []Post{}
			return 500, errorPost
		}
		posts = append(posts, p)
	}

	return 200, posts
}
