package handler

import (
	"github.com/Javex-Inc/panchito-backend/internal/factory"
	"github.com/Javex-Inc/panchito-backend/internal/model"
	"github.com/Javex-Inc/panchito-backend/internal/repository"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
)

type ProductHandler struct {
	repository *repository.ProductRepository
	factory    *factory.ProductFactory
}

func NewProductHandler(client *mongo.Client) *ProductHandler {
	return &ProductHandler{
		repository: repository.NewProductRepository(client, "Panchito", "products"),
		factory:    factory.NewProductFactory(),
	}
}

func (ph *ProductHandler) CreateProduct(c *fiber.Ctx) error {
	var productBody model.Product

	err := c.BodyParser(&productBody)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "internal server error",
		})
	}

	product := ph.factory.CreateProduct(productBody.Name, productBody.Description, productBody.Image, productBody.Category, productBody.Price, productBody.IsOnOffer)
	err = ph.repository.InsertProduct(product)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "failed to insert client",
		})
	}

	return c.Status(201).JSON(fiber.Map{
		"message": "success",
	})
}
