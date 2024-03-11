package database

import "time"

type CreatePost struct {
	PostImage       string
	PostDescription string
	CreateTime      time.Time
	UserId          int
}
