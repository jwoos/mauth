package main

import (
	"fmt"
	"database/sql"

	"otp"
)


type Account struct {
	id string
	username string
	account string
	description string
	hashType int
	otp OTP
}

const (
	TYPE_HASH = iota
	TYPE_TIME
)

func accountNew(id string, username string, account string, description string, hashType int, otp OTP) *Account {
	var acct Account

	acct.id = id
	acct.username = account
	acct.description = description
	acct.hashType = hashType
	acct.otp = otp

	return &acct
}

func (acct *Account) show() {
	fmt.Println("%s (%s): %s", acct.account, acct.username, acct.otp.Generate())
}

func (acct *Account) save() {
	id := acct.id
	account := acct.account
	username := acct.username
	hashType := acct.hashType
	secret := acct.otp.Secret
	length := acct.otp.Length

	timestep := 0

	switch acct.hashType {
	case TYPE_HASH:

	case TYPE_TIME:

	default:
	}

	if acct.hashType == TYPE_HASH {
	} else if acct.hashType == TYPE_TIME {
		db.Exec(`
		REPLACE INTO account (
			id,
			account,
			username,
			secret,
			type,
			length,
			timestep,
			base32,
			description
		) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
		`, acct.id, acct.account, acct.username, acct.otp.secret, acct.hashType,  acct.otp.Length, acct.OTP.timestep)
	}
}
