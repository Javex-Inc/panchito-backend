package database

import (
	"context"
	"fmt"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client

func Connect() error {
	URI := os.Getenv("DB_URI")

	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	options := options.Client().ApplyURI(URI).SetServerAPIOptions(serverAPI)

	client, err := mongo.Connect(context.Background(), options)
	if err != nil {
		return fmt.Errorf("failed to connect to mongodb: %w", err)
	}

	Client = client

	fmt.Println("DB connection successful!")

	return nil
}
