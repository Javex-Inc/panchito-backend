package main

import (
	"log"

	"github.com/Javex-Inc/panchito-backend/internal/database"
	"github.com/Javex-Inc/panchito-backend/internal/service"
	"github.com/Javex-Inc/panchito-backend/internal/utils"
)

func main() {
	err := utils.LoadEnv()
	if err != nil {
		log.Fatalf("failed to load .env variables: %v", err)
	}

	err = database.Connect()
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	service.StartServer(database.Client)
}
