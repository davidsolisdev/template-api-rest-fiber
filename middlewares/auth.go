package middlewares

import (
	"os"

	"github.com/gofiber/fiber/v2"
	jwtWare "github.com/gofiber/jwt/v3"
)

func AuthMiddleware() fiber.Handler {
	return jwtWare.New(jwtWare.Config{
		SigningKey:     os.Getenv("Secret_Signed_Token"),
		SigningMethod:  "HS256",
		AuthScheme:     "Bearer",
		TokenLookup:    "cookie:Authorization",
		SuccessHandler: jwtSuccess,
		ErrorHandler:   jwtError,
	})
}

func jwtSuccess(ctx *fiber.Ctx) error {
	return ctx.Next()
}

func jwtError(ctx *fiber.Ctx, err error) error {
	return ctx.Status(401).SendString("Token invalido o no has enviado el token")
}
