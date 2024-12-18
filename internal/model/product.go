package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Product struct {
	ID          primitive.ObjectID `json:"id"`
	Name        string             `json:"name"`
	Description string             `json:"description"`
	Image       string             `json:"image"`
	Price       float32            `json:"price"`
	Category    string             `json:"category"`    // TODO: Will be Category type
	IsOnOffer   bool               `json:"is_on_offer"` // When the product is on offer
}
