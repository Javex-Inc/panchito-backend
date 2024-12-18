package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Payment struct {
	ID        primitive.ObjectID `json:"id"`
	Method    string             `json:"method"`
	Status    Status             `json:"status"`
	Discount  float32            `json:"discount"`
	Fee       float32            `json:"fee"`
	Total     float32            `json:"total"`
	Timestamp time.Time          `json:"timestamp"`
}
