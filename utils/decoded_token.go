package utils

import (
	"errors"

	"github.com/golang-jwt/jwt/v4"
)

func DecodedToken(token string) (claims ClaimsJwt, err error) {
	_, err = jwt.ParseWithClaims(token, claims, func(t *jwt.Token) (interface{}, error) {
		_, ok := t.Claims.(*ClaimsJwt)
		if !ok {
			return nil, errors.New("no se puede extraer la información del token")
		}

		return nil, nil
	})

	return claims, err
}
