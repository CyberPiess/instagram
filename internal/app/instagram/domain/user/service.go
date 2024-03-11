//go:generate mockgen -source=service.go -destination=mocks/mock.go
package user

import (
	"fmt"
	"strconv"

	"github.com/CyberPiess/instagram/internal/app/instagram/infrastructure/database/user"
	"github.com/CyberPiess/instagram/internal/app/instagram/infrastructure/token"
)

type userStorage interface {
	Insert(newUser user.UserDTO) error
	IfUserExist(newUser user.UserDTO) (bool, error)
	IfEmailExist(newUSer user.UserDTO) (bool, error)
	SelectUser(username string) (int, string, error)
}

type tokenInteraction interface {
	VerifyToken(tokenString string) (*token.Credentials, error)
	CreateToken(userId int) (string, error)
}

type UserService struct {
	store userStorage
	token tokenInteraction
}

func NewUserService(storage userStorage, token tokenInteraction) *UserService {
	return &UserService{store: storage,
		token: token}
}

func (u *UserService) CreateUser(newUser User) error {

	err := u.VerifyData(newUser)
	if err != nil {
		return err
	}

	userDTO := user.UserDTO{
		Username:   newUser.Username,
		UserEmail:  newUser.UserEmail,
		Password:   hashAndSalt([]byte(newUser.Password)),
		CreateTime: newUser.CreateTime,
	}

	userExists, err := u.store.IfUserExist(userDTO)
	if err != nil {
		return err
	} else if userExists {
		return fmt.Errorf("user with this username already exists")
	}

	emailExists, err := u.store.IfEmailExist(userDTO)
	if err != nil {
		return err
	} else if emailExists {
		return fmt.Errorf("user with this email already exists")
	}

	return u.store.Insert(userDTO)
}

func (u *UserService) LoginUser(req *LoginUserReq) (*LoginUserRes, error) {
	userId, hashedPassword, err := u.store.SelectUser(req.Username)
	if err != nil {
		return &LoginUserRes{}, err
	}

	if !comparePasswords(hashedPassword, []byte(req.Password)) {
		return &LoginUserRes{}, fmt.Errorf("incorrect username or password")
	}

	token, err := u.token.CreateToken(userId)

	return &LoginUserRes{AccessToken: token, Username: req.Username, UserId: strconv.Itoa(userId)}, err
}

func (u *UserService) VerifyData(newUser User) error {
	switch {
	case newUser.Username == "":
		return fmt.Errorf("invalid input: empty username")
	case newUser.UserEmail == "":
		return fmt.Errorf("invalid input: empty email")
	case newUser.Password == "":
		return fmt.Errorf("invalid input: empty password")
	}
	return nil
}
