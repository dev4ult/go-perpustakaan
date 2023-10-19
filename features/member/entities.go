package member

import "gorm.io/gorm"

type Member struct {
	gorm.Model

	ID int `gorm:"type:int(11)"`
	FullName string `gorm:"type:varchar(255)"`
	CredentialNumber string `gorm:"type:varchar(50)"`
	Email string `gorm:"type:varchar(255)"`
	Password string `gorm:"type:varchar(255)"`
	PhoneNumber string `gorm:"type:varchar(20)"`
	Address string `gorm:"type:text"`
}