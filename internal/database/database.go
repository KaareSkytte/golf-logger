package database

import (
	"database/sql"
	"log"
)

type DB struct {
	conn *sql.DB
}

func InitDB() *DB {
	sqlDB, err := sql.Open("sqlite3", "golf_logger.db")
	if err != nil {
		log.Fatal("cannot open db: ", err)
	}

	createTable := `
	CREATE TABLE IF NOT EXISTS users (
		id TEXT PRIMARY KEY,
		username TEXT NOT NULL UNIQUE,
		password_hash TEXT NOT NULL
	);`

	_, err = sqlDB.Exec(createTable)
	if err != nil {
		log.Fatal("cannot create users: ", err)
	}
	return &DB{conn: sqlDB}
}
