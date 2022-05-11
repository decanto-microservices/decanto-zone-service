package env

import (
	"fmt"
	"os"
	"sync"

	"github.com/joho/godotenv"
)

var lock = &sync.Mutex{}

type Config struct {
	Port string
	DSN  string
}

func newConfig() *Config {
	godotenv.Load(".env")

	config := &Config{}

	// ----- SET Values -----
	config.Port = os.Getenv("PORT")
	config.DSN = fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_ADDR"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)

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
