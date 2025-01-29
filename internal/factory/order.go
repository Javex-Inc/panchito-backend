package factory

import (
	"time"

	"github.com/Javex-Inc/panchito-backend/internal/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type OrderFactory struct{}

func NewOrderFactory() *OrderFactory {
	return &OrderFactory{}
}

func (*OrderFactory) CreateOrder(number int, clientID primitive.ObjectID, status model.Status, isDelivery, isTakeaway bool, products []model.Product, payment model.Payment, total float64) *model.Order {
	return &model.Order{
		ID:         primitive.NewObjectID(),
		Number:     number,
		ClientID:   clientID,
		Status:     status,
		IsDelivery: isDelivery,
		IsTakeaway: isTakeaway,
		Products:   products,
		Payment:    payment,
		Timestamp:  time.Time{},
		Total:      total,
	}
}
