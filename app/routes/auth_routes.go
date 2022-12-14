package routes

import (
	"github.com/davidsolisdev/template-api-rest-fiber/app/controllers"
	validate "github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func AuthRoutes(app fiber.Router, validator *validate.Validate) {
	app.Post("/register-user", func(ctx *fiber.Ctx) error { return controllers.RegisterUser(ctx, validator) })

	app.Post("/register-moderator", func(ctx *fiber.Ctx) error { return controllers.RegisterModerator(ctx, validator) })

	app.Get("/email-confirmation/:id", func(ctx *fiber.Ctx) error { return controllers.EmailConfirmation(ctx, validator) })

	app.Post("/login", func(ctx *fiber.Ctx) error { return controllers.Login(ctx, validator) })

	app.Post("/send-email-recover-password", func(ctx *fiber.Ctx) error { return controllers.SendEmailRecoverPassword(ctx, validator) })
}
