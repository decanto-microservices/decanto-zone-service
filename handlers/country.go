package handlers

import (
	"github.com/Gprisco/decanto-zone-service/db"
	"github.com/Gprisco/decanto-zone-service/models"
)

func GetCountries() []models.Country {
	var countries []models.Country

	db.GetDB().Find(&countries)

	return countries
}
