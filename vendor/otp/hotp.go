package otp


import (
	"encoding/base32"
	"fmt"
	"math"
)


type HOTP struct {
	// 8-byte counter value, the moving factor [RFC4226 5.1]
	Counter uint64
	Secret string
	Base32Secret string
	Length uint8
}


func HOTPNew(counter uint64, secret string, isBase32 bool, length uint8) (*HOTP, error) {
	var hotp HOTP
	hotp.Counter = counter
	hotp.Length = length

	if isBase32 {
		hotp.Base32Secret = secret
		str, err := base32.StdEncoding.DecodeString(secret)
		if err != nil {
			return nil, err
		}

		hotp.Secret = string(str)
	} else {
		hotp.Secret = secret
	}

	return &hotp, nil
}

// [RFC4226 5.3]
func (hotp *HOTP) Generate() string {
	// step 1: Generate an HMAC-SHA-1 value Let HS = HMAC-SHA-1(K,C)
	hash := hmacSha1([]byte(hotp.Secret), uint64ToByte(hotp.Counter))

	/* step 2: Generate a 4-byte string (Dynamic Truncation)
	 * Let Sbits = DT(HS)
	 *
	 * step 3: Compute an HOTP value
	 * Let Snum  = StToNum(Sbits)
	 */
	truncatedHashNum := dynamicTruncation(hash)

	// Return D = Snum mod 10^Digit
	d := truncatedHashNum % uint(math.Pow10(int(hotp.Length)))

	lengthFormatter := fmt.Sprintf("%%0%dd", hotp.Length)

	defer func() {hotp.Counter++}()

	return fmt.Sprintf(lengthFormatter, d)
}
