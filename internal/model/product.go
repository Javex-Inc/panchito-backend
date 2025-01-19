package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Product struct {
	ID          primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name        string             `json:"name" bson:"name"`
	Description string             `json:"description" bson:"description"`
	Image       string             `json:"image" bson:"image"`
	Price       float64            `json:"price" bson:"price"`
	Category    Category           `json:"category" bson:"category"`
	IsOnOffer   bool               `json:"is_on_offer" bson:"is_on_offer"` // When the product is on offer
}
