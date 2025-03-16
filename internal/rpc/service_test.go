package rpc

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/The-Fox-Hunt/auth/internal/model"
	"github.com/The-Fox-Hunt/auth/pkg/auth"
	"github.com/golang-jwt/jwt"
	gomock "github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	os.Setenv("JWT_SECRET", "test-secret")
	jwtSecret = []byte("test-secret") 

	exitCode := m.Run() // Запускаем тесты
	os.Exit(exitCode)   // Выходим с кодом тестов
}

func TestService_Singup(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	//t.Setenv("JWT_SECRET", "test-secret")

	t.Run("sould return succes", func(t *testing.T) {
		ctx := context.Background()
		mockRepo := NewMockRepo(ctrl)
		s := New(mockRepo)

		mockRepo.EXPECT().Insert(ctx, "testuser", "testpass").Return(nil).Times(1)

		resp, err := s.Signup(ctx, &auth.SignupIn{
			Username: "testuser",
			Password: "testpass",
		})

		assert.NoError(t, err)
		assert.True(t, resp.Success)
	})

	t.Run("should return error if Insert fails", func(t *testing.T) {
		ctx := context.Background()
		mockRepo := NewMockRepo(ctrl)
		s := New(mockRepo)

		mockRepo.EXPECT().Insert(ctx, "testuser5", "testpass").Return(errors.New("DB error")).Times(1)

		resp, err := s.Signup(ctx, &auth.SignupIn{
			Username: "testuser5",
			Password: "testpass",
		})

		assert.NotNil(t, resp)
		assert.False(t, resp.Success)
		assert.ErrorContains(t, err, "DB error")
	})
}

func TestService_Login(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	t.Setenv("JWT_SECRET", "test-secret")

	t.Run("should return JWT token for correct password", func(t *testing.T) {
		ctx := context.Background()
		mockRepo := NewMockRepo(ctrl)
		s := New(mockRepo)

		mockRepo.EXPECT().GetPassword(ctx, "testuser").Return(model.UserPassword{Password: "testpass"}, nil).Times(1)

		resp, err := s.Login(ctx, &auth.LoginIn{
			Username: "testuser",
			Password: "testpass",
		})

		assert.NoError(t, err)
		assert.NotEmpty(t, resp.Token)

		// Декодируем токен, чтобы проверить username
		token, _ := jwt.Parse(resp.Token, func(token *jwt.Token) (interface{}, error) {
			return []byte("test-secret"), nil
		})
		claims, _ := token.Claims.(jwt.MapClaims)
		assert.Equal(t, "testuser", claims["username"])
	})

	t.Run("should return error if user not found", func(t *testing.T) {
		ctx := context.Background()
		mockRepo := NewMockRepo(ctrl)
		s := New(mockRepo)

		mockRepo.EXPECT().GetPassword(ctx, "testuser").Return(model.UserPassword{}, errors.New("user not found")).Times(1)

		resp, err := s.Login(ctx, &auth.LoginIn{
			Username: "testuser",
			Password: "testpass",
		})

		assert.Nil(t, resp)
		assert.ErrorContains(t, err, "failed to make request to get password")
	})

	t.Run("should return error if password is incorrect", func(t *testing.T) {
		ctx := context.Background()
		mockRepo := NewMockRepo(ctrl)
		s := New(mockRepo)

		mockRepo.EXPECT().GetPassword(ctx, "testuser").Return(model.UserPassword{Password: "correctpass"}, nil).Times(1)

		t.Setenv("JWT_SECRET", "")
		jwtSecret = nil

		resp, err := s.Login(ctx, &auth.LoginIn{
			Username: "testuser",
			Password: "wrongpass",
		})

		assert.Nil(t, resp)
		assert.EqualError(t, err, "invalid username or password")
	})

	t.Run("should return error if JWT token generation fails", func(t *testing.T) {
		ctx := context.Background()
		mockRepo := NewMockRepo(ctrl)
		s := New(mockRepo)
		
		mockRepo.EXPECT().GetPassword(ctx, "testuser").Return(model.UserPassword{Password: "testpass"}, nil).Times(1)
		
		jwtSecret = nil

		resp, err := s.Login(ctx, &auth.LoginIn{
			Username: "testuser",
			Password: "testpass",
		})

		// ✅ Ожидаем `nil` и ошибку
		assert.ErrorContains(t, err, "failed to generate JWT token")
		assert.Nil(t, resp)
	})
}
