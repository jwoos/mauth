package main

import (
	"fmt"
	"strconv"

	"otp"

	"satori/go.uuid"
)


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

func get() {
}

func edit() {

}

func del() {

}
