package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Order struct {
	ID         primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Number     int                `json:"number" bson:"number"`
	ClientID   primitive.ObjectID `json:"client_id" bson:"client_id"`
	Status     Status             `json:"status" bson:"status"`
	IsDelivery bool               `json:"is_delivery" bson:"is_delivery"`
	IsTakeaway bool               `json:"is_takeaway" bson:"is_takeaway"`
	Products   []Product          `json:"products" bson:"products"`
	Payment    Payment            `json:"payment" bson:"payment"`
	Timestamp  time.Time          `json:"timestamp" bson:"timestamp"`
}
