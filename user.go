package main

import (
	"github.com/google/uuid"
)

type User struct {
	ID       uuid.UUID
	Username string
	Password string
}

var fakeUsers = []User{
	{
		ID:       uuid.New(),
		Username: "alice",
		Password: "password123",
	},
}
