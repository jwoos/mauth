package main


import (
	"fmt"
	"math"
)


type HOTP struct {
	// 8-byte counter value, the moving factor [RFC4226 5.1]
	Counter uint64
	Secret string
	Length uint8
}


// [RFC4226 5.3]
func (hotp HTOP) Generate() uint {
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
	return truncatedHashNum % uint32(math.Pow10(int(hotp.Length)))
}
