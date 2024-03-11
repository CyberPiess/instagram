package post

import (
	"mime/multipart"
	"time"
)

type Post struct {
	PostDescription string
	CreateTime      time.Time
	UserId          int
	AccessToken     string
}

type Image struct {
	ObjectName  string
	File        multipart.File
	ContentType string
	FileSize    int64
}
