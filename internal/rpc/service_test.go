package rpc

import (
	"context"
	"testing"

	"github.com/The-Fox-Hunt/auth/pkg/auth"
	gomock "github.com/golang/mock/gomock"
)

func TestService_Singup(t *testing.T) {
	ctrl := gomock.NewController(t)

	t.Run("sould ok", func(t *testing.T) {

		ctx := context.Background()
		mockRepo := NewMockRepo(ctrl)
		s := New(mockRepo)
		s.Signup(ctx, &auth.SignupIn{
			Username : "test",
			Password : "test",
		})

	})
}