package book

import (
	"perpustakaan/features/author"
	"perpustakaan/features/publisher"

	"gorm.io/gorm"
)

type Book struct {
	gorm.Model

	ID int `gorm:"type:int(11)"`
	Title string `gorm:"type:varchar(255)"`
	Summary string `gorm:"type:text"`
	PublicationYear int `gorm:"type:int(7)"`
	Quantity int `gorm:"type:int(11)"`
	Language string `gorm:"type:varchar(255)"`
	NumberOfPages int `gorm:"type:int(11)"`
	
	CategoryID int `gorm:"type:int(11)"`
	PublisherID int `gorm:"type:int(11)"`
	
	Category Category
	Publisher publisher.Publisher

	Authors []author.Authorship
}

type Category struct {
	gorm.Model

	ID int `gorm:"type:int(11)"`
	Name string `gorm:"type:varchar(255)"`
}

type FineType struct {
	gorm.Model

	ID int `gorm:"type:int(11)"`
	Name string `gorm:"type:varchar(255)"`
	FineCost int `gorm:"type:int(11)"`
}