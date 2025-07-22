package main

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
)

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	UserID    uuid.UUID `json:"user_id"`
	AuthToken string    `json:"auth_token"`
}

func (cfg *apiConfig) loginHandler(w http.ResponseWriter, r *http.Request) {
	req := LoginRequest{}
	json.NewDecoder(r.Body).Decode(&req)

	for _, user := range fakeUsers {
		if user.Username == req.Username && user.Password == req.Password {
			authToken := uuid.New().String()
			resp := LoginResponse{UserID: user.ID, AuthToken: authToken}
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(resp)
			return
		}
	}
	http.Error(w, "Unauthorized", http.StatusUnauthorized)
}
