package database

import (
	"errors"

	"github.com/google/uuid"
)

type CreateUserParams struct {
	Email    string
	Password string
}

type User struct {
	ID    string
	Email string
}

func (db *DB) CreateUser(params CreateUserParams) (*User, error) {
	var exists int
	err := db.conn.QueryRow("SELECT EXISTS(SELECT 1 FROM users WHERE username = ?)", params.Email).Scan(&exists)
	if err != nil {
		return nil, err
	}
	if exists == 1 {
		return nil, errors.New("email already registered")
	}

	id := uuid.New().String()
	_, err = db.conn.Exec("INSERT INTO users (id, username, password_hash) VALUES (?, ?, ?)",
		id, params.Email, params.Password)
	if err != nil {
		return nil, err
	}

	return &User{ID: id, Email: params.Email}, nil
}
