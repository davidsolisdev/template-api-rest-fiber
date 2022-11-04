package authController

import (
	"os"
	"time"

	"github.com/davidsolisdev/template-api-rest-fiber/models"
	"github.com/davidsolisdev/template-api-rest-fiber/static"
	"github.com/davidsolisdev/template-api-rest-fiber/utils"
	validate "github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

var db *gorm.DB

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
	if body.Password != body.RepeatPassword {
		return ctx.Status(400).SendString("Las contraseñas no son iguales")
	}

	// * create new user without confirmed email
	tx := db.Table("users").Create(&models.User{
		Id:                   0,
		Name:                 body.Name,
		LastName:             body.LastName,
		Email:                body.Email,
		Password:             body.RepeatPassword,
		Role:                 os.Getenv("role_user"),
		ConfirmedEmail:       false,
		ConfirmedEmailSecret: uuid.NewString(),
	})
	if tx.Error != nil {
		utils.ErrorEndPoint("register user", tx.Error)
		return ctx.Status(500).SendString("Error al grabar el usuario")
	}

	// * send mail of confirmation
	var mail string = static.EmailConfirmation()
	_, err = utils.SendEmail(&utils.NewEmail{To: body.Email, Subject: "Confirmación de correo"}, mail)
	if err != nil {
		utils.ErrorEndPoint("register user", err)
		return ctx.Status(500).SendString("Error al enviar el correo de confirmación")
	}

	return ctx.Status(200).SendString("Usuario creado con exito")
}

func RegisterModerator(ctx *fiber.Ctx, validator *validate.Validate) error {
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
	if body.Password != body.RepeatPassword {
		return ctx.Status(400).SendString("Las contraseñas no son iguales")
	}

	// * create new user without confirmed email
	tx := db.Table("users").Create(&models.User{
		Id:                   0,
		Name:                 body.Name,
		LastName:             body.LastName,
		Email:                body.Email,
		Password:             body.RepeatPassword,
		Role:                 os.Getenv("role_moderator"),
		ConfirmedEmail:       false,
		ConfirmedEmailSecret: uuid.NewString(),
	})
	if tx.Error != nil {
		utils.ErrorEndPoint("register moderator", tx.Error)
		return ctx.Status(500).SendString("Error al grabar el moderador")
	}

	// * send mail of confirmation
	var mail string = static.EmailConfirmation()
	_, err = utils.SendEmail(&utils.NewEmail{To: body.Email, Subject: "Confirmación de correo"}, mail)
	if err != nil {
		utils.ErrorEndPoint("register user", err)
		return ctx.Status(500).SendString("Error al enviar el correo de confirmación")
	}

	return ctx.Status(200).SendString("Moderador creado con exito")
}

func EmailConfirmation(ctx *fiber.Ctx, validator *validate.Validate) error {
	// * validate body
	var body *ParamsConfirmMail = new(ParamsConfirmMail)
	err := ctx.BodyParser(&body)
	if err != nil {
		return ctx.Status(400).SendString("Los datos enviados son invalidos")
	}

	// * find user for email
	var user *models.User = new(models.User)
	tx := db.Table("users").Select("id", "confirmed_email_secret").Where("email = ?", body.Email).First(user)
	if tx.Error != nil {
		return ctx.Status(400).SendString("Los datos enviados son invalidos")
	}

	// * comparate secrets
	if user.ConfirmedEmailSecret != body.Id {
		return ctx.Status(400).SendString("Los datos enviados son invalidos")
	}

	// * update user with confirmed email
	txUpdate := db.Where("id = ?", user.Id).UpdateColumn("confirmed_email", true)
	if tx.Error != nil {
		utils.ErrorEndPoint("confirmate email", txUpdate.Error)
		return ctx.Status(400).SendString("Los datos enviados son invalidos")
	}

	return ctx.Status(200).SendString("¡Correo confirmado!")
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
	comprobation := utils.CompareHashedPassword(body.Password, user.Password)
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
		utils.ErrorEndPoint("login", err)
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

type BodyRegister struct {
	Name           string `json:"name" validate:"required,min=5"`
	LastName       string `json:"lastName" validate:"required,min=5"`
	Email          string `json:"email" validate:"required,min=5"`
	Password       string `json:"password" validate:"required,min=8"`
	RepeatPassword string `json:"repeatPassword" validate:"required,min=8"`
}

type BodyLogin struct {
	Email    string `json:"email" validate:"required,min=5"`
	Password string `json:"password" validate:"required,min=8"`
}

type ParamsConfirmMail struct {
	Email string
	Id    string
}
