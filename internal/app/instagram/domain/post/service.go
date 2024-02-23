package post

import (
	"fmt"
	"strconv"

	"github.com/golang-jwt/jwt/v5"
)

const (
	secretKey = "secret"
)

type postStorage interface {
	Create(CreatePOst CreatePostReq) error
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

	postCreate := CreatePostReq{
		PostImage:       newPost.PostImage,
		PostDescription: newPost.PostDescription,
		CreateTime:      newPost.CreateTime,
		UserId:          userId,
	}
	return p.store.Create(postCreate)
}

func (p *PostService) VerifyToken(tokenString string) (*MyJWTClaims, error) {
	var jwtClaims MyJWTClaims
	token, err := jwt.ParseWithClaims(tokenString, &jwtClaims, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})

	if err != nil {
		return &MyJWTClaims{}, err
	}

	if !token.Valid {
		return &MyJWTClaims{}, fmt.Errorf("invalid token")
	}

	return &jwtClaims, nil
}
