package service

import (
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) string {
	hashed, _ := bcrypt.GenerateFromPassword([]byte(password), 10)
	return string(hashed)
}

func CheckPassword(hashed string, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(password))
	if err != nil {
		return true
	} else {
		return false
	}
}
