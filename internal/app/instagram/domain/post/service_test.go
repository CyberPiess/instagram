package post

import (
	"fmt"
	"testing"
	"time"

	mock_post "github.com/CyberPiess/instagram/internal/app/instagram/domain/post/mocks"
	"github.com/CyberPiess/instagram/internal/app/instagram/infrastructure/token"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

type args struct {
	testPost Post
}

func TestCreatePost(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockPostStorage := mock_post.NewMockpostStorage(ctrl)
	mockToken := mock_post.NewMocktokenInteraction(ctrl)
	postService := NewPostService(mockPostStorage, mockToken)

	mockPostStorage.EXPECT().Create(gomock.Any()).Return(nil).AnyTimes()
	mockToken.EXPECT().VerifyToken(gomock.Any()).Return(&token.Credentials{UserId: "1"}, nil).AnyTimes()

	tests := []struct {
		name string
		args args
		want error
	}{
		{
			name: "Empty post image",
			args: args{testPost: Post{PostImage: "",
				PostDescription: "someDescription",
				CreateTime:      time.Now(),
				UserId:          1,
				AccessToken:     "someTokenString"},
			},
			want: fmt.Errorf("no post supplied"),
		},
		{
			name: "Correct data",
			args: args{testPost: Post{PostImage: "someImage",
				PostDescription: "someDescription",
				CreateTime:      time.Now(),
				UserId:          1,
				AccessToken:     "someTokenString"},
			},
			want: nil,
		},
	}

	for _, tt := range tests {
		err := postService.CreatePost(tt.args.testPost)
		assert.Equal(t, tt.want, err)
	}

}
