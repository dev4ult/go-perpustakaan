package publisher

import "gorm.io/gorm"

type Publisher struct {
	gorm.Model

	ID int `gorm:"type:int(11)"`
	Country string `gorm:"type:varchar(150)"`
	City string `gorm:"type:varchar(150)"`
	Address string `gorm:"type:varchar(255)"`
	PostalCode int `gorm:"type:int(11)"`
	PhoneNumber string `gorm:"type:varchar(20)"`
	Email string `gorm:"type:varchar(255)"`
}