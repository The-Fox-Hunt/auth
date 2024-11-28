package rpc

import "context"

type Repo interface {
	Insert(ctx context.Context, username, password string) error
}
