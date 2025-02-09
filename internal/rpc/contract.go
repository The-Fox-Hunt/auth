package rpc

import (
	"context"

	"github.com/The-Fox-Hunt/auth/internal/model"
)

type Repo interface {
	Insert(ctx context.Context, username, password string) error
	GetPassword(ctx context.Context, username string) (model.UserPassword, error)
}
