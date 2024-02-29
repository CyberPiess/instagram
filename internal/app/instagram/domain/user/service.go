package user

import (
	"fmt"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

//go:generate mockgen -source=service.go -destination=mocks/mock.go

const (
	secretKey = "secret"
)

type userStorage interface {
	Insert(newUser User) error
	IfUserExist(newUser User) (bool, error)
	IfEmailExist(newUSer User) (bool, error)
	SelectUser(username string) (int, string, error)
}

type UserService struct {
	store userStorage
}

func NewUserService(storage userStorage) *UserService {
	return &UserService{store: storage}
}

func (u *UserService) CreateUser(newUser User) error {

	err := u.VerifyData(newUser)
	if err != nil {
		return err
	}

	newUser.Password = hashAndSalt([]byte(newUser.Password))

	userExists, err := u.store.IfUserExist(newUser)
	if err != nil {
		return err
	} else if userExists {
		return fmt.Errorf("user with this username already exists")
	}

	emailExists, err := u.store.IfEmailExist(newUser)
	if err != nil {
		return err
	} else if emailExists {
		return fmt.Errorf("user with this email already exists")
	}

	return u.store.Insert(newUser)
}

func (u *UserService) LoginUser(req *LoginUserReq) (*LoginUserRes, error) {
	userId, hashedPassword, err := u.store.SelectUser(req.Username)
	if err != nil {
		return &LoginUserRes{}, err
	}

	if !comparePasswords(hashedPassword, []byte(req.Password)) {
		return &LoginUserRes{}, fmt.Errorf("incorrect username or password")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, MyJWTClaims{
		UserId: strconv.Itoa(userId),
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    strconv.Itoa(userId),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		},
	})

	ss, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return &LoginUserRes{}, err
	}

	return &LoginUserRes{AccessToken: ss, Username: req.Username, UserId: strconv.Itoa(userId)}, err
}

func (u *UserService) VerifyToken(tokenString string) (*MyJWTClaims, error) {
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
