package utils

import (
	//"errors"
	"os"

	"github.com/golang-jwt/jwt/v4"
)

func DecodedToken(token string) (claims jwt.Claims, err error) {
	_, err = jwt.ParseWithClaims(token, claims, func(t *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("SECRET_SIGNED_TOKEN")), nil
	})

	/*claims, ok := t.Claims.(*ClaimsJwt)
	if !ok {
		return nil, errors.New("no se puede extraer la información del token")
	}*/

	return claims, err
}
