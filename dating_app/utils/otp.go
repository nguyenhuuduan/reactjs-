package utils

import (
	"crypto/rand"
	"log"
	"math/big"
)

func GenerateOTP() string {
	const digits = "0123456789"
	otpLength := 6
	otp := make([]byte, otpLength)

	for i := range otp {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(digits))))
		if err != nil {
			log.Println("Error generating OTP:", err)
			return ""
		}
		otp[i] = digits[num.Int64()]
	}

	return string(otp)
}
