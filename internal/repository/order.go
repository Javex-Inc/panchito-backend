package repository

import (
	"context"
	"fmt"

	"github.com/Javex-Inc/panchito-backend/internal/model"
	"go.mongodb.org/mongo-driver/mongo"
)

type OrderRepository struct {
	collection *mongo.Collection
}

func NewOrderRepository(client *mongo.Client, dbName, collection string) *OrderRepository {
	return &OrderRepository{
		collection: client.Database(dbName).Collection(collection),
	}
}

func (o *OrderRepository) CreateOrder(order *model.Order) error {
	_, err := o.collection.InsertOne(context.Background(), order)

	if err != nil {
		return fmt.Errorf("error trying to create order: %w", err)
	}

	return nil
}
