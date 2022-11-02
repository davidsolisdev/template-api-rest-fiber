package authControllers

import (
	"time"

	"github.com/davidsolisdev/template-api-rest-fiber/utils"
	validate "github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"gorm.io/gorm"
)

type BodyLogin struct {
	UserName string `json:"userName" validate:"required,min=2"`
	PassWord string `json:"password" validate:"required,min=8"`
}

var DB *gorm.DB

func Login(ctx *fiber.Ctx, validator *validate.Validate) error {
	//Parse and validate body request
	var body *BodyLogin = new(BodyLogin)
	err := ctx.BodyParser(body)
	if err != nil {
		return ctx.Status(400).SendString("La información enviada es incorrecta")
	}
	errorValidate := validator.Struct(body)
	if errorValidate != nil {
		return ctx.Status(400).SendString(errorValidate.Error())
	}

	// find user for userName
	//user := DB.Where().Find()

	// Create token
	token, err := utils.CreateToken(utils.ClaimsJwt{
		Id: "",
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 10)),
		},
	})
	if err != nil {
		return ctx.Status(500).SendString("Error del servidor al crear el token")
	}

	// Create and set token on response cookie
	var cookieAuth fiber.Cookie = fiber.Cookie{
		Name:     "Authentication",
		Value:    token,
		Path:     "/",
		Expires:  time.Now().Add(time.Hour * 10),
		Secure:   false,
		HTTPOnly: false,
	}
	ctx.Cookie(&cookieAuth)

	return ctx.SendStatus(200)
}
