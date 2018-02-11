package otp


import (
	"encoding/binary"
	"crypto/sha1"
	"crypto/hmac"
)


type OTP interface {
	Generate() string
}


func uint64ToByte(num uint64) []byte {
	byteArr := make([]byte, 8)
	binary.BigEndian.PutUint64(byteArr, num)

	return byteArr
}

func hmacSha1(secret []byte, counter []byte) []byte {
	hmac := hmac.New(sha1.New, secret)
	hmac.Write(counter)

	return hmac.Sum(nil)
}

// [RFC4226 5.4]
func dynamicTruncation(hash []byte) uint {
	// sha1 is 20 bytes
	offset := uint(hash[19] & 0xF)
	truncatedVal := (uint((hash[offset])  & 0x7F) << 24) |
		((uint(hash[offset + 1]) & 0xFF) << 16) |
		((uint(hash[offset + 2]) & 0xFF) <<  8) |
		(uint(hash[offset + 3]) & 0xFF)

	return truncatedVal
}
