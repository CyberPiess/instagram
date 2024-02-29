package application

import (
	"net/http/httptest"
	"strings"
	"testing"

	appUser "github.com/CyberPiess/instagram/internal/app/instagram/application/user/mocks"
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
