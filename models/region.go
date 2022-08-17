package models

import (
	"github.com/Gprisco/decanto-zone-service/db"
	"github.com/Gprisco/decanto-zone-service/env"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Region struct {
	ID      primitive.ObjectID `bson:"_id" json:"_id"`
	Region  string             `bson:"region" json:"region"`
	Country *Country           `bson:"country,omitempty" json:"country,omitempty"`
}

var RegionCollection *mongo.Collection = db.GetInstance().Client().Database(env.GetInstance().DB).Collection("g_regions")
