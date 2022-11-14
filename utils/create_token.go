package utils

import (
	"os"

	"github.com/golang-jwt/jwt/v4"
)

type ClaimsJwt struct {
	Id uint
	jwt.RegisteredClaims
}

func CreateToken(claims ClaimsJwt) (token string, err error) {
	var tokenIncompleto *jwt.Token = jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	token, err = tokenIncompleto.SignedString([]byte(os.Getenv("SECRET_SIGNED_TOKEN")))
	if err != nil {
		return "", err
	}
	return token, nil
}
