package handler

import (
	"context"
	"os"
	"strconv"

	"github.com/Javex-Inc/panchito-backend/internal/factory"
	"github.com/Javex-Inc/panchito-backend/internal/repository"
	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
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

func (ph *ProductHandler) GetAllProducts(c *fiber.Ctx) error {
	products, err := ph.repository.GetAllProducts()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "failed to find all products",
			"error":   err.Error(),
		})
	}

	return c.Status(200).JSON(products)
}

func (ph *ProductHandler) CreateProduct(c *fiber.Ctx) error {
	cloudName := os.Getenv("CLOUD_NAME")
	apiKey := os.Getenv("API_KEY")
	apiSecret := os.Getenv("API_SECRET")

	form, err := c.MultipartForm()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "internal server error",
		})
	}

	name := form.Value["name"][0]
	description := form.Value["description"][0]
	image := form.File["image"][0]
	price, _ := strconv.ParseFloat(form.Value["price"][0], 64)
	category := form.Value["category"][0]
	isOnOffer, _ := strconv.ParseBool(form.Value["is_on_offer"][0])

	cloud, err := cloudinary.NewFromParams(cloudName, apiKey, apiSecret)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "internal server error in cloud connection",
		})
	}

	uploadParams := uploader.UploadParams{
		Folder: "products",
	}

	uploadResult, err := cloud.Upload.Upload(context.Background(), image, uploadParams)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "internal server error in upload result",
		})
	}

	product := ph.factory.CreateProduct(name, description, uploadResult.SecureURL, category, price, isOnOffer)
	err = ph.repository.InsertProduct(product)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "failed to insert client",
		})
	}

	return c.Status(201).JSON(fiber.Map{
		"message": "success",
		"product": product,
	})
}

// func (ph *ProductHandler) uploadPhoto
