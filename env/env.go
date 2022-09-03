package env

import (
	"fmt"
	"os"
	"sync"

	"github.com/Gprisco/decanto-zone-service/consul"
	"github.com/joho/godotenv"
)

var lock = &sync.Mutex{}

type Config struct {
	Port      string
	DSN       string
	DB        string
	BaseURL   string
	ServiceID string
}

func newConfig() *Config {
	godotenv.Load(".env")

	config := &Config{}

	consul := consul.GetInstance()

	mongoConnString, _, err := consul.KV().Get("mongodb/connection", nil)
	mongoDbPair, _, err := consul.KV().Get("mongodb/db", nil)

	if err != nil {
		panic(err)
	}

	// ----- SET Values -----
	config.Port = os.Getenv("PORT")
	config.DSN = fmt.Sprintf(string(mongoConnString.Value))
	config.DB = string(mongoDbPair.Value)

	config.BaseURL = os.Getenv("BASE_URL")
	config.ServiceID = os.Getenv("SERVICE_ID")

	return config
}

var singleton *Config

func GetInstance() *Config {
	if singleton == nil {
		lock.Lock()
		defer lock.Unlock()

		if singleton == nil {
			singleton = newConfig()
		}
	}

	return singleton
}
