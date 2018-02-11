package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	var debugFlag bool
	flag.BoolVar(&debugFlag, "debug", false, "Specify whether in debugging mode or not")

	var dbPath string
	flag.StringVar(&dbPath. "db", "./accounts.db", "The path of the file to use")

	flag.Parse()

	positionalArgs := flag.Args()

	if len(positionalArgs) != 1 {
		usage()
		os.Exit(1)
	}

	db, err := initDB()
	checkError(err, "Error initializing database")

	command := positionalArgs[0]

	switch command {
	case "get":
		get()
	case "add":
		add()
	case "delete":
		del()
	case "edit":
		edit()
	default:
		fmt.Println("Invalid command")
		usage()
		os.Exit(1)
	}
}
