package main

import (
	"github.com/Gprisco/decanto-zone-service/env"
	"github.com/Gprisco/decanto-zone-service/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/", handlers.GetCountries)
	r.GET("/:countryId", handlers.GetCountry)

	r.Run(env.GetInstance().Port)
}
