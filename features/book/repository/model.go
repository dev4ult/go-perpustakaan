package repository

import (
	"perpustakaan/features/book"

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

func (mdl *model) Paginate(page, size int) []book.Book {
	var books []book.Book

	offset := (page - 1) * size

	result := mdl.db.Offset(offset).Limit(size).Find(&books)
	
	if result.Error != nil {
		log.Error(result.Error)
		return nil
	}

	return books
}

func (mdl *model) Insert(newBook book.Book) int64 {
	result := mdl.db.Create(&newBook)

	if result.Error != nil {
		log.Error(result.Error)
		return -1
	}

	return int64(newBook.ID)
}

func (mdl *model) SelectByID(bookID int) *book.Book {
	var book book.Book
	result := mdl.db.First(&book, bookID)

	if result.Error != nil {
		log.Error(result.Error)
		return nil
	}

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