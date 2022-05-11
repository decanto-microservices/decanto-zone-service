package handlers

import (
	"net/http"
	"strconv"

	"github.com/Gprisco/decanto-zone-service/db"
	"github.com/Gprisco/decanto-zone-service/helpers"
	"github.com/Gprisco/decanto-zone-service/models"
	"github.com/gin-gonic/gin"
)

func GetCountries(c *gin.Context) {
	var countries []models.Country

	result := db.GetInstance().Find(&countries)

	if result.Error != nil {
		helpers.HandleGenericError(c, result.Error)
		return
	}

	c.JSON(http.StatusOK, countries)
}

func GetCountry(c *gin.Context) {
	countryId, err := strconv.Atoi(c.Param("countryId"))

	if err != nil {
		helpers.HandleGenericError(c, err)
		return
	}

	var country = &models.Country{CountryId: countryId}

	result := db.GetInstance().First(&country)

	if result.Error != nil {
		helpers.HandleGenericError(c, result.Error)
		return
	}

	c.JSON(http.StatusOK, country)
}
