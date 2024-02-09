package user

import (
	"github.com/CyberPiess/instagram/internal/app/instagram/infrastructure/database"
)

type AuthService struct {
	repository database.Authorization
}

func NewAuthService(repository database.Authorization) *AuthService {
	return &AuthService{repository: repository}
}
