package rpc

import (
	"context"

	"github.com/The-Fox-Hunt/auth/pkg/auth"
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

func (s *Service) Login(ctx context.Context, in *auth.LoginIn) (*auth.LoginOut, error) {
	return &auth.LoginOut{
		Token: in.Username + in.Password,
	}, nil
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
