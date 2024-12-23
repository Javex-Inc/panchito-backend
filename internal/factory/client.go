package factory

import (
	"github.com/Javex-Inc/panchito-backend/internal/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ClientFactory struct{}

func NewClientFactory() *ClientFactory {
	return &ClientFactory{}
}

func (*ClientFactory) CreateClient(name, phone string, address []model.Address) *model.Client {
	return &model.Client{
		ID:      primitive.NewObjectID(),
		Name:    name,
		Phone:   phone,
		Address: address,
	}
}
