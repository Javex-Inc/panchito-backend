package main

import (
	"fmt"
	"log"

	"github.com/Javex-Inc/panchito-backend/internal/utils"
)

func main() {
	err := utils.LoadEnv()
	if err != nil {
		log.Fatalf("failed to load .env variables: %w", err)
	}

	fmt.Println("Hello World!")
}
