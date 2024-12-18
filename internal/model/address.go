package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Address struct {
	ID              primitive.ObjectID `json:"id"`
	ZipCode         string             `json:"zipCode"`
	Block           string             `json:"block"`
	ResidentialUnit string             `json:"residential_unit"`
	Number          string             `json:"number"`
	Sector          string             `json:"sector"`
}
