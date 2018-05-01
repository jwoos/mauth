package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)


func usage(bool exit) {
	fmt.Println("Usage: go_auth [--debug] [ACTION]")
	fmt.Println("	ACTION: show|get|add|delete|edit")

	if exit {
		os.Exit(1)
	}
}

func checkError(err error, msg string) {
	if err != nil {
		fmt.Println(msg)
		os.Exit(1)
	}
}

func read(msg string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(msg)
	text, _ := reader.ReadString('\n')
	text = strings.replace(text, "\n", "", -1)
}
