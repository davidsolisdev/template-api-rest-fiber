package main

import (
	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/davidsolisdev/template-api-rest-fiber/database"
	_ "github.com/davidsolisdev/template-api-rest-fiber/docs"
	"github.com/davidsolisdev/template-api-rest-fiber/middlewares"
	"github.com/davidsolisdev/template-api-rest-fiber/routes"
	validate "github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/joho/godotenv"
)

// @title Example API Rest with Fiber
// @version 1.0
// @description This is a sample API Rest with best practices.
// @termsOfService http://myDomain/terms

// @contact.name Support
// @contact.url http://myDomain.com/support
// @contact.email support@myDomain.com

// @license.name MIT
// @license.url https://github.com/davidsolisdev/template-api-rest-fiber/blob/main/LICENSE

// @tag.name This is the name of the tag
// @tag.description Cool Description
// @tag.docs.url https://myDomain.com/docs
// @tag.docs.description Best example documentation

// @host localhost:3005
// @schemes https
// @BasePath /api
// @accept json
// @produce json
// @securityDefinitions.apikey ApiKeyAuth
func App() (app *fiber.App) {
	// ! development only!
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}
	// ! development only!

	// * connect to db
	//database.ConnectMsSql()
	//database.ConnectMySql()
	database.ConnectPostqreSql()

	// inicialization
	app = fiber.New(fiber.Config{Prefork: true})
	var validator *validate.Validate = validate.New()

	// middlewares
	app.Use(cors.New(cors.Config{}))
	app.Use(recover.New(recover.Config{}))

	// ! Swagger
	app.Get("/docs/*", swagger.HandlerDefault)

	// * --------------------------- routes without auth ---------------------------
	var appPublic fiber.Router = app.Group("/")

	routes.AuthRoutes(appPublic, validator)

	// * --------------------------- routes with auth ---------------------------
	var appPrivate fiber.Router = app.Group("/", middlewares.AuthMiddleware())

	routes.ChangeUserDataRoutes(appPrivate, validator)

	// return configured server
	return app
}
