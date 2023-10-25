package transaction

import (
	"gorm.io/gorm"
)

type Transaction struct {
	gorm.Model

	ID int `gorm:"type:int(11)"`
	Name string `gorm:"type:varchar(255)"`
}

