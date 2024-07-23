package utils

import (
	"crypto/rand"
	"fmt"
	"log"
	"math/big"
)

func GenerateOTP() (string, error) {
	min := int64(100000)
	max := int64(999999)

	randomInt, err := rand.Int(rand.Reader, new(big.Int).Sub(big.NewInt(max), big.NewInt(min)))
	if err != nil {
		return "", err
	}
	otpValue := randomInt.Int64() + min

	otp := fmt.Sprintf("%05d", otpValue)
	log.Println("otpValue: ", otp)
	return otp, nil
}
