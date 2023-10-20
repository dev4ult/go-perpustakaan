package usecase

import (
	"mime/multipart"
	"perpustakaan/features/book"
	"perpustakaan/features/book/dtos"
	"perpustakaan/helpers"

	"github.com/labstack/gommon/log"
	"github.com/mashingan/smapping"
)

type service struct {
	model book.Repository
}

func New(model book.Repository) book.Usecase {
	return &service {
		model: model,
	}
}

func (svc *service) FindAll(page, size int) []dtos.ResBook {
	var books, err = svc.model.Paginate(page, size)

	if err != nil {
		log.Error(err.Error())
		return nil
	}

	return books
}

func (svc *service) FindByID(bookID int) *dtos.ResBook {
	book := svc.model.SelectByID(bookID)

	if book == nil {
		return nil
	}

	// err := smapping.FillStruct(&res, smapping.MapFields(book))
	// if err != nil {
	// 	log.Error(err)
	// 	return nil
	// }

	return book
}

func (svc *service) Create(newBook dtos.InputBook, bookCover multipart.File) *dtos.AfterInsert {
	book := book.Book{}
	
	err := smapping.FillStruct(&book, smapping.MapFields(newBook))
	if err != nil {
		log.Error(err.Error())
		return nil
	}

	imageURL, err := helpers.UploadImage("book-cover", bookCover)
	
	if err != nil {
		log.Error(err.Error())
		return nil
	}

	book.CoverImage = imageURL
	bookID := svc.model.Insert(book)

	if bookID == -1 {
		return nil
	}

	resAfterInsert := dtos.AfterInsert{}
	errRes := smapping.FillStruct(&resAfterInsert, smapping.MapFields(newBook))
	if errRes != nil {
		log.Error(errRes)
		return nil
	}
	resAfterInsert.CoverImage = imageURL

	return &resAfterInsert
}

func (svc *service) Modify(bookData dtos.InputBook, bookID int) bool {
	newBook := book.Book{}

	err := smapping.FillStruct(&newBook, smapping.MapFields(bookData))
	if err != nil {
		log.Error(err)
		return false
	}

	newBook.ID = bookID
	rowsAffected := svc.model.Update(newBook)

	if rowsAffected <= 0 {
		log.Error("There is No Book Updated!")
		return false
	}
	
	return true
}

func (svc *service) Remove(bookID int) bool {
	rowsAffected := svc.model.DeleteByID(bookID)

	if rowsAffected <= 0 {
		log.Error("There is No Book Deleted!")
		return false
	}

	return true
}