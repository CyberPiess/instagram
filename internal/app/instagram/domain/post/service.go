package post

import (
	"fmt"
	"strconv"

	database "github.com/CyberPiess/instagram/internal/app/instagram/infrastructure/database/post"
	"github.com/CyberPiess/instagram/internal/app/instagram/infrastructure/token"
)

//go:generate mockgen -source=service.go -destination=mocks/mock.go

type postStorage interface {
	Create(CreatePost database.CreatePost) error
}

type tokenInteraction interface {
	VerifyToken(tokenString string) (*token.Credentials, error)
}

type PostService struct {
	store postStorage
	token tokenInteraction
}

func NewPostService(store postStorage, token tokenInteraction) *PostService {
	return &PostService{store: store,
		token: token}
}

func (p *PostService) CreatePost(newPost Post) error {
	jwtClaims, err := p.token.VerifyToken(newPost.AccessToken)

	if newPost.PostImage == "" {
		return fmt.Errorf("no post supplied")
	}
	if err != nil {
		return err
	}
	userId, err := strconv.Atoi(jwtClaims.UserId)
	if err != nil {
		return err
	}

	postCreate := database.CreatePost{
		PostImage:       newPost.PostImage,
		PostDescription: newPost.PostDescription,
		CreateTime:      newPost.CreateTime,
		UserId:          userId,
	}
	return p.store.Create(postCreate)
}
