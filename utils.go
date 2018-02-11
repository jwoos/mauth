package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)


func usage() {
	fmt.Println("Usage: go_auth [--debug] [--db=DB] [ACTION]")
	fmt.Println("	ACTION: get|add|delete|edit")
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
