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

	mongoUserPair, _, err := consul.KV().Get("mongodb/user/root", nil)
	mongoPassPair, _, err := consul.KV().Get("mongodb/user/root/password", nil)
	mongoAddressPair, _, err := consul.KV().Get("mongodb/address", nil)
	mongoPortPair, _, err := consul.KV().Get("mongodb/port", nil)
	mongoDbPair, _, err := consul.KV().Get("mongodb/db", nil)

	if err != nil {
		panic(err)
	}

	// ----- SET Values -----
	config.Port = os.Getenv("PORT")
	config.DSN = fmt.Sprintf(
		"mongodb://%s:%s@%s:%s/?retryWrites=true&w=majority",
		string(mongoUserPair.Value),
		string(mongoPassPair.Value),
		string(mongoAddressPair.Value),
		string(mongoPortPair.Value),
	)
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
