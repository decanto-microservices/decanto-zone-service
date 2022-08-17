package env

import (
	"fmt"
	"os"
	"sync"

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

	// ----- SET Values -----
	config.Port = os.Getenv("PORT")
	config.DSN = fmt.Sprintf(
		"mongodb://%s:%s@%s:%s/?retryWrites=true&w=majority",
		os.Getenv("MONGO_USER"),
		os.Getenv("MONGO_PASS"),
		os.Getenv("MONGO_ADDR"),
		os.Getenv("MONGO_PORT"),
	)
	config.DB = os.Getenv("MONGO_DB")
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
