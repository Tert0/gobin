package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"gobin/handler"
	"gobin/middleware"
	"time"
)

func SetupRoutes(app *fiber.App)  {
	api := app.Group("/api/v1/", logger.New())
	api.Get("/", handler.HelloWorld)
	api.Get("/secret", middleware.Protected(), handler.Secret)

	api.Post("/paste", limiter.New(limiter.Config{
		Max: 5,
		Expiration: time.Minute * 10,
	}), handler.CreatePaste)

	getPasteLimiter := limiter.New(limiter.Config{
		Max: 50,
		Expiration: time.Minute * 10,
	})

	api.Get("/paste/:id", getPasteLimiter, handler.GetPaste)
	api.Get("/paste/:id/raw", getPasteLimiter, handler.GetRawPaste)

	authApi := api.Group("/auth")
	authApi.Post("/login", handler.Login)
}
