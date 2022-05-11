package models

type Region struct {
	RegionId  int     `gorm:"column:regionId;primarykey" json:"regionId"`
	Region    string  `gorm:"column:region" json:"region"`
	CountryId int     `gorm:"column:countryId" json:"countryId"`
	Country   Country `gorm:"foreignKey:countryId;references:countryId"`
}

func (Region) TableName() string {
	return "g_regions"
}
