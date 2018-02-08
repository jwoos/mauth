package main

import (
	"os"
)

func usage() {
	fmt.Println("Usage: go_auth [ACCOUNT]")
	fmt.Println("	action - get|add|delete|edit")
	fmt.Println("	action - get|add|delete|edit")
}

func checkError(err error, msg string) {
	if err != nil {
		fmt.Println(msg)
		os.Exit(1)
	}
}
