package handlers

import (
	"net/http"
	"strconv"

	"github.com/Gprisco/decanto-zone-service/helpers"
	"github.com/Gprisco/decanto-zone-service/services"
	"github.com/gin-gonic/gin"
)

func GetRegions(c *gin.Context) {
	regions := services.GetRegions()
	c.JSON(http.StatusOK, regions)
}

func GetRegion(c *gin.Context) {
	regionId, err := strconv.Atoi(c.Param("regionId"))

	helpers.CheckForError(c, err)

	region := services.GetRegion(regionId)

	c.JSON(http.StatusOK, region)
}
