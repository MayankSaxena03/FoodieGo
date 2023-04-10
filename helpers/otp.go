package helpers

import (
	"fmt"
	"math/rand"
	"regexp"
)

func RandomOTP(length int) string {
	//six digit otp generator
	otp := ""
	for i := 0; i < length; i++ {
		otp += fmt.Sprint(rand.Intn(9) + 1)
	}
	return otp
}

func ValidatePhoneNumber(phoneNumber string) bool {
	pattern := `^\d{10}$`

	regex := regexp.MustCompile(pattern)

	return regex.MatchString(phoneNumber)
}

func ValidateEmail(email string) bool {
	pattern := `^[\w-\.]+@([\w-]+\.)+[\w-]{2,4}$`

	regex := regexp.MustCompile(pattern)

	return regex.MatchString(email)
}
