package user

import (
	"fmt"
	"testing"
	"time"

	u "github.com/CyberPiess/instagram/internal/app/instagram/domain/user/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

type args struct {
	user User
}

func TestCreateUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserStorage := u.NewMockuserStorage(ctrl)
	mockTokenInteraction := u.NewMocktokenInteraction(ctrl)
	userService := NewUserService(mockUserStorage, mockTokenInteraction)

	mockUserStorage.EXPECT().IfUserExist(gomock.Any()).Return(false, nil).AnyTimes()
	mockUserStorage.EXPECT().IfEmailExist(gomock.Any()).Return(false, nil).AnyTimes()
	mockUserStorage.EXPECT().Insert(gomock.Any()).Return(nil).AnyTimes()

	tests := []struct {
		name string
		args args
		want error
	}{
		{
			name: "Correct data",
			args: args{User{Username: "someUsername",
				UserEmail:  "someUserEmail",
				Password:   "somePassword",
				CreateTime: time.Now()},
			},
			want: nil,
		},
	}

	for _, tt := range tests {
		err := userService.CreateUser(tt.args.user)
		assert.Equal(t, tt.want, err)
	}
}

func TestLoginUser(t *testing.T) {

	type localArgs struct {
		loginUser LoginUserReq
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserStorage := u.NewMockuserStorage(ctrl)
	mockTokenInteraction := u.NewMocktokenInteraction(ctrl)
	userService := NewUserService(mockUserStorage, mockTokenInteraction)

	mockTokenInteraction.EXPECT().CreateToken(gomock.Any()).Return("someToken", nil).AnyTimes()
	mockUserStorage.EXPECT().SelectUser(gomock.Any()).Return(1, "$2a$04$daDkOqpORBtgVFjFFcmBj.BzVq9lFz7J932T3dTdU0RGCjvhQSmNy", nil).AnyTimes()

	tests := []struct {
		name string
		args localArgs
		want error
	}{{
		name: "Incorrect password",
		args: localArgs{loginUser: LoginUserReq{Username: "somename",
			Password: "somepwd"},
		},
		want: fmt.Errorf("incorrect username or password"),
	},
		{
			name: "Correct password",
			args: localArgs{loginUser: LoginUserReq{Username: "test",
				Password: "test"},
			},
			want: nil,
		}}

	for _, tt := range tests {
		_, err := userService.LoginUser(&tt.args.loginUser)
		assert.Equal(t, tt.want, err)

	}
}

func TestVerifyDate(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserStorage := u.NewMockuserStorage(ctrl)
	mockTokenInteraction := u.NewMocktokenInteraction(ctrl)
	userService := NewUserService(mockUserStorage, mockTokenInteraction)

	tests := []struct {
		name string
		args args
		want error
	}{{
		name: "Empty username",
		args: args{
			User{Username: "",
				UserEmail: "someUserEmail",
				Password:  "somePassword"},
		},
		want: fmt.Errorf("invalid input: empty username"),
	},
		{
			name: "Empty user email",
			args: args{
				User{Username: "someUsername",
					UserEmail: "",
					Password:  "somePassword"},
			},
			want: fmt.Errorf("invalid input: empty email"),
		},
		{
			name: "Empty password",
			args: args{
				User{Username: "someUsername",
					UserEmail: "someUserEmail",
					Password:  ""},
			},
			want: fmt.Errorf("invalid input: empty password"),
		},
		{
			name: "Correct data",
			args: args{
				User{Username: "someUsername",
					UserEmail: "someUserEmail",
					Password:  "somePassord"},
			},
			want: nil,
		},
	}

	for _, tt := range tests {
		err := userService.VerifyData(tt.args.user)
		assert.Equal(t, tt.want, err)
	}

}
