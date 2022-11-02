package main

import (
	//"github.com/davidsolisdev/template-api-rest-fiber/middlewares"
	"github.com/davidsolisdev/template-api-rest-fiber/routes"
	validate "github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func App() (app *fiber.App) {
	//inicialization
	app = fiber.New(fiber.Config{Prefork: true})
	var validator *validate.Validate = validate.New()

	//middlewares
	app.Use(cors.New(cors.Config{}))
	app.Use(recover.New(recover.Config{}))

	//routes without auth
	var appPublic fiber.Router = app.Group("/api")

	routes.AuthRoutes(appPublic, validator)

	//routes with auth
	//var appPrivate fiber.Router = app.Group("/api", middlewares.AuthMiddleware())

	//return configured server
	return app
}
