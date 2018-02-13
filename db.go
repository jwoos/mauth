package main


import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)


func initDB(path string) sql.DB {
	database, err := sql.Open("sqlite3", path)
	if err {
		checkError(err, "Error opening database")
	}

	result, err := database.Exec(`CREATE TABLE IF NOT EXISTS account (
		id STRING PRIMARY KEY,
		account TEXT NOT NULL,
		username TEXT NOT NULL,
		secret TEXT NOT NULL,
		type INTEGER NOT NULL,
		length INTEGER,
		timestep INTEGER,
		base32 BOOLEAN,
		description TEXT
	)`)
	if err {
		checkError(err, "Error initiating database")
	}

	return database
}
