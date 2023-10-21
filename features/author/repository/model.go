package repository

import (
	"perpustakaan/features/author"
	"perpustakaan/features/author/dtos"
	"perpustakaan/features/book"

	"github.com/labstack/gommon/log"
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

func (mdl *model) Paginate(page, size int) []author.Author {
	var authors []author.Author

	offset := (page - 1) * size

	result := mdl.db.Offset(offset).Limit(size).Find(&authors)
	
	if result.Error != nil {
		log.Error(result.Error)
		return nil
	}

	return authors
}

func (mdl *model) Insert(newAuthor author.Author) int64 {
	result := mdl.db.Create(&newAuthor)

	if result.Error != nil {
		log.Error(result.Error)
		return -1
	}

	return int64(newAuthor.ID)
}

func (mdl *model) SelectByID(authorID int) *author.Author {
	var author author.Author
	result := mdl.db.First(&author, authorID)

	if result.Error != nil {
		log.Error(result.Error)
		return nil
	}

	return &author
}

func (mdl *model) Update(author author.Author) int64 {
	result := mdl.db.Save(&author)

	if result.Error != nil {
		log.Error(result.Error)
	}

	return result.RowsAffected
}

func (mdl *model) DeleteByID(authorID int) int64 {
	result := mdl.db.Delete(&author.Author{}, authorID)
	
	if result.Error != nil {
		log.Error(result.Error)
		return 0
	}

	return result.RowsAffected
}

func (mdl *model) SelectAuthorshipByID(authorshipID int) *author.Authorship {
	var authorship author.Authorship
	
	if result := mdl.db.Table("authorships").Where("id = ?", authorshipID).First(&authorship); result.Error != nil {
		log.Error(result.Error.Error())
		return nil
	}

	return &authorship
}

func (mdl *model) IsAuthorshipExist(bookID, authorID int) bool {
	var authorship author.Authorship
	result := mdl.db.Where("book_id = ? AND author_id = ?", bookID, authorID).First(&authorship)

	return result.Error == nil
}

func (mdl *model) InsertAuthorship(authorship dtos.InputAuthorshipIDS) (*dtos.BookAuthors, error) {
	var bookAuthors dtos.BookAuthors

	result := mdl.db.Table("authorships").Create(&authorship)

	if result.Error != nil {
		return nil, result.Error
	}

	var book book.Book
	mdl.db.Table("books").Where("id = ?", authorship.BookID).First(&book)
	
	var authors []dtos.ResAuthor
	mdl.db.Table("authorships").Select("authors.*").Where("book_id = ?", authorship.BookID).Joins("LEFT JOIN authors ON authors.id = authorships.author_id").Find(&authors)
	
	if err := smapping.FillStruct(&bookAuthors, smapping.MapFields(book)); err != nil {
		return nil, err
	}

	for _, author := range authors {
		bookAuthors.Authors = append(bookAuthors.Authors, author)
	}

	return &bookAuthors, nil
}

func (mdl *model) DeleteAuthorshipByID(authorshipID int) int64 {
	result := mdl.db.Delete(&author.Authorship{}, authorshipID)
	
	if result.Error != nil {
		log.Error(result.Error)
		return 0
	}

	return result.RowsAffected
}