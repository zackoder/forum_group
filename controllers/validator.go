package controllers

import (
	"regexp"

	"golang.org/x/crypto/bcrypt"
)

func IsValidUsername(username string) bool {
	username_RGX := regexp.MustCompile(`^[a-zA-Z0-9_-]{3,20}$`)
	return username_RGX.MatchString(username)
}

// This for valid email

func IsValidEmail(email string) bool {
	re := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	regex := regexp.MustCompile(re)
	return regex.MatchString(email)
}

func HasPassowd(password string) (string, error) {
	hashpassord, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashpassord), nil
}

func isValidPassword(password string) bool {
	lowercase := regexp.MustCompile(`[a-z]`)
	uppercase := regexp.MustCompile(`[A-Z]`)
	digit := regexp.MustCompile(`\d`)
	specialChar := regexp.MustCompile(`[.\+\-\*/_/@]`)
	length := len(password) >= 8
	return lowercase.MatchString(password) &&
		uppercase.MatchString(password) &&
		digit.MatchString(password) &&
		specialChar.MatchString(password) &&
		length
}
