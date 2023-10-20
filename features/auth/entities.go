package auth

import (
	"gorm.io/gorm"
)

type Librarian struct {
	gorm.Model

	ID int `gorm:"type:int(11)"`
	FullName string `gorm:"type:varchar(255)"`
	StaffID string `gorm:"type:varchar(255)"`
	NIK int `gorm:"type:varchar(16)"`
	PhoneNumber string `gorm:"type:varchar(20)"`
	Address string `gorm:"text"`
}

