package repository

import (
	"context"
	"fmt"

	"github.com/Javex-Inc/panchito-backend/internal/model"
	"go.mongodb.org/mongo-driver/bson"
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
		return fmt.Errorf("error trying to insert product: %w", err)
	}

	return nil
}

func (p *ProductRepository) GetAllProducts() ([]model.Product, error) {
	var products []model.Product

	cur, err := p.collection.Find(context.Background(), bson.M{})
	if err != nil {
		return nil, fmt.Errorf("failed to find all products: %w", err)
	}

	for cur.Next(context.Background()) {
		var product model.Product

		err := cur.Decode(&product)
		if err != nil {
			return nil, fmt.Errorf("failed to decode product: %w", err)
		}

		products = append(products, product)
	}

	return products, nil
}
