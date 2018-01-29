package otp


import (
	"time"
)


/* [RFC6238 4.2]
 * TOTP = HOTP(K, T)
 * where T = is an integer and represents the
 * number of time steps between the initial counter
 * time T0 and the current Unix time
 */
type TOTP struct {
	Timestep uint8
	Hotp HOTP
}


func (totp *TOTP) New(counter uint64, secret string, base32 bool, length uint8, timestep uint8) {
	totp.Timestep = timestep
	totp.Hotp.New(counter, secret, base32, length)
}


func (totp *TOTP) Generate() string {
	totp.Hotp.Counter = (uint64(time.Now().Unix()) / uint64(totp.Timestep))

	return totp.Hotp.Generate()
}
