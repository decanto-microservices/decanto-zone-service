package main

import (
	"github.com/Gprisco/decanto-zone-service/env"
	"github.com/Gprisco/decanto-zone-service/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, handlers.GetCountries())
	})

	r.Run(env.GetInstance().Port)
}
