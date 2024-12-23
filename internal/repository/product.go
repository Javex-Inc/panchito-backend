package repository

import (
	"context"
	"fmt"

	"github.com/Javex-Inc/panchito-backend/internal/model"
	"go.mongodb.org/mongo-driver/mongo"
)

type ProductRepository struct {
	collection *mongo.Collection
}

func NewProductRepository(client *mongo.Client, dbName string, collection string) *ProductRepository {
	return &ProductRepository{
		collection: client.Database(dbName).Collection(collection),
	}
}

func (p *ProductRepository) InsertProduct(product *model.Product) error {
	_, err := p.collection.InsertOne(context.Background(), product)

	if err != nil {
		fmt.Errorf("error trying to insert product: %w", err)
	}

	return nil
}
