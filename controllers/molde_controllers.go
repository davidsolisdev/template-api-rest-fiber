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
// @Failure 401 {string} string "Titulo del estado"
// @Failure 503 {array} string
// @Router /v1/books/{id} [get]
func Get(ctx *fiber.Ctx) error {

	return ctx.Status(200).JSON("")
}
