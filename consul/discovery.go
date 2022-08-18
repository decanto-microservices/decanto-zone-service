package consul

import (
	consulapi "github.com/hashicorp/consul/api"
)

func Discovery() map[string]*consulapi.AgentService {
	consul := GetInstance()

	services, error := consul.Agent().Services()

	if error != nil {
		panic(error)
	}

	return services
}
