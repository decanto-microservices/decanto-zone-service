package handlers

import (
	"net/http"

	"github.com/Gprisco/decanto-zone-service/services"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetRegions(c *gin.Context) {
	c.JSON(http.StatusOK, services.GetRegions())
}

func GetRegion(c *gin.Context) {
	regionId, err := primitive.ObjectIDFromHex(c.Param("regionId"))

	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, services.GetRegion(regionId))
}
