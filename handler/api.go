package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"gobin/database"
	"gobin/key_generators"
	"gobin/model"
	"time"
)

func HelloWorld(c *fiber.Ctx) error  {
	return c.JSON(fiber.Map{"message": "Hello World"})
}

func Secret(c *fiber.Ctx) error  {
	return c.JSON(fiber.Map{
		"message": "This is secret!",
		"username": c.Locals("token").(*jwt.Token).Claims.(jwt.MapClaims)["username"],
	})
}

func CreatePaste(c *fiber.Ctx) error {
	var pasteData model.PasteBase
	err := c.BodyParser(&pasteData)
	if err != nil {
		c.Status(422)
		return c.JSON(fiber.Map{"error": "invalid data"})
	}

	if pasteData.Content == "" {
		c.Status(400)
		return c.JSON(fiber.Map{"error": "content should not be empty"})
	}

	if !checkContentType(pasteData.ContentType) {
		c.Status(400)
		return c.JSON(fiber.Map{"error": "invalid content type"})
	}

	id := ""
	for id == "" {
		tempID := key_generators.GetKeyGenerator("random")()
		result := database.DB.Find(&model.PasteModel{ID: tempID})
		if result.RowsAffected < 1 {
			id = tempID
		}
	}

	paste := model.PasteModel{
		ID: id,
		PasteBase: pasteData,
		Timestamp: time.Now(),
	}

	database.DB.Create(&paste)

	return c.JSON(fiber.Map{"id": paste.ID})
}

func GetPaste(c *fiber.Ctx) error {
	id := c.Params("id")

	var paste model.PasteModel
	result := database.DB.First(&paste, &model.PasteModel{ID: id})
	if result.RowsAffected < 1 {
		c.Status(404)
		return c.JSON(fiber.Map{"error": "paste not found"})
	}
	return c.JSON(fiber.Map{"id": paste.ID, "paste": paste.PasteBase, "timestamp": paste.Timestamp.Unix()})
}

func GetRawPaste(c *fiber.Ctx) error {
	id := c.Params("id")

	var paste model.PasteModel
	result := database.DB.First(&paste, &model.PasteModel{ID: id})
	if result.RowsAffected < 1 {
		c.Status(404)
		return c.JSON(fiber.Map{"error": "paste not found"})
	}
	return c.SendString(paste.Content)
}

func checkContentType(contentType string) bool {
	contentTypes := []string{"Text", "GoLang", "Python"}
	for _, i := range contentTypes {
		if i == contentType {
			return true
		}
	}
	return false
}
