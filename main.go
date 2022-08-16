package main

import (
	"github.com/Gprisco/decanto-zone-service/env"
	"github.com/Gprisco/decanto-zone-service/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	baseUrl := env.GetInstance().BaseURL

	r.GET(baseUrl+"/country", handlers.GetCountries)
	r.GET(baseUrl+"/country/:countryId", handlers.GetCountry)

	r.GET(baseUrl+"/region", handlers.GetRegions)
	r.GET(baseUrl+"/region/:regionId", handlers.GetRegion)

	r.Run(env.GetInstance().Port)
}
