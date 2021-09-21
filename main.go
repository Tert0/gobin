package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
	"gobin/database"
	"gobin/key_generators"
	"gobin/router"
	"log"
	"os"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Failed to load .env")
	}

	serverHost := os.Getenv("HOST")
	if serverHost == "" {
		serverHost = "0.0.0.0"
	}
	serverPort := os.Getenv("PORT")
	if serverPort == "" {
		serverPort = "3000"
	}


	database.ConnectDB()

	key_generators.RegisterKeyGenerator("random", key_generators.GenerateRandomKey)

	app := fiber.New(fiber.Config{
		AppName: "GoBin",
		Prefork: false,
	})

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "*",
		AllowMethods: "*",
	}))
	app.Use(logger.New())

	router.SetupRoutes(app)


	log.Fatal(app.Listen(fmt.Sprintf("%s:%s", serverHost, serverPort)))
}
