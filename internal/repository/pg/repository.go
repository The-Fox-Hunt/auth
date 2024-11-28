package pg

import (
	"context"
	"fmt"
)

type Repository struct {
	storage map[string]string
}

func NewRepository() *Repository {
	return &Repository{
		storage: make(map[string]string),
	}
}

func (r *Repository) Insert(ctx context.Context, username, password string) error {
	_, ok := r.storage[username]
	if ok {
		return fmt.Errorf("username %s already exists", username)
	}
	r.storage[username] = password
	return nil
}
