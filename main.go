package main

import (
	"github.com/Gprisco/decanto-zone-service/env"
	"github.com/Gprisco/decanto-zone-service/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/country", handlers.GetCountries)
	r.GET("/country/:countryId", handlers.GetCountry)

	r.GET("/region", handlers.GetRegions)
	r.GET("/region/:regionId", handlers.GetRegion)

	r.Run(env.GetInstance().Port)
}
