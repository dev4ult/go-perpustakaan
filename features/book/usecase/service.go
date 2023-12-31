package usecase

import (
	"fmt"
	"mime/multipart"
	"perpustakaan/features/book"
	"perpustakaan/features/book/dtos"
	"perpustakaan/helpers"

	"github.com/mashingan/smapping"
)

type service struct {
	model book.Repository
	helper helpers.Helper
}

func New(model book.Repository, helper helpers.Helper) book.Usecase {
	return &service {
		model: model,
		helper: helper,
	}
}

func (svc *service) FindAll(page int, size int, searchKey string) ([]dtos.ResBook, string) {
	books, err := svc.model.Paginate(page, size, searchKey)

	if err != nil {
		return nil, err.Error() 
	}

	return books, ""
}

func (svc *service) FindByID(bookID int) (*dtos.ResBook, string) {
	book, err := svc.model.SelectByID(bookID)

	if err != nil {
		return nil, err.Error()
	}

	return book, ""
}

func (svc *service) Create(newBook dtos.InputBook, bookCover multipart.File) (*dtos.AfterInsert, string) {
	fmt.Println("In Service, formFile type : ", bookCover)
	var book book.Book
	
	if err := smapping.FillStruct(&book, smapping.MapFields(newBook)); err != nil {
		return nil, err.Error()
	}
	
	imageURL := svc.helper.UploadImage("book_cover", bookCover)

	book.CoverImage = imageURL
	bookID, err := svc.model.Insert(book)

	if err != nil {
		return nil, err.Error()
	}

	var resAfterInsert dtos.AfterInsert
	resAfterInsert.ID = bookID
	if err := smapping.FillStruct(&resAfterInsert, smapping.MapFields(newBook)); err != nil {
		return nil, err.Error()
	}
	
	resAfterInsert.CoverImage = imageURL

	return &resAfterInsert, ""
}

func (svc *service) Modify(bookData dtos.InputBook, bookID int, bookCover multipart.File) (bool, string) {
	var newBook book.Book
	
	if err := smapping.FillStruct(&newBook, smapping.MapFields(bookData)); err != nil {
		return false, err.Error()
	}

	imageURL := svc.helper.UploadImage("book_cover", bookCover)


	newBook.CoverImage = imageURL
	newBook.ID = bookID
	_, err := svc.model.Update(newBook)

	if err != nil {
		return false, err.Error()
	}
	
	return true, ""
}

func (svc *service) Remove(bookID int) (bool, string) {
	_, err := svc.model.DeleteByID(bookID)

	if err != nil {
		return false, err.Error()
	}

	return true, ""
}