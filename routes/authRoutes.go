package routes

import (
	authController "github.com/davidsolisdev/template-api-rest-fiber/controllers"
	validate "github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func AuthRoutes(app fiber.Router, validator *validate.Validate) {
	app.Post("/register-user", func(ctx *fiber.Ctx) error { return authController.RegisterUser(ctx, validator) })

	app.Post("/register-moderator", func(ctx *fiber.Ctx) error { return authController.RegisterModerator(ctx, validator) })

	app.Post("/email-confirmation", func(ctx *fiber.Ctx) error { return authController.EmailConfirmation(ctx, validator) })

	app.Post("/login", func(ctx *fiber.Ctx) error { return authController.Login(ctx, validator) })
}
