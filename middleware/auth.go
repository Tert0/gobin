package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/jwt/v2"
	"os"
)

func Protected() func(*fiber.Ctx) error {
	signingKey := os.Getenv("JWT_SECRET_KEY")
	if signingKey == "" {
		panic("JWT Secret Key is not set")
	}
	return jwtware.New(jwtware.Config{
		SigningKey: signingKey,
		ErrorHandler: jwtError,
		ContextKey: "token",
		TokenLookup: "header:Authorization",
	})
}

func jwtError(c *fiber.Ctx, err error) error {
	if err.Error() == "Missing or malformed JWT" {
		c.Status(400)
		return c.JSON(fiber.Map{"message": "Missing or malformed JWT"})

	} else {
		c.Status(401)
		return c.JSON(fiber.Map{"message": "Invalid or expired JWT"})
	}
}
