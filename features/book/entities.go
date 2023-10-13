package book

import "gorm.io/gorm"

type Book struct {
	gorm.Model
	ID int `gorm:"type:int(11)"`
	Title string `gorm:"type:varchar(255)"`
}