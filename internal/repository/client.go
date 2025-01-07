package repository

import (
	"context"
	"fmt"

	"github.com/Javex-Inc/panchito-backend/internal/model"
	"go.mongodb.org/mongo-driver/mongo"
)

type ClientRepository struct {
	collection *mongo.Collection
}

func NewClientRepository(client *mongo.Client, dbName string, collection string) *ClientRepository {
	return &ClientRepository{
		collection: client.Database(dbName).Collection(collection),
	}
}

func (c *ClientRepository) InsertClient(client *model.Client) error {
	_, err := c.collection.InsertOne(context.Background(), client)
	if err != nil {
		return fmt.Errorf("error trying to insert user into the collection: %w", err)
	}

	return nil
}
