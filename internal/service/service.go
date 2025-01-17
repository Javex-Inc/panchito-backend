package service

import (
	"os"

	"github.com/Javex-Inc/panchito-backend/internal/handler"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"go.mongodb.org/mongo-driver/mongo"
)

func StartServer(client *mongo.Client) {
	app := fiber.New()
	port := os.Getenv("SERVER_PORT")

	ch := handler.NewClientHandler(client)
	ph := handler.NewProductHandler(client)
	oh := handler.NewOrderHandler(client)

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	app.Post("/register/client", ch.CreateClient)
	app.Post("/register/product", ph.CreateProduct)
	app.Post("register/order", oh.CreateOrder)
	app.Get("/products", ph.GetAllProducts)

	app.Listen(":" + port)
}
