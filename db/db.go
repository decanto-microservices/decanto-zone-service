package db

import (
	"sync"

	"github.com/Gprisco/decanto-zone-service/env"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var lock = &sync.Mutex{}

var singleton *gorm.DB

func GetDB() *gorm.DB {
	if singleton == nil {
		lock.Lock()
		defer lock.Unlock()

		if singleton == nil {
			db, err := gorm.Open(mysql.Open(env.GetInstance().DSN))

			if err != nil {
				panic("Failed to connect to db")
			}

			singleton = db
		}
	}

	return singleton
}
