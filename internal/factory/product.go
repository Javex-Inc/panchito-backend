package factory

import (
	"github.com/Javex-Inc/panchito-backend/internal/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ProductFactory struct{}

func NewProductFactory() *ProductFactory {
	return &ProductFactory{}
}

func (p *ProductFactory) CreateProduct(name, description, image string, category model.Category, price float64, isOnOffer bool) *model.Product {
	return &model.Product{
		ID:          primitive.NewObjectID(),
		Name:        name,
		Description: description,
		Image:       image,
		Price:       price,
		Category:    category,
		IsOnOffer:   isOnOffer,
	}
}
