package utils

import (
	"golang.org/x/crypto/bcrypt"
)

func EncryptPassword(password string) (hashedPassword string, err error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		return "", err
	}

	hashedPassword = string(hash)
	return hashedPassword, nil
}
