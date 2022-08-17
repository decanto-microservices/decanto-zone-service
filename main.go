package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/Gprisco/decanto-zone-service/env"
	"github.com/Gprisco/decanto-zone-service/handlers"
	"github.com/gin-gonic/gin"

	consulapi "github.com/hashicorp/consul/api"
)

func main() {
	serviceRegistryWithConsul()

	r := gin.Default()
	baseUrl := env.GetInstance().BaseURL

	r.GET(baseUrl+"/check", (func(c *gin.Context) {
		c.JSON(http.StatusOK, nil)
	}))

	r.GET(baseUrl+"/country", handlers.GetCountries)
	r.GET(baseUrl+"/country/:countryId", handlers.GetCountry)

	r.GET(baseUrl+"/region", handlers.GetRegions)
	r.GET(baseUrl+"/region/:regionId", handlers.GetRegion)

	r.Run(env.GetInstance().Port)
}

func serviceRegistryWithConsul() {
	config := consulapi.DefaultConfig()
	consul, err := consulapi.NewClient(config)
	if err != nil {
		log.Println(err)
	}

	serviceID := "decanto-zone-service"
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
