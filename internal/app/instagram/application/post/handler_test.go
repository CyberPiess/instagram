package application

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
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
	body   map[string]interface{}
}

func TestPostCreate(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockPostService := mocks.NewMockpostService(ctrl)

	postHandler := NewPostHandler(mockPostService)

	mockPostService.EXPECT().CreatePost(gomock.Any(), gomock.Any()).Return(nil).AnyTimes()

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Correct format",
			args: args{method: "POST",
				cookie: customCookie{cookieName: "jwt",
					cookieValue: "testcookie"},
				body: map[string]interface{}{"image": "someImage", "description": "some_description"}},
			want: 200,
		},
		{
			name: "Method GET",
			args: args{method: "GET",
				cookie: customCookie{cookieName: "",
					cookieValue: ""},
				body: map[string]interface{}{"image": "someImage",
					"description": "some_description"}},
			want: 400,
		},
		{
			name: "Incorrect cookie",
			args: args{method: "POST",
				cookie: customCookie{cookieName: "",
					cookieValue: ""},
				body: map[string]interface{}{"image": "someImage", "description": "some_description"}},
			want: 400,
		},
	}

	for _, tt := range tests {
		body, _ := json.Marshal(tt.args.body)
		req := httptest.NewRequest(tt.args.method, "/createUser", bytes.NewReader(body))

		w := httptest.NewRecorder()
		http.SetCookie(w, &http.Cookie{
			Name:   tt.args.cookie.cookieName,
			Value:  tt.args.cookie.cookieValue,
			MaxAge: 3600,
			Path:   "/",

			Secure:   false,
			HttpOnly: true,
		})

		if tt.args.cookie.cookieName != "" {
			req.Header.Set("Cookie", w.HeaderMap["Set-Cookie"][0])
		}

		postHandler.PostCreate().ServeHTTP(w, req)
		assert.Equal(t, tt.want, w.Code)
	}
}
