package main

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/kaareskytte/golf-logger/internal/auth"
)

func (cfg *apiConfig) handlerChangeClubStatus(w http.ResponseWriter, r *http.Request) {
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

	type Params struct {
		ClubName string `json:"clubName"`
		InBag    bool   `json:"inBag"`
	}

	var params Params
	if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
		respondWithError(w, http.StatusInternalServerError, "Invalid request body", err)
		return
	}

	err = cfg.db.UpdateClubStatus(userID.String(), params.ClubName, params.InBag)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Failed to update club", err)
		return
	}

	respondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
}

func (cfg *apiConfig) handlerChangeClubDistance(w http.ResponseWriter, r *http.Request) {
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

	type Params struct {
		ClubName string `json:"clubName"`
		Distance int    `json:"distance"`
	}

	var params Params
	if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
		respondWithError(w, http.StatusInternalServerError, "Invalid request body", err)
		return
	}

	err = cfg.db.UpdateClubDistance(userID.String(), params.ClubName, params.Distance)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Failed to update club distance", err)
		return
	}

	respondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
}
