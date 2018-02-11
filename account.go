package main

import (
	"fmt"

	"otp"
)


type Account struct {
	id string
	username string
	account string
	description string
	otp OTP
}

func accountNew(id string, username string, account string, description string, otp OTP) *Account {
	var acct Account

	acct.id = id
	acct.username = account
	acct.description = description
	acct.otp = otp

	return &acct
}

func (acct *Account) show() {
	fmt.Println("%s (%s): %s", acct.account, acct.username, acct.otp.Generate())
}

func (acct *Account) save() {
	/*
	 *db.Exec(`
	 *REPLACE INTO account (
	 *    id,
	 *    account,
	 *    username,
	 *    secret,
	 *    type,
	 *    length,
	 *    timestep,
	 *    base32,
	 *    description
	 *) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
	 *`, acct.id, acct.account, acct.username, acct.otp.secret,)
	 */
}
