package validator

import (
	"fmt"
	"net/mail"
	"regexp"
)

var (
	isValidUsername = regexp.MustCompile(`^[a-z0-9_]+$`).MatchString
	isValidFullname = regexp.MustCompile(`^[a-zA-Z\s]+$`).MatchString
)

func ValidateString(value string, minLen, maxLen int) error {
	if len(value) < minLen || len(value) > maxLen {
		return fmt.Errorf("must contain from %d - %d characters", minLen, maxLen)
	}

	return nil
}

func ValidateUsername(value string) error {
	if err := ValidateString(value, 6, 100); err != nil {
		return err
	}

	if !isValidUsername(value) {
		return fmt.Errorf("must contain only letters, digits or underscore")
	}

	return nil
}

func ValidatePassword(value string) error {
	if err := ValidateString(value, 6, 80); err != nil {
		return err
	}

	return nil
}

func ValidateFullname(value string) error {
	if err := ValidateString(value, 6, 100); err != nil {
		return err
	}

	if !isValidFullname(value) {
		return fmt.Errorf("must contain only letters or space")
	}

	return nil
}

func ValidateEmail(value string) error {
	if err := ValidateString(value, 6, 100); err != nil {
		return err
	}

	if _, err := mail.ParseAddress(value); err != nil {
		return fmt.Errorf("not a valid email")
	}

	return nil
}
