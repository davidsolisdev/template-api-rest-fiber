package utils

import (
	"golang.org/x/crypto/bcrypt"
)

func CompareHashedPassword(hashedPassword string, password string) (ok bool) {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}
