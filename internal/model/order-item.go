package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type OrderItem struct {
	ID        primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	ProductID string             `json:"product_id" bson:"product_id"`
	Name      string             `json:"name" bson:"name"`
	Quantity  int                `json:"quantity" bson:"quantity"`
	Price     float64            `json:"price" bson:"price"`
	Subtotal  float64            `json:"subtotal" bson:"subtotal"`
}
