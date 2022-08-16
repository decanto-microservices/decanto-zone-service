package services

import (
	"context"

	"github.com/Gprisco/decanto-zone-service/db"
	"github.com/Gprisco/decanto-zone-service/env"
	"github.com/Gprisco/decanto-zone-service/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var countryColl *mongo.Collection = db.GetInstance().Client().Database(env.GetInstance().DB).Collection(models.CountryCollection)

func GetCountries() []models.Country {
	cursor, err := countryColl.Find(context.TODO(), bson.D{})
	var countries []models.Country

	if err != nil {
		panic(err)
	}

	cursor.All(context.TODO(), &countries)

	return countries
}

func GetCountry(id primitive.ObjectID) models.Country {
	var country *models.Country
	err := countryColl.FindOne(context.TODO(), bson.D{{Key: "_id", Value: id}}).Decode(&country)

	if err != nil {
		panic(err)
	}

	return *country
}
