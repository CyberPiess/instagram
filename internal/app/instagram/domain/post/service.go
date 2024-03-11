//go:generate mockgen -source=service.go -destination=mocks/mock.go
package post

import (
	"fmt"
	"io"
	"os"
	"strconv"

	database "github.com/CyberPiess/instagram/internal/app/instagram/infrastructure/database/post"
	"github.com/CyberPiess/instagram/internal/app/instagram/infrastructure/minio/post"
	"github.com/CyberPiess/instagram/internal/app/instagram/infrastructure/token"
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

	dst, err := os.Create(image.ObjectName)
	if err != nil {
		return fmt.Errorf("error creating file")

	}
	defer dst.Close()
	if _, err := io.Copy(dst, image.File); err != nil {
		return fmt.Errorf("error copying file")
	}

	imageDTO := post.ImageDTO{
		ObjectName:  image.ObjectName,
		FilePath:    image.ObjectName,
		ContentType: image.ContentType,
		FileSize:    image.FileSize,
	}

	p.minio.UploadFile(imageDTO)

	return p.store.Create(postCreate)
}
