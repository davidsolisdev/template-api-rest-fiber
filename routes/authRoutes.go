package routes

import (
	authController "github.com/davidsolisdev/template-api-rest-fiber/controllers"
	validate "github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func AuthRoutes(app fiber.Router, validator *validate.Validate) {
	app.Get("/login", func(ctx *fiber.Ctx) error { return authController.Login(ctx, validator) })
}
