package middlewares

import (
	"os"

	"github.com/gofiber/fiber/v2"
	jwtWare "github.com/gofiber/jwt/v3"
)

func AuthMiddleware() fiber.Handler {
	return jwtWare.New(jwtWare.Config{
		SigningKey:     []byte(os.Getenv("SECRET_SIGNED_TOKEN")),
		SigningMethod:  "HS256",
		AuthScheme:     "Bearer",
		TokenLookup:    "cookie:Auth",
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
