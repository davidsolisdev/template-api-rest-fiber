package utils

import (
	"errors"
	"os"

	"github.com/golang-jwt/jwt/v4"
)

func DecodedToken(token string) (claims *ClaimsJwt, err error) {
	t, err := jwt.Parse(token, func(tk *jwt.Token) (interface{}, error) { return []byte(os.Getenv("SECRET_SIGNED_TOKEN")), nil })
	claims, ok := t.Claims.(*ClaimsJwt)
	if !ok {
		return nil, errors.New("no se puede extraer la información del token")
	}
	return claims, err
}
