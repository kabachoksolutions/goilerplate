package authbp

import (
	"errors"
	"unicode"
)

// ValidatePassword checks the strength of a password by ensuring it meets certain criteria:
//
//   - password is not empty;
//
//   - password is at least 8 characters long;
//
//   - password contains at least one uppercase letter;
//
//   - password contains at least one lowercase letter;
//
//   - password contains at least one digit;
//
//   - password contains at least one special character (punctuation or symbol).
func ValidatePassword(password string) error {
	if password == "" {
		return errors.New("validator: password is required")
	}
	if len(password) < 8 {
		return errors.New("validator: password must be at least 8 characters long")
	}

	var (
		hasUpper   bool
		hasLower   bool
		hasDigit   bool
		hasSpecial bool
	)

	for _, char := range password {
		switch {
		case unicode.IsUpper(char):
			hasUpper = true
		case unicode.IsLower(char):
			hasLower = true
		case unicode.IsDigit(char):
			hasDigit = true
		case unicode.IsPunct(char) || unicode.IsSymbol(char):
			hasSpecial = true
		}
	}

	if !hasUpper || !hasLower || !hasDigit || !hasSpecial {
		return errors.New("validator: password must contain at least one uppercase letter, one lowercase letter, one digit, and one special character")
	}

	return nil
}
