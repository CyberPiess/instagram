package post

import "bytes"

type ImageDTO struct {
	ObjectName  string
	FileBuff    *bytes.Buffer
	ContentType string
	FileSize    int64
}
