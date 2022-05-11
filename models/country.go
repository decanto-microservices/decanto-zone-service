package models

type Country struct {
	CountryId int    `gorm:"column:countryId;primarykey" json:"countryId"`
	Country   string `gorm:"column:country" json:"country"`
}

func (Country) TableName() string {
	return "g_countries"
}
