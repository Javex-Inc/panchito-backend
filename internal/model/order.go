package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Order struct {
	ID         primitive.ObjectID `json:"id"`
	Number     int                `json:"number"`
	ClientID   primitive.ObjectID `json:"clientId"`
	Status     Status             `json:"status"`
	IsDelivery bool               `json:"isDelivery"`
	IsTakeaway bool               `json:"isTakeaway"`
	Products   []Product          `json:"products"`
	Payment    Payment            `json:"payment"`
	Timestamp  time.Time          `json:"timestamp"`
}
