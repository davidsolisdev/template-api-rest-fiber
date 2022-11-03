package authController

import (
	"strings"
	"time"

	"github.com/davidsolisdev/template-api-rest-fiber/models"
	"github.com/davidsolisdev/template-api-rest-fiber/utils"
	validate "github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"gorm.io/gorm"
)

var db *gorm.DB

type BodyRegister struct {
	Name           string `json:"name" validate:"required,min=5"`
	LastName       string `json:"lastName" validate:"required,min=5"`
	Email          string `json:"email" validate:"required,min=5"`
	PassWord       string `json:"password" validate:"required,min=8"`
	RepeatPassWord string `json:"repeatPassword" validate:"required,min=8"`
}

type BodyLogin struct {
	Email    string `json:"email" validate:"required,min=5"`
	PassWord string `json:"password" validate:"required,min=8"`
}

func RegisterUser(ctx *fiber.Ctx, validator *validate.Validate) error {
	// * body validation
	var body *BodyRegister = new(BodyRegister)
	err := ctx.BodyParser(&body)
	if err != nil {
		return ctx.Status(400).SendString("Los datos enviados son incorrectos")
	}
	err = validator.Struct(body)
	if err != nil {
		return ctx.Status(400).SendString("Los datos enviados son incorrectos")
	}
	// * password comparison
	body.PassWord = strings.TrimSpace(body.PassWord)
	body.RepeatPassWord = strings.TrimSpace(body.RepeatPassWord)
	if body.PassWord != body.RepeatPassWord {
		return ctx.Status(400).SendString("Las contraseñas no son iguales")
	}

	return ctx.SendStatus(200)
}

func RegisterModerator(ctx *fiber.Ctx, validator *validate.Validate) error {
	return ctx.SendStatus(200)
}

func Login(ctx *fiber.Ctx, validator *validate.Validate) error {
	// * parse and validate body request
	var body *BodyLogin = new(BodyLogin)
	err := ctx.BodyParser(body)
	if err != nil {
		return ctx.Status(400).SendString("La información enviada es incorrecta")
	}
	errorValidate := validator.Struct(body)
	if errorValidate != nil {
		return ctx.Status(400).SendString(errorValidate.Error())
	}

	// * find user for email
	var user *models.User = new(models.User)
	tx := db.Table("users").Select("password").Where("email = ?", body.Email).First(&user)
	if tx.Error != nil {
		return ctx.Status(400).SendString("Los datos enviados son invalidos")
	}

	// * password comprobation
	comprobation := utils.CompareHashedPassword(body.PassWord, user.Password)
	if !comprobation {
		return ctx.Status(400).SendString("Los datos enviados son invalidos")
	}

	// * create token
	token, err := utils.CreateToken(utils.ClaimsJwt{
		Id: user.Id,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 10)),
		},
	})
	if err != nil {
		return ctx.Status(500).SendString("Error del servidor al crear el token")
	}

	// * create and set token on response cookie
	var cookieAuth fiber.Cookie = fiber.Cookie{
		Name:     "Authorization",
		Value:    token,
		Path:     "/",
		Expires:  time.Now().Add(time.Hour * 10),
		Secure:   false,
		HTTPOnly: false,
	}
	ctx.Cookie(&cookieAuth)

	return ctx.Status(200).SendString(token)
}

func RecoverPassword(ctx *fiber.Ctx, validator *validate.Validate) error {
	return ctx.SendStatus(200)
}

func ChangePassword(ctx *fiber.Ctx, validator *validate.Validate) error {
	return ctx.SendStatus(200)
}

func ChangeEmail(ctx *fiber.Ctx, validator *validate.Validate) error {
	return ctx.SendStatus(200)
}
