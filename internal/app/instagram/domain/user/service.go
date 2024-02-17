package user

type userStorage interface {
	Insert(newUser User) error
}

type UserService struct {
	store userStorage
}

func NewUserService(storage userStorage) *UserService {
	return &UserService{store: storage}
}

func (u *UserService) CreateUser(newUser User) error {
	if newUser.Username == "" || newUser.User_email == "" || newUser.Password == "" {
		return nil
	}

	newUser.Password = hashAndSalt([]byte(newUser.Password))

	err := u.store.Insert(newUser)

	return err
}
