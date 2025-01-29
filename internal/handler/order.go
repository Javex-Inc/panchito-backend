package handler

import (
	"github.com/Javex-Inc/panchito-backend/internal/factory"
	"github.com/Javex-Inc/panchito-backend/internal/model"
	"github.com/Javex-Inc/panchito-backend/internal/repository"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
)

type OrderHandler struct {
	orderRepo  repository.OrderRepository
	clientRepo repository.ClientRepository
	factory    factory.OrderFactory
}

func NewOrderHandler(client *mongo.Client) *OrderHandler {
	return &OrderHandler{
		orderRepo:  *repository.NewOrderRepository(client, "Panchito", "orders"),
		clientRepo: *repository.NewClientRepository(client, "Panchito", "clients"),
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

	order := oh.factory.CreateOrder(orderBody.Number, orderBody.ClientID, orderBody.Status, orderBody.IsDelivery, orderBody.IsTakeaway, orderBody.Products, orderBody.Payment, orderBody.Total)
	err = oh.orderRepo.CreateOrder(order)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "internal server error",
		})
	}

	return c.Status(201).JSON(fiber.Map{
		"message": "success",
	})
}

func (oh *OrderHandler) GetAllOrders(c *fiber.Ctx) error {
	var orders []model.Order

	res, err := oh.orderRepo.GetAllOrders()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"code":        500,
			"message":     "internal server error",
			"description": err.Error(),
		})
	}

	orders = res

	return c.Status(200).JSON(orders)
}

func (oh *OrderHandler) GetAllFullOrders(c *fiber.Ctx) error {
	var orders []model.FullOrder

	res, err := oh.orderRepo.GetAllOrders()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"code":        500,
			"message":     "internal server error",
			"description": err.Error(),
		})
	}

	for _, order := range res {
		client, err := oh.clientRepo.FindClientByID(order.ClientID)
		if err != nil {
			return c.Status(500).JSON(fiber.Map{
				"code":        500,
				"message":     "internal server error",
				"description": err.Error(),
			})
		}

		orders = append(orders, model.FullOrder{
			ID:         order.ID,
			Number:     order.Number,
			Client:     *client,
			Status:     order.Status,
			IsDelivery: order.IsDelivery,
			IsTakeaway: order.IsTakeaway,
			Products:   order.Products,
			Payment:    order.Payment,
			Timestamp:  order.Timestamp,
			Total:      order.Total,
		})
	}

	return c.Status(200).JSON(orders)
}
