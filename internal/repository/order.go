package repository

import (
	"context"
	"fmt"

	"github.com/Javex-Inc/panchito-backend/internal/model"
	"go.mongodb.org/mongo-driver/bson"
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

func (o *OrderRepository) GetAllOrders() ([]model.Order, error) {
	var orders []model.Order

	cur, err := o.collection.Find(context.Background(), bson.M{})
	if err != nil {
		return nil, fmt.Errorf("failed to find orders: %w", err)
	}

	for cur.Next(context.Background()) {
		var order model.Order

		err = cur.Decode(&order)
		if err != nil {
			return nil, fmt.Errorf("failed to decode order: %w", err)
		}

		orders = append(orders, order)
	}

	return orders, nil
}
