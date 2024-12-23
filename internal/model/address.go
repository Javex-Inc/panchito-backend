package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Address struct {
	ID              primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	ZipCode         string             `json:"zip_code" bson:"zip_code"`
	Block           string             `json:"block" bson:"block"`
	ResidentialUnit string             `json:"residential_unit" bson:"residential_unit"`
	Number          string             `json:"number" bson:"number"`
	Sector          string             `json:"sector" bson:"sector"`
}
