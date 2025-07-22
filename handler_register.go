package main

import (
	"encoding/json"
	"net/http"

	"github.com/kaareskytte/golf-logger/internal/auth"
	"github.com/kaareskytte/golf-logger/internal/database"
)

func (cfg *apiConfig) handlerRegister(w http.ResponseWriter, r *http.Request) {
	type Params struct {
		Email    string
		Password string
	}

	userParams := Params{}
	err := json.NewDecoder(r.Body).Decode(&userParams)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request body (must be JSON with email and password)", err)
		return
	}

	if userParams.Email == "" || userParams.Password == "" {
		respondWithError(w, http.StatusBadRequest, "Email or Password must both be provided", nil)
		return
	}

	hashed, err := auth.HashPassword(userParams.Password)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Failed to hash password", err)
		return
	}

	user, err := cfg.db.CreateUser(database.CreateUserParams{
		Email:    userParams.Email,
		Password: hashed,
	})
	if err != nil {
		msg := "Failed to create user"

		if err.Error() == "email already registered" {
			msg = "Email already registered"
			respondWithError(w, http.StatusConflict, msg, nil)
			return
		}
		respondWithError(w, http.StatusInternalServerError, msg, err)
		return
	}

	respondWithJSON(w, http.StatusCreated, struct {
		ID    string `json:"id"`
		Email string `json:"email"`
	}{
		ID:    user.ID,
		Email: user.Email,
	})
}
