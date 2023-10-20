package repository

import (
	"perpustakaan/features/book"
	"perpustakaan/features/book/dtos"

	"github.com/labstack/gommon/log"
	"gorm.io/gorm"
)

type model struct {
	db *gorm.DB
}

func New(db *gorm.DB) book.Repository {
	return &model {
		db: db,
	}
}

func (mdl *model) Paginate(page, size int) ([]dtos.ResBook, error){
	var books []dtos.ResBook

	offset := (page - 1) * size

	result := mdl.db.Table("books").Select("books.*, categories.name as category, publishers.name as publisher").Joins("LEFT JOIN categories ON categories.id = books.category_id").Joins("LEFT JOIN publishers ON publishers.id = books.publisher_id").Offset(offset).Limit(size).Scan(&books)

	if result.Error != nil {
		return nil, result.Error
	}

	return books, nil
}

func (mdl *model) Insert(newBook book.Book) int64 {
	result := mdl.db.Create(&newBook)

	if result.Error != nil {
		log.Error(result.Error)
		return -1
	}

	return int64(newBook.ID)
}

func (mdl *model) SelectByID(bookID int) *dtos.ResBook {
	var book dtos.ResBook
	result := mdl.db.Table("books").Where("books.id = ?", bookID).Select("books.*, categories.name as category, publishers.name as publisher").Joins("LEFT JOIN categories ON categories.id = books.category_id").Joins("LEFT JOIN publishers ON publishers.id = books.publisher_id").First(&book)

	if result.Error != nil {
		log.Error(result.Error)
		return nil
	}
	// fmt.Println(book)

	return &book
}

func (mdl *model) Update(book book.Book) int64 {
	result := mdl.db.Save(&book)

	if result.Error != nil {
		log.Error(result.Error)
	}

	return result.RowsAffected
}

func (mdl *model) DeleteByID(bookID int) int64 {
	result := mdl.db.Delete(&book.Book{}, bookID)
	
	if result.Error != nil {
		log.Error(result.Error)
		return 0
	}

	return result.RowsAffected
}