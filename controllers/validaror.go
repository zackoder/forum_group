package controllers

import (
	"regexp"
	"unicode"

	"golang.org/x/crypto/bcrypt"
)

// This for valid username

func IsValidUsername(username string) bool {
	if username == "" {
		return false
	}
	last := []rune(username)[0]
	for _, c := range username {
		if c == '_' {
			if last == '_' {
				return false
			}
			last = c
			continue
		}
		if !unicode.IsLetter(c) && !unicode.IsDigit(c) {
			return false
		}
		last = c
	}
	return true
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
