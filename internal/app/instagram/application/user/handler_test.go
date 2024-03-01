package application

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	appUser "github.com/CyberPiess/instagram/internal/app/instagram/application/user/mocks"
	"github.com/CyberPiess/instagram/internal/app/instagram/domain/user"
	"github.com/golang/mock/gomock"

	"github.com/stretchr/testify/assert"
)

type args struct {
	method    string
	parametrs string
}

func TestUserCreate(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserService := appUser.NewMockuserService(ctrl)

	userHandler := NewUserHandler(mockUserService)

	mockUserService.EXPECT().CreateUser(gomock.Any()).Return(nil)

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Method POST",
			args: args{method: "POST",
				parametrs: ""},
			want: 200,
		},
		{
			name: "Method GET",
			args: args{method: "GET",
				parametrs: ""},
			want: 400,
		},
	}

	for _, tt := range tests {
		req := httptest.NewRequest(tt.args.method, "/createUser", strings.NewReader(tt.args.parametrs))
		w := httptest.NewRecorder()

		userHandler.UserCreate().ServeHTTP(w, req)
		assert.Equal(t, tt.want, w.Code)
	}
}

func TestUserLogin(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserService := appUser.NewMockuserService(ctrl)

	userHandler := NewUserHandler(mockUserService)

	userResult := &user.LoginUserRes{
		AccessToken: "",
		UserId:      "",
		Username:    "",
	}
	mockUserService.EXPECT().LoginUser(gomock.Any()).Return(userResult, nil)

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Method POST",
			args: args{method: "POST",
				parametrs: ""},
			want: 200,
		},
		{
			name: "Method GET",
			args: args{method: "GET",
				parametrs: ""},
			want: 400,
		},
	}

	for _, tt := range tests {
		req := httptest.NewRequest(tt.args.method, "/loginUser", strings.NewReader(tt.args.parametrs))
		w := httptest.NewRecorder()

		userHandler.UserLogin().ServeHTTP(w, req)
		assert.Equal(t, tt.want, w.Code)
	}
}

func TestUserLogout(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserService := appUser.NewMockuserService(ctrl)

	userHandler := NewUserHandler(mockUserService)
	tests := []struct {
		args args
		want string
	}{
		{
			args: args{method: "GET"},
			want: "",
		},
	}

	for _, tt := range tests {
		req := httptest.NewRequest(tt.args.method, "/logoutUser", nil)
		w := httptest.NewRecorder()

		http.SetCookie(w, &http.Cookie{
			Name:   "jwt",
			Value:  "testcookie",
			MaxAge: 3600,
			Path:   "/",

			Secure:   false,
			HttpOnly: true,
		})
		userHandler.UserLogout().ServeHTTP(w, req)
		assert.Equal(t, tt.want, w.Result().Cookies()[1].Value)
	}
}
