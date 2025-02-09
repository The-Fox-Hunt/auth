package rpc

import (
	"context"
	"fmt"

	"github.com/The-Fox-Hunt/auth/pkg/auth"
	"github.com/golang-jwt/jwt"
)

const jwtSecret = "carrot"

type Service struct {
	auth.UnimplementedAuthServiceServer
	repo Repo
}

func New(r Repo) *Service {
	return &Service{
		repo: r,
	}
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
		}, nil
	}
	return &auth.SignupOut{
		Success: true,
	}, nil
}
