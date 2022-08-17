package services

import (
	"context"

	"github.com/Gprisco/decanto-zone-service/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetRegions() []models.Region {
	var regions []models.Region

	lookup := bson.D{
		{Key: "$lookup", Value: bson.D{
			{Key: "from", Value: models.CountryCollection.Name()},
			{Key: "localField", Value: "countryId"},
			{Key: "foreignField", Value: "_id"},
			{Key: "as", Value: "country"},
		}},
	}

	unwind := bson.D{
		{Key: "$unwind", Value: bson.D{
			{Key: "path", Value: "$country"},
		}},
	}

	cursor, err := models.RegionCollection.Aggregate(context.TODO(), mongo.Pipeline{lookup, unwind}, &options.AggregateOptions{})

	if err != nil {
		panic(err)
	}

	cursor.All(context.TODO(), &regions)

	return regions
}

func GetRegion(id primitive.ObjectID) models.Region {
	var region models.Region
	err := models.RegionCollection.FindOne(context.TODO(), bson.D{{Key: "_id", Value: id}}).Decode(&region)

	if err != nil {
		panic(err)
	}

	return region
}
