package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/google/uuid"
	_ "github.com/mattn/go-sqlite3"
)

type apiConfig struct {
	CurrentUserID uuid.UUID
	AuthToken     string
	APIBaseURL    string
	db            *sql.DB
}

type bag struct {
	playerID uuid.UUID
	clubs    []club
}

type club struct {
	clubName string
	clubType string
	distance int
	inBag    bool
}

var allPossibleClubs = []club{
	{"Driver", "Wood", 0, false},
	{"3-wood", "Wood", 0, false},
	{"5-wood", "Wood", 0, false},
	{"Hybrid-3", "Hybrid", 0, false},
	{"Hybrid-4", "Hybrid", 0, false},
	{"Hybrid-5", "Hybrid", 0, false},
	{"3-iron", "Iron", 0, false},
	{"4-iron", "Iron", 0, false},
	{"5-iron", "Iron", 0, false},
	{"6-iron", "Iron", 0, false},
	{"7-iron", "Iron", 0, false},
	{"8-iron", "Iron", 0, false},
	{"9-iron", "Iron", 0, false},
	{"Pitching Wedge", "Wedge", 0, false},
	{"Gap Wedge", "Wedge", 0, false},
	{"Sand Wedge", "Wedge", 0, false},
	{"Lop Wedge", "Wedge", 0, false},
	{"Putter", "Putter", 0, false},
}

func initDB() *sql.DB {
	db, err := sql.Open("sqlite3", "golf_logger.db")
	if err != nil {
		log.Fatal("cannot open db: ", err)
	}

	createTable := `
	CREATE TABLE IF NOT EXISTS users (
		id TEXT PRIMARY KEY,
		username TEXT NOT NULL UNIQUE,
		password_hash TEXT NOT NULL
	);`

	_, err = db.Exec(createTable)
	if err != nil {
		log.Fatal("cannot create users: ", err)
	}
	return db
}

func main() {
	db := initDB()

	cfg := apiConfig{
		db: db,
	}

	http.HandleFunc("/api/login", cfg.loginHandler)

	http.ListenAndServe(":8080", nil)
}
