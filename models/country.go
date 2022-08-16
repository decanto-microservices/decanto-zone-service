package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Country struct {
	ID      primitive.ObjectID `bson:"_id" json:"_id"`
	Country string             `json:"country"`
}

var CountryCollection string = "g_countries"
