package author

import (
	"gorm.io/gorm"
)

type Author struct {
	gorm.Model
	
	ID int `gorm:"type:int(11)"`
	FullName string `gorm:"type:varchar(255)"`
	DOB string `gorm:"type:date"`
	Biography string `gorm:"type:text"`
}

type Authorship struct {
	gorm.Model
	
	ID int `gorm:"type:int(11)"`
	BookID int `gorm:"type:int(11)"`
	AuthorID int `gorm:"type:int(11)"`
}