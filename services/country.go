package services

import (
	"context"

	"github.com/Gprisco/decanto-zone-service/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetCountries() []models.Country {
	var countries []models.Country
	cursor, err := models.CountryCollection.Find(context.TODO(), bson.D{})

	if err != nil {
		panic(err)
	}

	cursor.All(context.TODO(), &countries)

	return countries
}

func GetCountry(id primitive.ObjectID) models.Country {
	var country models.Country
	err := models.CountryCollection.FindOne(context.TODO(), bson.D{{Key: "_id", Value: id}}).Decode(&country)

	if err != nil {
		panic(err)
	}

	return country
}
