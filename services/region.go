package services

import (
	"context"

	"github.com/Gprisco/decanto-zone-service/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetRegions(ids []primitive.ObjectID) []models.Region {
	var regions []models.Region

	aggregationPipeline := mongo.Pipeline{}

	if len(ids) > 0 {
		// If ids are specified, we append a match step in the aggregation pipeline
		aggregationPipeline = append(aggregationPipeline, bson.D{
			{Key: "$match", Value: bson.D{
				{Key: "_id", Value: bson.D{{Key: "$in", Value: ids}}},
			},
			}})
	}

	// After matching our ids, we add the lookup step
	aggregationPipeline = append(aggregationPipeline, bson.D{{Key: "$lookup", Value: bson.D{
		{Key: "from", Value: models.CountryCollection.Name()},
		{Key: "localField", Value: "countryId"},
		{Key: "foreignField", Value: "_id"},
		{Key: "as", Value: "country"},
	}}})

	// Lastly we add an unwind stage
	aggregationPipeline = append(aggregationPipeline, bson.D{
		{Key: "$unwind", Value: bson.D{
			{Key: "path", Value: "$country"},
		}},
	})

	cursor, err := models.RegionCollection.Aggregate(context.TODO(), aggregationPipeline, &options.AggregateOptions{})

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
