package service

import (
	"os"

	"github.com/gofiber/fiber/v2"
)

type APIService struct {
	router *fiber.App
}

func NewAPIService() {
	app := fiber.New()
	port := os.Getenv("SERVER_PORT")

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello world!")
	})

	app.Listen(port)
}
