package factory

import "github.com/Javex-Inc/panchito-backend/internal/model"

type ProductFactory struct{}

func NewProductFactory() *ProductFactory {
	return &ProductFactory{}
}

func (p *ProductFactory) CreateProduct(name, description, image, category string, price float32, isOnOffer bool) *model.Product {
	return &model.Product{
		Name:        name,
		Description: description,
		Image:       image,
		Price:       price,
		Category:    category,
		IsOnOffer:   isOnOffer,
	}
}
