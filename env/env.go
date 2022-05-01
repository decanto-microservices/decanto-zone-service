package env

import (
	"os"
	"sync"

	"github.com/joho/godotenv"
)

var lock = &sync.Mutex{}

type Config struct {
	Port string
}

func newConfig() *Config {
	godotenv.Load(".env")

	config := &Config{}

	// ----- SET Values -----
	config.Port = os.Getenv("PORT")

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
