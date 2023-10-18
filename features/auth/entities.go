package auth

import (
	"gorm.io/gorm"
)

type Librarian struct {
	gorm.Model

	ID int `gorm:"type:int(11)"`
	FullName string `gorm:"type:varchar(255)"`
	StaffID string 
	NIK int
	PhoneNumber string
	Address string
}

