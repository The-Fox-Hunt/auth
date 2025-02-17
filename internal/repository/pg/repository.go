package pg

import (
	"context"
	"fmt"
	"log"

	"github.com/The-Fox-Hunt/auth/internal/model"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
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
	query := "INSERT INTO participants (user_uuid, login, password) VALUES($1, $2, $3)"
	_, err := r.conn.ExecContext(ctx, query, "ef969bba-6192-4175-b63e-832481e6b563", username, password)
	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) GetPassword(ctx context.Context, username string) (model.UserPassword, error) {
	query := "SELECT password FROM participants WHERE login = $1"
	var storedPassword model.UserPassword
	err := r.conn.GetContext(ctx, &storedPassword, query, username)

	if err != nil {
		return model.UserPassword{}, err
	}

	return storedPassword, nil
}

func (r *Repository) UpdatePassword(ctx context.Context, username string, newPassword model.UserPassword) error {
	query := "UPDATE participants SET password = $1 WHERE login = $2 RETURNING password"
	var updatedPassword string
	err := r.conn.GetContext(ctx, &updatedPassword, query, newPassword.Password, username)

	if err != nil {
		return fmt.Errorf("cant changed password: %v", err)
	}

	return nil
}
