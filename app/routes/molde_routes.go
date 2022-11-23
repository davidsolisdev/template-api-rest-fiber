package routes

import (
	validate "github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func MoldeRoutes(app fiber.Router, validator *validate.Validate) {
	app.Get("/example/:id", func(ctx *fiber.Ctx) error { /*return controllers.ChangePassword(ctx, validator)*/
		return ctx.SendStatus(200)
	})

	app.Post("/example1", func(ctx *fiber.Ctx) error { /*return controllers.ChangeEmail(ctx, validator)*/
		return ctx.SendStatus(200)
	})

	app.Put("/example2/:id", func(ctx *fiber.Ctx) error { /*return controllers.ChangeEmail(ctx, validator)*/
		return ctx.SendStatus(200)
	})

	app.Delete("/example3/:id", func(ctx *fiber.Ctx) error { /*return controllers.ChangeEmail(ctx, validator)*/
		return ctx.SendStatus(200)
	})
}
