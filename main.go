package main

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/kaareskytte/golf-logger/internal/database"
	_ "github.com/mattn/go-sqlite3"
)

type apiConfig struct {
	CurrentUserID uuid.UUID
	AuthToken     string
	APIBaseURL    string
	db            *database.DB
}

func main() {
	db := database.InitDB()

	cfg := apiConfig{
		db: db,
	}

	http.HandleFunc("/api/login", cfg.loginHandler)
	http.HandleFunc("/api/register", cfg.handlerRegister)
	http.HandleFunc("/api/bag", cfg.handlerGetBag)
	http.HandleFunc("/api/bag/club", cfg.handlerChangeClubStatus)
	http.HandleFunc("/api/bag/distance", cfg.handlerChangeClubDistance)

	http.ListenAndServe(":8080", nil)
}
