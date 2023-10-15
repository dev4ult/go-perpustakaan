package member

import "gorm.io/gorm"

type Member struct {
	gorm.Model

	ID int `gorm:"type:int(11)"`
	FullName string
	Email string
	Password string
}