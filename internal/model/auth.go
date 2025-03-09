package model

type UserPassword struct {
	Password string `db:"password"`
}

type Key string

const (
	Username Key = "username"
)
