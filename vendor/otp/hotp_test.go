package otp

import (
	"testing"
)

func TestHotp(t *testing.T) {
	hotp, _ := HOTPNew(0, "12345678901234567890", false, 6)
	expected := []string{
		"755224",
		"287082",
		"359152",
		"969429",
		"338314",
	}

	for _, val := range expected {
		if val != hotp.Generate() {
			t.Fail()
		}
	}
}

func TestHotpBase32(t *testing.T) {
	hotp, _ := HOTPNew(0, "ZVZG5UZU4D7MY4DH", true, 6)
	expected := []string{
		"269371",
		"940502",
		"877748",
		"197256",
		"516658",
	}

	for _, val := range expected {
		if val != hotp.Generate() {
			t.Fail()
		}
	}
}
