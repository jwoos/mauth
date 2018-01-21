package main


import (
	"fmt"
	"math"
	"time"
)


/* TOTP = HOTP(K, T)
 * where T = is an integer and represents the
 * number of time steps between the initial counter
 * time T0 and the current Unix time
 */
type TOTP struct {
	Timestep uint8
	hotp HOTP
}


func  (totp TOTP) Generate() uint {
	totp.hotp.Counter = (uint64(time.Now().Unix()) / uint64(totp.Timestep))

	return totp.hotp.Generate()
}
