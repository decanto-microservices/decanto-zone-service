package models

type Country struct {
	CountryId int    `gorm:"column:countryId;primarykey"`
	Country   string `gorm:"column:country"`
}

func (Country) TableName() string {
	return "g_countries"
}
