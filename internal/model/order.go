package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Order struct {
	ID         primitive.ObjectID `json:"id" bson:"_id"`
	Number     int                `json:"number"`
	ClientID   primitive.ObjectID `json:"client_id"`
	Status     Status             `json:"status"`
	IsDelivery bool               `json:"is_delivery"`
	IsTakeaway bool               `json:"is_takeaway"`
	Products   []Product          `json:"products"`
	Payment    Payment            `json:"payment"`
	Timestamp  time.Time          `json:"timestamp"`
}
