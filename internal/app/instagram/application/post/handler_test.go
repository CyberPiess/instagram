package application

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	mocks "github.com/CyberPiess/instagram/internal/app/instagram/application/post/mocks"
	"github.com/golang/mock/gomock"

	"github.com/stretchr/testify/assert"
)

type customCookie struct {
	cookieName  string
	cookieValue string
}

type args struct {
	method string
	cookie customCookie
	body   string
}

func TestPostCreate(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockPostService := mocks.NewMockpostService(ctrl)

	postHandler := NewPostHandler(mockPostService)

	mockPostService.EXPECT().CreatePost(gomock.Any()).Return(nil)

	tests := []struct {
		name string
		args args
		want int
	}{{
		name: "Method POST",
		args: args{method: "POST",
			cookie: customCookie{cookieName: "",
				cookieValue: ""},
			body: "image:someImage, description:some_description"},
		want: 400,
	},
	}

	for _, tt := range tests {
		req := httptest.NewRequest(tt.args.method, "/createUser", strings.NewReader(tt.args.body))
		w := httptest.NewRecorder()
		http.SetCookie(w, &http.Cookie{
			Name:   tt.args.cookie.cookieName,
			Value:  tt.args.cookie.cookieValue,
			MaxAge: 3600,
			Path:   "/",

			Secure:   false,
			HttpOnly: true,
		})

		postHandler.PostCreate().ServeHTTP(w, req)
		assert.Equal(t, tt.want, w.Code)
	}
}
