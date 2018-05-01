package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/jwoos/go_auth/otp"

	"github.com/satori/go.uuid"
)


func show([]string args) {
	if len(args) == 2 {
		rows, err := db.Query(`SELECT * FROM accounts WHERE id=$1`, args[1])
		checkError(err)

		defer rows.Close()

		s := AccountQuery{}

		err := rows.Scan(&s)
		checkError(err)

		fmt.Println(s)
	} else {
		rows, err := db.Query(`SELECT * FROM accounts`)
		checkError(err)

		defer rows.Close()

		var str strings.Builder

		for rows.Next() {
			s := AccountQuery{}

			err := rows.Scan(&s)
			checkError(err)

			str.WriteString(s.String())
		}

		fmt.Println(str)
	}
}

func get([]string args) {
	if len(args) < 2 {
		usage(true)
	}

	rows, err := db.Query(`SELECT * FROM accounts WHERE id=$1`, args[1])
	checkError(err)

	defer rows.Close()

	s := AccountQuery{}

	err := rows.Scan(&s)
	checkError(err)
}

func add() {
	accountName := read("Account name: ")
	username := read("Username: ")
	desc := read("Description: ")
	typeOTP := read("Type (h/t): ")
	length, _ := strconv.Atoi(read("Length: "))
	secret := read("Secret: ")
	base32 := false
	if read("Base32 (y/n): ") == "y" {
		base32 = true
	}

	var otp *OTP
	if typeOTP == "h" {
		counter, _ := strconv.Atoi(read("Counter: "))
		otp = HOTPNew(0, secret, base32, length)
	} else if typeOTP == "t" {
		timeStep, _ := strconv.Atoi(read("Timestep: "))
		otp = TOTPNew(secret, base32, length, timeStep)
	}

	uuid, err := uuid.NewV4()
	account := accountNew(uuid, username, accountName, description, hashType, otp)

	// commit
	account.save();
}

func edit() {

}

func del() {

}
