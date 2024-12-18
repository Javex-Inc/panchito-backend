package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Client struct {
	ID      primitive.ObjectID `json:"id" bson:"_id"`
	Name    string             `json:"name" bson:"name"`
	Phone   string             `json:"phone" bson:"phone"`
	Address []Address          `json:"address" bson:"address"`
}
