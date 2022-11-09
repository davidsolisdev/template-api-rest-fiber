package main

import (
	//"github.com/davidsolisdev/template-api-rest-fiber/database"
	"github.com/davidsolisdev/template-api-rest-fiber/middlewares"
	"github.com/davidsolisdev/template-api-rest-fiber/routes"
	validate "github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/joho/godotenv"
)

// @title           Swagger Example API
// @version         1.0
// @description     This is a sample server celler server.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /api/v1

// @securityDefinitions.basic  BasicAuth
func App() (app *fiber.App) {
	// ! development only!
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	// ! development only!

	// * connect to db
	//database.ConnectMsSql()
	//database.ConnectMySql()
	//database.ConnectPostqreSql()

	// inicialization
	app = fiber.New(fiber.Config{Prefork: true})
	var validator *validate.Validate = validate.New()

	// middlewares
	app.Use(cors.New(cors.Config{}))
	app.Use(recover.New(recover.Config{}))

	// * --------------------------- routes without auth ---------------------------
	var appPublic fiber.Router = app.Group("/api")

	routes.AuthRoutes(appPublic, validator)

	// * --------------------------- routes with auth ---------------------------
	var appPrivate fiber.Router = app.Group("/api", middlewares.AuthMiddleware())

	routes.ChangeUserDataRoutes(appPrivate, validator)

	// return configured server
	return app
}
