package services

import (
	"github.com/Gprisco/decanto-zone-service/db"
	"github.com/Gprisco/decanto-zone-service/models"
)

func GetRegions() []models.Region {
	var regions []models.Region

	result := db.GetInstance().Joins("Country").Find(&regions)

	if result.Error != nil {
		return nil
	}

	return regions
}

func GetRegion(regionId int) *models.Region {
	var region = &models.Region{RegionId: regionId}

	result := db.GetInstance().Joins("Country").Find(&region)

	if result.Error != nil {
		return nil
	}

	return region
}
