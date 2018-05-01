package main

import (
	"fmt"
	"database/sql"
	"strings"

	"github.com/jwoos/go_auth/otp"
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

func accountNewFromQuery(AccountQuery query) *Account {
}

func (acct *Account) String() {
	var str strings.Builder

	str.WriteString(fmt.Sprintf("ID: %d\n", acct.id))

	if acct.hashType == TYPE_HASH {
		str.WriteString("Type: HOTP\n")
	} else {
		str.WriteString("Type: TOTP\n")
	}

	str.WriteString(fmt.Sprintf("Account: %s\n", acct.account))

	str.WriteString(fmt.Sprintf("Username: %s\n", acct.username))

	return str.String()
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
