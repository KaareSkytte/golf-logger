package main

import (
	"encoding/json"
	"net/http"
	"os"
	"time"

	"github.com/google/uuid"
	"github.com/kaareskytte/golf-logger/internal/auth"
)

func (cfg *apiConfig) loginHandler(w http.ResponseWriter, r *http.Request) {
	type LoginRequest struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	type LoginResponse struct {
		UserID    uuid.UUID `json:"user_id"`
		AuthToken string    `json:"auth_token"`
	}

	loginReq := LoginRequest{}

	err := json.NewDecoder(r.Body).Decode(&loginReq)
	if err != nil || loginReq.Email == "" || loginReq.Password == "" {
		respondWithError(w, http.StatusBadRequest, "Invalid login body", err)
		return
	}

	user, err := cfg.db.FindUserByEmail(loginReq.Email)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "Invalid email or password", nil)
		return
	}

	if err := auth.CheckPasswordHash(loginReq.Password, user.PasswordHash); err != nil {
		respondWithError(w, http.StatusUnauthorized, "Invalid email or password", nil)
		return
	}

	userUUID, err := uuid.Parse(user.ID)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "User ID is not a valid UUID", err)
		return
	}

	tokenSecret := os.Getenv("GOLF_LOGGER_SECRET_TOKEN")
	token, err := auth.MakeJWT(userUUID, tokenSecret, time.Hour*24)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Failed to generate token", err)
		return
	}

	resp := LoginResponse{
		UserID:    userUUID,
		AuthToken: token,
	}

	respondWithJSON(w, http.StatusOK, resp)
}
