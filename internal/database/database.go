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
		email TEXT NOT NULL UNIQUE,
		password_hash TEXT NOT NULL
	);`

	_, err = sqlDB.Exec(createTable)
	if err != nil {
		log.Fatal("cannot create users: ", err)
	}

	createTable = `
	CREATE TABLE IF NOT EXISTS user_clubs (
		id TEXT PRIMARY KEY,         
		user_id TEXT NOT NULL,       
		club_name TEXT NOT NULL,     
		club_type TEXT NOT NULL,     
		distance INTEGER NOT NULL,   
		in_bag BOOLEAN NOT NULL,     
		UNIQUE(user_id, club_name),  
		FOREIGN KEY(user_id) REFERENCES users(id)
	);`

	_, err = sqlDB.Exec(createTable)
	if err != nil {
		log.Fatal("cannot create user_clubs: ", err)
	}
	return &DB{conn: sqlDB}
}
