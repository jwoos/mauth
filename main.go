package main

import (
	//"otp"

	"database/sql"
	"flag"
	"fmt"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

func usage() {
	fmt.Println("Usage: otp [action]")
	fmt.Println("	action - get|add|delete|edit")
}

func checkError(err error, msg string) {
	if err != nil {
		fmt.Println(msg)
		os.Exit(1)
	}
}

func get() {
}

func main() {
	/*
	 *var typeFlag string
	 *flag.StringVar(&typeFlag, "type", "", "The type of OTP ")
	 */

	/*
	 *var actionFlag string
	 *flag.StringVar(&actionFlag, "action"))
	 */

	flag.Parse()

	positionalArgs := flag.Args()

	if len(positionalArgs) != 1 {
		usage()
		os.Exit(1)
	}

	database, err := sql.Open("sqlite3", "./otp.db")
	checkError(err, "Error opening database")

	statement, err := database.Prepare(`CREATE TABLE IF NOT EXISTS account (
		id INTEGER PRIMARY KEY,
		username TEXT,
		secret TEXT,
		type INTEGER,
		length INTEGER,
		timestep INTEGER,
		base32 BOOLEAN,
		description TEXT
	)`)
	checkError(err, "Error preparing statement")

	statement.Exec()

	switch action := positionalArgs[0]; action {
	case "get":
	case "add":
	case "delete":
	case "edit":
	default:
		usage()
		os.Exit(1)
	}
}
