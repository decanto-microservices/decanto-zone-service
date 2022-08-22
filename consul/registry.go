package consul

import (
	"fmt"
	"log"
	"os"
	"strconv"

	consulapi "github.com/hashicorp/consul/api"
	"github.com/joho/godotenv"
)

func Register() {
	godotenv.Load(".env")
	consul := GetInstance()

	serviceID := os.Getenv("SERVICE_ID")
	port, _ := strconv.ParseInt(os.Getenv("PORT")[1:len(os.Getenv("PORT"))], 10, 64)
	address, _ := os.Hostname()

	registration := &consulapi.AgentServiceRegistration{
		ID:      serviceID,
		Name:    serviceID,
		Port:    int(port),
		Address: address,
		Check: &consulapi.AgentServiceCheck{
			HTTP:     fmt.Sprintf("http://%s:%v/%s/check", address, port, os.Getenv("BASE_URL")),
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
