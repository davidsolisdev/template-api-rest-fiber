package utils

import (
	"errors"
	"fmt"
	"os"
	"strconv"

	"github.com/golang-jwt/jwt/v4"
)

func DecodedToken(token string) (claims *ClaimsJwt, err error) {
	claims = new(ClaimsJwt)
	t, err := jwt.Parse(token, func(tk *jwt.Token) (interface{}, error) { return []byte(os.Getenv("SECRET_SIGNED_TOKEN")), nil })

	data, ok := t.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("no se puede extraer la información del token")
	}

	idfloat := data["id"].(float64)
	id, e := strconv.Atoi(fmt.Sprintf("%v", idfloat))
	if e != nil {
		return nil, e
	}

	(*claims).Id = uint(id)
	//TODO: extraer datos faltantes

	return claims, err
}
