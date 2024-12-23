package handler

import (
	"github.com/Javex-Inc/panchito-backend/internal/factory"
	"github.com/Javex-Inc/panchito-backend/internal/model"
	"github.com/Javex-Inc/panchito-backend/internal/repository"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
)

type ClientHandler struct {
	repository *repository.ClientRepository
	factory    *factory.ClientFactory
}

func NewClientHandler(client *mongo.Client) *ClientHandler {
	return &ClientHandler{
		repository: repository.NewClientRepository(client, "Panchito", "clients"),
		factory:    factory.NewClientFactory(),
	}
}

func (ch *ClientHandler) CreateClient(c *fiber.Ctx) error {
	var userBody model.Client

	err := c.BodyParser(&userBody)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "internal server error",
		})
	}

	client := ch.factory.CreateClient(userBody.Name, userBody.Phone, userBody.Address)
	err = ch.repository.InsertClient(client)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "failed to insert client",
		})
	}

	return c.Status(201).JSON(fiber.Map{
		"message": "success",
	})
}
