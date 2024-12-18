package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Order struct {
	ID         primitive.ObjectID `json:"id"`
	Number     int                `json:"number"`
	Client     Client             `json:"client"`
	Status     Status             `json:"status"`
	IsDelivery bool               `json:"isDelivery"`
	IsTakeaway bool               `json:"isTakeaway"`
	Products   []Product          `json:"products"`
	Payment    Payment            `json:"payment"`
	Discount   float32            `json:"discount"`
	Total      float32            `json:"total"`
	Timestamp  time.Time          `json:"timestamp"`
}
