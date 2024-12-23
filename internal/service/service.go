package service

import (
	"os"

	"github.com/Javex-Inc/panchito-backend/internal/handler"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
)

func StartServer(client *mongo.Client) {
	app := fiber.New()
	port := os.Getenv("SERVER_PORT")
	ch := handler.NewClientHandler(client)

	app.Post("/register", ch.CreateClient)

	app.Listen(":" + port)
}
