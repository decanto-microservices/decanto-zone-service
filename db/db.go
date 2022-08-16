package db

import (
	"context"
	"sync"

	"github.com/Gprisco/decanto-zone-service/env"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var lock = &sync.Mutex{}

var singleton *mongo.Database

func GetInstance() *mongo.Database {
	if singleton == nil {
		lock.Lock()
		defer lock.Unlock()

		if singleton == nil {
			uri := env.GetInstance().DSN

			client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))

			if err != nil {
				panic(err)
			}

			singleton = (client.Database(env.GetInstance().DB))
		}
	}

	return singleton
}
