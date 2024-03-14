//go:generate mockgen -source=service.go -destination=mocks/mock.go
package post

import (
	"bytes"
	"io"
	"strconv"

	database "github.com/CyberPiess/instagram/internal/infrastructure/database/post"
	"github.com/CyberPiess/instagram/internal/infrastructure/minio/post"
	"github.com/CyberPiess/instagram/internal/infrastructure/token"
	"github.com/google/uuid"
)

type postStorage interface {
	Create(CreatePost database.CreatePost) error
}

type tokenInteraction interface {
	VerifyToken(tokenString string) (*token.Credentials, error)
	CreateToken(userId int) (string, error)
}

type minioStorage interface {
	UploadFile(post.ImageDTO)
}
type PostService struct {
	store postStorage
	token tokenInteraction
	minio minioStorage
}

func NewPostService(store postStorage, token tokenInteraction, minio minioStorage) *PostService {
	return &PostService{store: store,
		token: token,
		minio: minio}
}

func (p *PostService) CreatePost(newPost Post, image Image) error {
	jwtClaims, err := p.token.VerifyToken(newPost.AccessToken)

	if err != nil {
		return err
	}
	userId, err := strconv.Atoi(jwtClaims.UserId)
	if err != nil {
		return err
	}

	postCreate := database.CreatePost{
		PostDescription: newPost.PostDescription,
		CreateTime:      newPost.CreateTime,
		UserId:          userId,
	}

	buf := bytes.NewBuffer(nil)
	if _, err := io.Copy(buf, image.File); err != nil {
		return err
	}

	photoId := uuid.New()

	imageDTO := post.ImageDTO{
		ObjectName:  photoId.String(),
		FileBuff:    buf,
		ContentType: image.ContentType,
		FileSize:    image.FileSize,
	}

	p.minio.UploadFile(imageDTO)
	postCreate.PostImage = photoId.String()

	return p.store.Create(postCreate)
}
