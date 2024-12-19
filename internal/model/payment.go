package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Payment struct {
	ID        primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Method    string             `json:"method" bson:"method"`
	IsPaid    bool               `json:"isPaid" bson:"isPaid"`
	Discount  float32            `json:"discount" bson:"discount"`
	Fee       float32            `json:"fee" bson:"fee"`
	Total     float32            `json:"total" bson:"total"`
	Timestamp time.Time          `json:"timestamp" bson:"timestamp"`
}
