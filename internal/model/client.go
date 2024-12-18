package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Client struct {
	ID      primitive.ObjectID `json:"id"`
	Name    string             `json:"name"`
	Phone   string             `json:"phone"`
	Address []Address          `json:"address"`
}
