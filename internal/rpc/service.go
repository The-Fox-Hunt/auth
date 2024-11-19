package rpc

import (
	"context"
	"github.com/The-Fox-Hunt/auth/pkg/auth"
)

type Service struct {
	auth.UnimplementedAuthServiceServer
}

func New() *Service {
	return &Service{}
}

func (s *Service) Login(ctx context.Context, in *auth.LoginIn) (*auth.LoginOut, error) {
	return &auth.LoginOut{
		Token: in.Username + in.Password,
	}, nil
}
