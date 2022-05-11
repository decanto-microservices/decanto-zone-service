package services

import (
	"github.com/Gprisco/decanto-zone-service/db"
	"github.com/Gprisco/decanto-zone-service/models"
)

func GetCountries() []models.Country {
	var countries []models.Country

	result := db.GetInstance().Find(&countries)

	if result.Error != nil {
		return nil
	}

	return countries
}

func GetCountry(countryId int) *models.Country {
	var country = &models.Country{CountryId: countryId}

	result := db.GetInstance().First(&country)

	if result.Error != nil {
		return nil
	}

	return country
}
