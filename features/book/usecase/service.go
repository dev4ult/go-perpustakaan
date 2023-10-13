package usecase

import (
	"perpustakaan/features/book"
	"perpustakaan/features/book/dtos"

	"github.com/labstack/gommon/log"
	"github.com/mashingan/smapping"
)

type service struct {
	model book.Repository
}

func (svc *service) GetAll(page, size int) []dtos.ResBook {
	if page < 1 || size < 1 {
		return nil
	}
	
	return nil
}

func (svc *service) GetDetail(bookID int) *dtos.ResBook {
	res := dtos.ResBook{}
	book := svc.model.SelectByID(bookID)

	err := smapping.FillStruct(&res, smapping.MapFields(book))
	if err != nil {
		log.Error(err)
		return nil
	}

	return &res
}

func (svc *service) Create(newBook dtos.InputBook) *dtos.ResBook {
	book := book.Book{}
	
	err := smapping.FillStruct(&book, smapping.MapFields(newBook))
	if err != nil {
		log.Error(err)
		return nil
	}

	bookID := svc.model.Insert(book)

	if bookID == -1 {
		return nil
	}

	resBook := dtos.ResBook{}
	errRes := smapping.FillStruct(&resBook, smapping.MapFields(newBook))
	if errRes != nil {
		log.Error(errRes)
		return nil
	}

	return &resBook
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
		log.Error("There is No Books Updated!")
		return false
	}
	
	return true
}

func (svc *service) Remove(bookId int) bool {
	return true
}