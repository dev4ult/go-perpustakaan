package ayam

import (
	"gorm.io/gorm"
)

type Ayam struct {
	gorm.Model

	ID int `gorm:"type:int(11)"`
	Name string `gorm:"type:varchar(255)"`
}

