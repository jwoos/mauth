package main

import (
	//"otp"

	"database/sql"
	"flag"
	"fmt"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	var commandFlag string
	flag.StringVar(&commandFlag, "command", "g", "The type of action, default is (g)et with options for (a)dd, (e)dit, and (d)elete")

	var userFlag string
	flag.StringVar(&actionFlag, "user", "", "The username to specify"))

	var debugFlag bool
	flag.BoolVar(&debugFlag, "debug", false, "Specify whether in debugging mode or not")

	var verboseFlag bool
	flag.BoolVar(&verboseFlag, "verbose", false, "Print out more things than usual")

	var dbPath string
	flag.StringVar(&dbPath. "db", "./accounts.db", "The path of the file to use")

	flag.Parse()

	positionalArgs := flag.Args()

	if len(positionalArgs) != 1 {
		usage()
		os.Exit(1)
	}

	/*
	 *db, err := initDB()
	 *checkError(err, "Error initializing database")
	 */

	account := positionalArgs[0]

	switch commandFlag {
	case "get":
	case "add":
	case "delete":
	case "edit":
	}
}
