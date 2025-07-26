package database

import (
	"errors"

	"github.com/google/uuid"
	"github.com/kaareskytte/golf-logger/pkg/clubs"
)

type CreateUserParams struct {
	Email    string
	Password string
}

type User struct {
	ID           string
	Email        string
	PasswordHash string
}

func (db *DB) CreateUser(params CreateUserParams) (*User, error) {
	var exists int
	err := db.conn.QueryRow("SELECT EXISTS(SELECT 1 FROM users WHERE email = ?)", params.Email).Scan(&exists)
	if err != nil {
		return nil, err
	}
	if exists == 1 {
		return nil, errors.New("email already registered")
	}

	id := uuid.New().String()
	_, err = db.conn.Exec("INSERT INTO users (id, email, password_hash) VALUES (?, ?, ?)",
		id, params.Email, params.Password)
	if err != nil {
		return nil, err
	}

	for _, club := range clubs.AllPossibleClubs {
		_, err := db.conn.Exec(`
            INSERT INTO user_clubs (id, user_id, club_name, club_type, distance, in_bag)
            VALUES (?, ?, ?, ?, ?, ?)`,
			uuid.New().String(),
			id,
			club.ClubName,
			club.ClubType,
			club.Distance,
			club.InBag,
		)
		if err != nil {
			return nil, err
		}
	}

	return &User{ID: id, Email: params.Email}, nil
}

func (db *DB) FindUserByEmail(email string) (*User, error) {
	user := &User{}
	err := db.conn.QueryRow(
		"SELECT id, email, password_hash FROM users WHERE email = ?",
		email,
	).Scan(&user.ID, &user.Email, &user.PasswordHash)
	if err != nil {
		return nil, err
	}
	return user, nil
}
