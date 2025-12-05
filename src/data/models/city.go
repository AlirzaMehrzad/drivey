package models

type City struct {
	BaseModel
	Name      string  `gorm:"size:10;type:string;not null"`
	Country   Country `gorm:"foreignKey:CountryId"`
	CountryId int
}
