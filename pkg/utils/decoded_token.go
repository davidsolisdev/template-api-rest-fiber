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

	// * extract id
	idfloat := data["Id"].(float64)
	id, e := strconv.Atoi(fmt.Sprintf("%v", idfloat))
	if e != nil {
		return nil, e
	}

	// * extract expiration time
	_ = data["exp"].(float64)

	(*claims).Id = uint(id)
	//TODO: asignar la expiracion del token

	return claims, err
}
