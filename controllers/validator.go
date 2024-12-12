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

func ValidPassword(passwor string) bool {
	password_RGX := regexp.MustCompile(`^(?=(.*[a-z]))(?=(.*[A-Z]))(?=(.*\d)).{8,}$`)
	return password_RGX.MatchString(passwor)
}
