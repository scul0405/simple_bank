package util

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

// HashPassword returns the bcrypt hash of the password
func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("failed to hash passowrd: %w", err)
	}

	return string(hashedPassword), nil
}

// CheckPassword check if the provided password is correct or not
func CheckPassword(hashedPassword, password string) (error) {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}