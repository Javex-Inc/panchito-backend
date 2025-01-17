package handler

import (
	"github.com/Javex-Inc/panchito-backend/internal/factory"
	"github.com/Javex-Inc/panchito-backend/internal/model"
	"github.com/Javex-Inc/panchito-backend/internal/repository"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
)

type OrderHandler struct {
	repository repository.OrderRepository
	factory    factory.OrderFactory
}

func NewOrderHandler(client *mongo.Client) *OrderHandler {
	return &OrderHandler{
		repository: *repository.NewOrderRepository(client, "Panchito", "orders"),
		factory:    *factory.NewOrderFactory(),
	}
}

func (oh *OrderHandler) CreateOrder(c *fiber.Ctx) error {
	var orderBody model.Order

	err := c.BodyParser(&orderBody)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "internal server error",
		})
	}

	return c.Status(201).JSON(fiber.Map{
		"message": "success",
	})
}
