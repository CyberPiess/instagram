package post

import (
	"fmt"
	"strconv"

	"github.com/golang-jwt/jwt/v5"
)

//go:generate mockgen -source=service.go -destination=mocks/mock.go

const (
	secretKey = "secret"
)

type postStorage interface {
	Create(CreatePOst CreatePost) error
}

type PostService struct {
	store postStorage
}

func NewPostService(store postStorage) *PostService {
	return &PostService{store: store}
}

func (p *PostService) CreatePost(newPost Post) error {
	jwtClaims, err := p.VerifyToken(newPost.AccessToken)
	if err != nil {
		return err
	}
	userId, err := strconv.Atoi(jwtClaims.UserId)
	if err != nil {
		return err
	}

	postCreate := CreatePost{
		PostImage:       newPost.PostImage,
		PostDescription: newPost.PostDescription,
		CreateTime:      newPost.CreateTime,
		UserId:          userId,
	}
	return p.store.Create(postCreate)
}

func (p *PostService) VerifyToken(tokenString string) (*Credentials, error) {
	var jwtClaims Credentials
	token, err := jwt.ParseWithClaims(tokenString, &jwtClaims, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})

	if err != nil {
		return &Credentials{}, err
	}

	if !token.Valid {
		return &Credentials{}, fmt.Errorf("invalid token")
	}

	return &jwtClaims, nil
}
