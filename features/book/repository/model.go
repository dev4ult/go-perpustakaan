package repository

import (
	"perpustakaan/features/book"
	"perpustakaan/features/book/dtos"

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

func (mdl *model) Paginate(page int, size int, searchKey string) ([]dtos.ResBook, error){
	var books []dtos.ResBook

	offset := (page - 1) * size
	title := "%" + searchKey + "%"

	if err := mdl.db.Table("books").Select("books.*, categories.name as category, publishers.name as publisher").Joins("LEFT JOIN categories ON categories.id = books.category_id").Joins("LEFT JOIN publishers ON publishers.id = books.publisher_id").Where("books.title LIKE ?", title).Offset(offset).Limit(size).Scan(&books).Error; err != nil {
		return nil, err
	}

	return books, nil
}

func (mdl *model) Insert(newBook book.Book) (int, error) {
	if err := mdl.db.Create(&newBook).Error; err != nil {
		return 0, err 
	}

	return newBook.ID, nil
}

func (mdl *model) SelectByID(bookID int) (*dtos.ResBook, error) {
	var book dtos.ResBook
	

	if  err := mdl.db.Table("books").Where("books.id = ?", bookID).Select("books.*, categories.name as category, publishers.name as publisher").Joins("LEFT JOIN categories ON categories.id = books.category_id").Joins("LEFT JOIN publishers ON publishers.id = books.publisher_id").First(&book).Error; err != nil {
		return nil, err
	}

	return &book, nil
}

func (mdl *model) Update(book book.Book) (int, error) {
	result := mdl.db.Save(&book)

	if result.Error != nil {
		return 0, result.Error
	}

	return int(result.RowsAffected), nil
}

func (mdl *model) DeleteByID(bookID int) (int, error) {
	result := mdl.db.Delete(&book.Book{}, bookID)
	
	if result.Error != nil {
		return 0, result.Error
	}

	return int(result.RowsAffected), nil
}