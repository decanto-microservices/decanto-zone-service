package consul

import (
	"log"
	"sync"

	consulapi "github.com/hashicorp/consul/api"
)

var lock = &sync.Mutex{}

var singleton *consulapi.Client

func GetInstance() *consulapi.Client {
	if singleton == nil {
		lock.Lock()
		defer lock.Unlock()

		if singleton == nil {
			config := consulapi.DefaultConfig()
			consul, err := consulapi.NewClient(config)

			if err != nil {
				log.Println(err)
			}

			singleton = consul
		}
	}

	return singleton
}
