package controllers

import (
	"os"
	"time"

	"github.com/davidsolisdev/template-api-rest-fiber/app/models"
	"github.com/davidsolisdev/template-api-rest-fiber/app/static"
	"github.com/davidsolisdev/template-api-rest-fiber/internal/database"
	"github.com/davidsolisdev/template-api-rest-fiber/pkg/utils"
	validate "github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

// @Summary Registrar Usuario
// @Description Ruta para creación de usuarios normales
// @Tags Auth
// @Accept json
// @Produce json
// @Param Body body controllers.BodyRegister true "Body peticion"
// @Success 200 {string} string
// @Failure 400 {string} string
// @Failure 500 {string} string
// @Router /register-user [post]
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

	// * check if user exists
	var user *models.User = new(models.User)
	findTx := database.DBPostgres.Table("users").Select("created").Where("email = ?", body.Email).First(user)
	if findTx.Error == nil {
		return ctx.Status(400).SendString("El usuario ya existe")
	}

	// * encrypt user password
	password, err := utils.EncryptPassword(body.RepeatPassword)
	if err != nil {
		utils.ErrorEndPoint("register user -> encrypt password", err)
		return ctx.Status(500).SendString("Error al registrar al usuario")
	}

	// * create new user without confirmed email
	var confirmedemailsecret string = uuid.NewString()
	tx := database.DBPostgres.Table("users").Create(&models.User{
		Id:                   0,
		Name:                 body.Name,
		LastName:             body.LastName,
		Email:                body.Email,
		Password:             password,
		Role:                 os.Getenv("ROLE_USER"),
		ConfirmedEmail:       false,
		ConfirmedEmailSecret: confirmedemailsecret,
		Created:              time.Now(),
		Updated:              time.Now(),
	})
	if tx.Error != nil {
		utils.ErrorEndPoint("register user", tx.Error)
		return ctx.Status(500).SendString("Error al grabar el usuario")
	}

	// * send mail of confirmation
	var mail string = static.EmailConfirmation(confirmedemailsecret)
	_, err = utils.SendEmail(&utils.NewEmail{To: body.Email, Subject: "Confirmación de correo"}, mail)
	if err != nil {
		utils.ErrorEndPoint("register user", err)
		return ctx.Status(500).SendString("Error al enviar el correo de confirmación")
	}

	return ctx.Status(200).SendString("Usuario creado con exito")
}

// @Summary Registrar Usuario Moderador
// @Description Ruta para creación de usuarios moderadores
// @Tags Auth
// @Accept json
// @Produce json
// @Param Body body controllers.BodyRegister true "Body peticion"
// @Success 200 {string} string
// @Failure 400 {string} string
// @Failure 500 {string} string
// @Router /register-moderator [post]
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

	// * check if user exists
	var user *models.User = new(models.User)
	findTx := database.DBPostgres.Table("users").Select("created").Where("email = ?", body.Email).First(&user)
	if findTx.Error == nil {
		return ctx.Status(400).SendString("El usuario ya existe")
	}

	// * encrypt user password
	password, err := utils.EncryptPassword(body.RepeatPassword)
	if err != nil {
		utils.ErrorEndPoint("register moderator -> encrypt password", err)
		return ctx.Status(500).SendString("Error al registrar al usuario")
	}

	// * create new user without confirmed email
	var confirmedemailsecret string = uuid.NewString()
	tx := database.DBPostgres.Table("users").Create(&models.User{
		Id:                   0,
		Name:                 body.Name,
		LastName:             body.LastName,
		Email:                body.Email,
		Password:             password,
		Role:                 os.Getenv("ROLE_MODERATOR"),
		ConfirmedEmail:       false,
		ConfirmedEmailSecret: confirmedemailsecret,
		Created:              time.Now(),
		Updated:              time.Now(),
	})
	if tx.Error != nil {
		utils.ErrorEndPoint("register moderator", tx.Error)
		return ctx.Status(500).SendString("Error al grabar el moderador")
	}

	// * send mail of confirmation
	var mail string = static.EmailConfirmation(confirmedemailsecret)
	_, err = utils.SendEmail(&utils.NewEmail{To: body.Email, Subject: "Confirmación de correo"}, mail)
	if err != nil {
		utils.ErrorEndPoint("register user", err)
		return ctx.Status(500).SendString("Error al enviar el correo de confirmación")
	}

	return ctx.Status(200).SendString("Moderador creado con exito")
}

// @Summary Confirmar email del usuario
// @Description Ruta para confirmar el correo electronico
// @Tags Auth
// @Accept json
// @Produce json
// @Param id path string true "code"
// @Success 200 {string} string
// @Failure 400 {string} string
// @Router /email-confirmation/{id} [get]
func EmailConfirmation(ctx *fiber.Ctx, validator *validate.Validate) error {
	// * validate params
	if len(ctx.Params("id")) < 5 {
		return ctx.Status(400).SendString("Los datos enviados son invalidos")
	}

	// * update user with confirmed email
	tx := database.DBPostgres.Table("users").Where("confirmed_email_secret = ?", ctx.Params("id")).UpdateColumn("confirmed_email", true)
	if tx.Error != nil {
		utils.ErrorEndPoint("confirmate email", tx.Error)
		return ctx.Status(400).SendString("Los datos enviados son invalidos")
	}

	return ctx.Status(200).SendString("¡Correo confirmado!")
}

// @Summary Login de usuarios
// @Description Ruta para ingresar al sistema
// @Tags Auth
// @Accept json
// @Produce json
// @Param Body body controllers.BodyLogin true "Body peticion"
// @Success 200 {string} string
// @Failure 400 {string} string
// @Failure 500 {string} string
// @Router /login [post]
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
	tx := database.DBPostgres.Table("users").Select("id", "password").Where("email = ?", body.Email).First(&user)
	if tx.Error != nil {
		return ctx.Status(400).SendString("Los datos enviados son invalidos")
	}

	// * password comprobation
	comprobation := utils.CompareHashedPassword(user.Password, body.Password)
	if !comprobation {
		return ctx.Status(400).SendString("Los datos enviados son invalidos")
	}

	// * create token
	token, err := utils.CreateToken(utils.ClaimsJwt{
		Id: uint(user.Id),
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
		Name:     "Auth",
		Value:    token,
		Path:     "/",
		Expires:  time.Now().Add(time.Hour * 10),
		Secure:   true,
		HTTPOnly: true,
	}
	ctx.Cookie(&cookieAuth)

	return ctx.Status(200).SendString(token)
}

// @Summary Recuperar contraseña
// @Description Ruta para recuperar la contraseña del usuario
// @Tags Auth
// @Accept json
// @Produce json
// @Param Body body controllers.BodyRecoverPassword true "Body peticion"
// @Success 200 {string} string
// @Failure 400 {string} string
// @Failure 500 {string} string
// @Router /recover-password [post]
func RecoverPassword(ctx *fiber.Ctx, validator *validate.Validate) error {
	// * body validation
	var body *BodyRecoverPassword = new(BodyRecoverPassword)
	err := ctx.BodyParser(&body)
	if err != nil {
		return ctx.Status(400).SendString("Los datos enviados son incorrectos")
	}
	err = validator.Struct(body)
	if err != nil {
		return ctx.Status(400).SendString("Los datos enviados son incorrectos")
	}

	// * send recover password mail
	var bodyEmail string = static.EmailRecoverPassword()
	_, err = utils.SendEmail(&utils.NewEmail{To: body.Email, Subject: "Recuperación de contraseña"}, bodyEmail)
	if err != nil {
		utils.ErrorEndPoint("recover password", err)
		return ctx.Status(500).SendString("Error interno al enviar el correo, intenta de nuevo")
	}

	return ctx.Status(200).SendString("Correo de recuperación ha sido enviado")
}

// @Summary Cambiar contraseña
// @Description Ruta para cambiar la contraseña del usuario
// @Tags Auth
// @Accept json
// @Produce json
// @Param Body body controllers.BodyChangePassword true "Body peticion"
// @Success 200 {string} string
// @Failure 400 {string} string
// @Failure 401 {string} string
// @Failure 500 {string} string
// @Router /change-password [post]
func ChangePassword(ctx *fiber.Ctx, validator *validate.Validate) error {
	// * body validation
	var body *BodyChangePassword = new(BodyChangePassword)
	err := ctx.BodyParser(&body)
	if err != nil {
		return ctx.Status(400).SendString("Los datos enviados son incorrectos")
	}
	err = validator.Struct(body)
	if err != nil {
		return ctx.Status(400).SendString("Los datos enviados son incorrectos")
	}

	// * get token of cookies
	var token string = ctx.Cookies("Auth")
	if len(token) < 8 {
		return ctx.Status(401).SendString("No has enviado tu token")
	}

	// * extract data of token
	dataToken, err := utils.DecodedToken(token)
	if err != nil {
		return ctx.Status(400).SendString("Ha ocurrido un error con tu token")
	}

	// * password comparison
	if body.Newpassword != body.RepeatNewPassword {
		return ctx.Status(400).SendString("Las contraseñas no son iguales")
	}

	// * find user for id
	var user *models.User = new(models.User)
	tx := database.DBPostgres.Table("users").Select("created", "password").Where("id = ?", dataToken.Id).First(&user)
	if tx.Error != nil {
		utils.ErrorEndPoint("chage password", err)
		return ctx.Status(400).SendString("Ha ocurrido un error")
	}

	// * encrypt new password
	password, err := utils.EncryptPassword(body.RepeatNewPassword)
	if err != nil {
		utils.ErrorEndPoint("chage password > encrypt new password", err)
		return ctx.Status(500).SendString("Ha ocurrido un error al cambiar la contraseña")
	}

	// * change last password
	tx = database.DBPostgres.Table("users").Where("id = ?", dataToken.Id).UpdateColumn("last_password", user.Password)
	if tx.Error != nil {
		utils.ErrorEndPoint("chage password > change last password", err)
		return ctx.Status(500).SendString("Ha ocurrido un error al cambiar la contraseña")
	}

	// * change password
	txNew := database.DBPostgres.Table("users").Where("id = ?", dataToken.Id).UpdateColumn("password", password)
	if txNew.Error != nil {
		utils.ErrorEndPoint("chage password > change password", err)
		return ctx.Status(500).SendString("Ha ocurrido un error al cambiar la contraseña")
	}

	return ctx.Status(200).SendString("¡Se ha cambiado la contraseña satisfactoriamente!")
}

// @Summary Cambiar email
// @Description Ruta para cambiar el correo del usuario
// @Tags Auth
// @Accept json
// @Produce json
// @Param Body body controllers.BodyChangeEmail true "Body peticion"
// @Success 200 {string} string
// @Failure 400 {string} string
// @Failure 401 {string} string
// @Failure 500 {string} string
// @Router /change-email [post]
func ChangeEmail(ctx *fiber.Ctx, validator *validate.Validate) error {
	// * body validation
	var body *BodyChangeEmail = new(BodyChangeEmail)
	err := ctx.BodyParser(&body)
	if err != nil {
		return ctx.Status(400).SendString("Los datos enviados son incorrectos")
	}
	err = validator.Struct(body)
	if err != nil {
		return ctx.Status(400).SendString("Los datos enviados son incorrectos")
	}

	// * get token of cookies
	var token string = ctx.Cookies("Auth")
	if len(token) < 8 {
		return ctx.Status(401).SendString("No has enviado tu token")
	}

	// * extract data of token
	dataToken, err := utils.DecodedToken(token)
	if err != nil {
		return ctx.Status(400).SendString("Ha ocurrido un error con tu token")
	}

	// * email comparison
	if body.NewEmail != body.RepeatNewEmail {
		return ctx.Status(400).SendString("Los emails no son iguales")
	}

	// * get user email
	var user *models.User = new(models.User)
	txU := database.DBPostgres.Table("users").Select("email").Where("id = ?", dataToken.Id).First(user)
	if txU.Error != nil {
		return ctx.Status(500).SendString("Ha ocurrido un error al cambiar la contraseña")
	}

	// * change confirmed_email
	tx := database.DBPostgres.Table("users").Where("id = ?", dataToken.Id).UpdateColumn("confirmed_email", false)
	if tx.Error != nil {
		utils.ErrorEndPoint("chage email > change confirmed_email = false", err)
		return ctx.Status(500).SendString("Ha ocurrido un error al cambiar la contraseña")
	}

	// * change confirmed_email_secret
	tx = database.DBPostgres.Table("users").Where("id = ?", dataToken.Id).UpdateColumn("confirmed_email_secret", uuid.NewString())
	if tx.Error != nil {
		utils.ErrorEndPoint("chage email > change confirmed_email_secret", err)
		return ctx.Status(500).SendString("Ha ocurrido un error al cambiar la contraseña")
	}

	// * change last email
	txLast := database.DBPostgres.Table("users").Where("id = ?", dataToken.Id).UpdateColumn("last_email", user.Email)
	if txLast.Error != nil {
		utils.ErrorEndPoint("chage email > change last_email", err)
		return ctx.Status(500).SendString("Ha ocurrido un error al cambiar la contraseña")
	}

	// * change email
	txNew := database.DBPostgres.Table("users").Where("id = ?", dataToken.Id).UpdateColumn("email", body.RepeatNewEmail)
	if txNew.Error != nil {
		utils.ErrorEndPoint("chage email > change email", err)
		return ctx.Status(500).SendString("Ha ocurrido un error al cambiar la contraseña")
	}

	return ctx.Status(200).SendString("¡Se ha cambiado el email satisfactoriamente!")
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

type BodyRecoverPassword struct {
	Email string `json:"email" validate:"required,min=5"`
}

type BodyChangePassword struct {
	Newpassword       string `json:"newPassword" validate:"required,min=8"`
	RepeatNewPassword string `json:"repeatNewPassword" validate:"required,min=8"`
}

type BodyChangeEmail struct {
	NewEmail       string `json:"newEmail" validate:"required,min=8"`
	RepeatNewEmail string `json:"repeatNewEmail" validate:"required,min=8"`
}
