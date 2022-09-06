package db

import (
	"context"
	"crypto/tls"
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

			tlsConfig := tls.Config{
				MinVersion: tls.VersionTLS12,
			}
			client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri).SetTLSConfig(&tlsConfig))

			if err != nil {
				panic(err)
			}

			singleton = client.Database(env.GetInstance().DB)
		}
	}

	return singleton
}
