package handlers

import (
	"net/http"
	"strconv"

	"github.com/Gprisco/decanto-zone-service/helpers"
	"github.com/Gprisco/decanto-zone-service/services"
	"github.com/gin-gonic/gin"
)

func GetCountries(c *gin.Context) {
	countries := services.GetCountries()
	c.JSON(http.StatusOK, countries)
}

func GetCountry(c *gin.Context) {
	countryId, err := strconv.Atoi(c.Param("countryId"))

	helpers.CheckForError(c, err)

	country := services.GetCountry(countryId)

	c.JSON(http.StatusOK, country)
}
