package controllers

import "github.com/gofiber/fiber/v2"

// @Summary Metodo Get del molde
// @Description Ruta Get para solicitar datos
// @Tags books
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Param Body body controllers.BodyRegister true "Body peticion"
// @Success 200 {object} models.User
// @Failure 400 {string} string "Titulo del estado"
// @Failure 500 {array} string
// @Router /v1/books/{id} [get]
func Get(ctx *fiber.Ctx) error {

	return ctx.Status(200).JSON("")
}

// @Summary Metodo Post del molde
// @Description Ruta Post para crear datos
// @Tags books
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Param Body body controllers.BodyRegister true "Body peticion"
// @Success 200 {object} models.User
// @Failure 400 {string} string "Titulo del estado"
// @Failure 500 {array} string
// @Router /v1/books/{id} [post]
func Post(ctx *fiber.Ctx) error {

	return ctx.Status(200).JSON("")
}

// @Summary Metodo Put del molde
// @Description Ruta Put para solicitar datos
// @Tags books
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Param Body body controllers.BodyRegister true "Body peticion"
// @Success 200 {object} models.User
// @Failure 400 {string} string "Titulo del estado"
// @Failure 500 {array} string
// @Router /v1/books/{id} [put]
func Put(ctx *fiber.Ctx) error {

	return ctx.Status(200).JSON("")
}

// @Summary Metodo Delete del molde
// @Description Ruta Delete para solicitar datos
// @Tags books
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Param Body body controllers.BodyRegister true "Body peticion"
// @Success 200 {object} models.User
// @Failure 400 {string} string "Titulo del estado"
// @Failure 500 {array} string
// @Router /v1/books/{id} [delete]
func Delete(ctx *fiber.Ctx) error {

	return ctx.Status(200).JSON("")
}
