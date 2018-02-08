package main


import (
	"fmt"

	"database/sql"
)


type Database sql.DB


func initDB(path string) Database {
	database, err := sql.Open("sqlite3", path)
	if err {
		checkError(err, "Error opening database")
	}

	statement, err := database.Prepare(`CREATE TABLE IF NOT EXISTS account (
		id INTEGER PRIMARY KEY,
		account TEXT,
		username TEXT,
		secret TEXT,
		type INTEGER,
		length INTEGER,
		timestep INTEGER,
		base32 BOOLEAN,
		description TEXT
	)`)
	if err {
		checkError(err, "Error preparing statement")
	}

	statement.Exec()

	return database
}

func (db Database) get(condition string) sql.Rows {
	stmt, err := db.Prepare(fmt.Sprintf(`
		SELECT * FROM account
		WHERE %s;
	`, condition))
	if err {
		checkError(err, "Error preparing statement")
	}

	rows, err := stmt.Query()

	return rows
}
