package pg

import (
	"context"
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
)

type Repository struct {
	storage map[string]string
	conn    *sqlx.DB
}

func NewRepository() *Repository {
	connectCmd := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=disable",
		"master", "master", "master", "localhost", "3015")

	db, err := sqlx.Connect("postgres", connectCmd)
	if err != nil {
		log.Fatal(err)
	}
	return &Repository{
		storage: make(map[string]string),
		conn:    db,
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
