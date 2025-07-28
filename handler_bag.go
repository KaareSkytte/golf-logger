package main

import (
	"net/http"
	"os"

	"github.com/kaareskytte/golf-logger/internal/auth"
)

func (cfg *apiConfig) handlerGetUserBag(w http.ResponseWriter, r *http.Request) {
	token, err := auth.GetBearerToken(r.Header)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "Missing or invalid auth token", err)
		return
	}

	secret := os.Getenv("GOLF_LOGGER_SECRET_TOKEN")
	userID, err := auth.ValidateJWT(token, secret)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "Invalid token", err)
		return
	}

	clubs, err := cfg.db.GetUserBag(userID.String())
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Failed to fetch bag", err)
		return
	}

	respondWithJSON(w, http.StatusOK, clubs)
}

func (cfg *apiConfig) handlerGetFullBag(w http.ResponseWriter, r *http.Request) {
	token, err := auth.GetBearerToken(r.Header)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "Missing or invalid auth token", err)
		return
	}

	secret := os.Getenv("GOLF_LOGGER_SECRET_TOKEN")
	userID, err := auth.ValidateJWT(token, secret)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "Invalid token", err)
		return
	}

	clubs, err := cfg.db.GetFullBag(userID.String())
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Failed to fetch bag", err)
		return
	}

	respondWithJSON(w, http.StatusOK, clubs)
}
