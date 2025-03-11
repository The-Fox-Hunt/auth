package rpc

import (
	"context"
	"fmt"
	"log"

	"github.com/The-Fox-Hunt/auth/config"
	"github.com/The-Fox-Hunt/auth/internal/model"
	"github.com/The-Fox-Hunt/auth/pkg/auth"
	"github.com/golang-jwt/jwt"
)

type Service struct {
	auth.UnimplementedAuthServiceServer
	repo Repo
}

func New(r Repo) *Service {
	return &Service{
		repo: r,
	}
}

var jwtSecret []byte

func init() {
	jwtSecret, err := config.GetSecret("JWT_SECRET")
	if err != nil {
		log.Fatalf("Ошибка загрузки JWT: %v", err)
	}

	fmt.Println("JWT загружен:", jwtSecret)
}

func (s *Service) ChangePassword(ctx context.Context, in *auth.ChangePasswordIn) (*auth.ChangePasswordOut, error) {
	username, ok := ctx.Value(model.Username).(string)
	if !ok {
		return nil, fmt.Errorf("username not found in context")
	}

	err := s.repo.UpdatePassword(ctx, username, model.UserPassword{Password: in.NewPassword})

	if err != nil {
		return nil, fmt.Errorf("failed to change password: %w", err)
	}

	return &auth.ChangePasswordOut{
		Success: true,
	}, nil
}

func (s *Service) Login(ctx context.Context, in *auth.LoginIn) (*auth.LoginOut, error) {
	storedPassword, err := s.repo.GetPassword(ctx, in.Username)
	if err != nil {
		return nil, fmt.Errorf("failed to make request to get password: %w", err)
	}

	if storedPassword.Password == in.Password {
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"username": in.Username,
		})

		tokenString, err := token.SignedString([]byte(jwtSecret))

		if err != nil {
			return nil, fmt.Errorf("failed to generate JWT token: %w", err)
		}

		return &auth.LoginOut{
			Token: tokenString,
		}, nil
	}

	return nil, fmt.Errorf("invalid username or password")
}

func (s *Service) Signup(ctx context.Context, in *auth.SignupIn) (*auth.SignupOut, error) {
	err := s.repo.Insert(ctx, in.Username, in.Password)
	if err != nil {
		return &auth.SignupOut{
			Success: false,
		}, fmt.Errorf("failed to signup: %w", err)
	}
	return &auth.SignupOut{
		Success: true,
	}, nil
}
