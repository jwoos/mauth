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

	if len(positionalArgs) < 1 || len(positionalArgs) > 2 {
		usage(true)
	}

	db, err := initDB()
	checkError(err, "Error initializing database")

	command := positionalArgs[0]

	switch command {
	case "show":
		show(positionalArgs)
	case "get":
		get(positionalArgs)
	case "add":
		add()
	case "delete":
		del(positionalArgs)
	case "edit":
		edit(positionalArgs)
	default:
		fmt.Println("Invalid command")
		usage(true)
	}
}
