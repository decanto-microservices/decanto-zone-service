package handlers

import (
	"net/http"
	"strings"

	"github.com/Gprisco/decanto-zone-service/services"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetRegions(c *gin.Context) {
	ids := c.Query("ids")
	objectIds := []primitive.ObjectID{}

	if ids != "" {
		for _, id := range strings.Split(ids, ",") {
			objectId, err := primitive.ObjectIDFromHex(string(id))

			if err != nil {
				c.JSON(http.StatusBadRequest, "The ids must be valid ObjectIDs")
				return
			}

			objectIds = append(objectIds, objectId)
		}
	}

	c.JSON(http.StatusOK, services.GetRegions(objectIds))
}

func GetRegion(c *gin.Context) {
	regionId, err := primitive.ObjectIDFromHex(c.Param("regionId"))

	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, services.GetRegion(regionId))
}
