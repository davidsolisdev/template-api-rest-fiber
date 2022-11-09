package routes

import (
	authController "github.com/davidsolisdev/template-api-rest-fiber/controllers"
	validate "github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func ChangeUserDataRoutes(app fiber.Router, validator *validate.Validate) {
	app.Post("/change-password", func(ctx *fiber.Ctx) error { return authController.ChangePassword(ctx, validator) })

	app.Post("/change-email", func(ctx *fiber.Ctx) error { return authController.ChangeEmail(ctx, validator) })
}
