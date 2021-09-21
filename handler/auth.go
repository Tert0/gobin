package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"log"
	"os"
	"time"
)

func Login(c *fiber.Ctx) error {
	type LoginData struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	var data LoginData

	if err := c.BodyParser(&data); err != nil {
		return c.SendStatus(400)
	}

	if data.Username != "mik" || data.Password != "1234" {
		return c.SendStatus(401)
	}

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = data.Username
	claims["admin"] = true
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	signingKey := os.Getenv("JWT_SECRET_KEY")
	if signingKey == "" {
		panic("JWT Secret Key is not set")
	}

	accessToken, err := token.SignedString(signingKey)
	if err != nil {
		log.Println(err)
		return c.SendStatus(500)
	}
	return c.JSON(fiber.Map{"access_token": accessToken})
}
