package rpc

//go:generate mockgen -destination=mock_contract_test.go -package=${GOPACKAGE} -source=contract.go

import (
	"context"

	"github.com/The-Fox-Hunt/auth/internal/model"
)

type Repo interface {
	Insert(ctx context.Context, username, password string) error
	GetPassword(ctx context.Context, username string) (model.UserPassword, error)
	UpdatePassword(ctx context.Context, username string, newPassword model.UserPassword) error
}
