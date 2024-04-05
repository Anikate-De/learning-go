package utils

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (string, error) {
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hashedBytes), err
}

func CompareHashAndPassword(password, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}
