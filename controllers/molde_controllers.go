package controllers

import "github.com/gofiber/fiber/v2"

// @Summary Metodo Get del molde
// @Description Ruta Get para solicitar datos
// @Tags books
// @Accept json
// @Produce json
// @Param id path int true "Book ID"
// @Success 200 {object} ResponseHTTP{data=[]models.example}
// @Failure 401 {object} ResponseHTTP{}
// @Failure 503 {object} ResponseHTTP{}
// @Router /v1/books/{id} [get]
func Get(ctx *fiber.Ctx) error {

	return ctx.Status(200).JSON("")
}
