package post

import (
	"time"
)

type Post struct {
	PostImage       string `json:"image"`
	PostDescription string `json:"description"`
	CreateTime      time.Time
	UserId          int
	AccessToken     string
}
