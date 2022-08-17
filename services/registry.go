package services

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/Gprisco/decanto-zone-service/env"
	consulapi "github.com/hashicorp/consul/api"
)

func Register() {
	config := consulapi.DefaultConfig()
	consul, err := consulapi.NewClient(config)
	if err != nil {
		log.Println(err)
	}

	serviceID := env.GetInstance().ServiceID
	port, _ := strconv.Atoi(env.GetInstance().Port[1:len(env.GetInstance().Port)])
	address, _ := os.Hostname()

	registration := &consulapi.AgentServiceRegistration{
		ID:      serviceID,
		Name:    serviceID,
		Port:    port,
		Address: address,
		Check: &consulapi.AgentServiceCheck{
			HTTP:     fmt.Sprintf("http://%s:%v/%s/check", address, port, env.GetInstance().BaseURL),
			Interval: "10s",
			Timeout:  "30s",
		},
	}

	regiErr := consul.Agent().ServiceRegister(registration)

	if regiErr != nil {
		log.Printf("Failed to register service: %s:%v", address, port)
	} else {
		log.Printf("successfully register service: %s:%v", address, port)
	}
}
