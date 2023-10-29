package repository

import (
	"perpustakaan/features/author"
	"perpustakaan/features/author/dtos"
	"perpustakaan/features/book"

	"github.com/mashingan/smapping"
	"gorm.io/gorm"
)

type model struct {
	db *gorm.DB
}

func New(db *gorm.DB) author.Repository {
	return &model {
		db: db,
	}
}

func (mdl *model) Paginate(page int, size int, searchKey string) ([]author.Author, error) {
	var authors []author.Author

	offset := (page - 1) * size
	name := "%" + searchKey + "%"
	
	if err := mdl.db.Where("full_name LIKE ?", name).Offset(offset).Limit(size).Find(&authors).Error; err != nil {
		return nil, err
	}

	return authors, nil
}

func (mdl *model) Insert(newAuthor author.Author) (int, error) {
	if err := mdl.db.Create(&newAuthor).Error; err != nil {
		return 0, err
	}

	return newAuthor.ID, nil
}

func (mdl *model) SelectByID(authorID int) (*author.Author, error) {
	var author author.Author

	if err := mdl.db.First(&author, authorID).Error; err != nil {
		return nil, err
	}

	return &author, nil
}

func (mdl *model) Update(author author.Author) (int, error) {
	result := mdl.db.Save(&author)

	if result.Error != nil {
		return 0, result.Error
	}

	return int(result.RowsAffected), nil
}

func (mdl *model) DeleteByID(authorID int) (int, error) {
	result := mdl.db.Delete(&author.Author{}, authorID)
	
	if result.Error != nil {
		return 0, result.Error
	}

	return int(result.RowsAffected), nil
}

func (mdl *model) SelectAuthorshipByID(authorshipID int) (*author.Authorship, error) {
	var authorship author.Authorship
	
	if err := mdl.db.Table("authorships").Where("id = ?", authorshipID).First(&authorship).Error; err != nil {
		return nil, err
	}

	return &authorship, nil
}

func (mdl *model) IsAuthorshipExist(bookID, authorID int) (bool, error) {
	var authorship author.Authorship
	if err := mdl.db.Where("book_id = ? AND author_id = ?", bookID, authorID).First(&authorship).Error; err != nil {
		return false, err
	}

	return true, nil
}

func (mdl *model) InsertAuthorship(authorship dtos.InputAuthorshipIDS) (*dtos.BookAuthors, error) {
	var bookAuthors dtos.BookAuthors

	if err := mdl.db.Table("authorships").Create(&authorship).Error; err != nil {
		return nil, err
	}

	var book book.Book
	if err := mdl.db.Table("books").Where("id = ?", authorship.BookID).First(&book).Error; err != nil {
		return nil, err
	}
	
	var authors []dtos.ResAuthor
	if err := mdl.db.Table("authorships").Select("authors.*").Where("book_id = ?", authorship.BookID).Joins("LEFT JOIN authors ON authors.id = authorships.author_id").Find(&authors).Error; err != nil {
		return nil, err
	}
	
	if err := smapping.FillStruct(&bookAuthors, smapping.MapFields(book)); err != nil {
		return nil, err
	}

	for _, author := range authors {
		bookAuthors.Authors = append(bookAuthors.Authors, author)
	}

	return &bookAuthors, nil
}

func (mdl *model) DeleteAuthorshipByID(authorshipID int) (int, error) {
	result := mdl.db.Delete(&author.Authorship{}, authorshipID)
	
	if result.Error != nil {
		return 0, result.Error
	}

	return int(result.RowsAffected), nil
}