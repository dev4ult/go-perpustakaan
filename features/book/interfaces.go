package book

import (
	"mime/multipart"
	"perpustakaan/features/book/dtos"

	"github.com/labstack/echo/v4"
)

type Repository interface {
	Paginate(page, size int) ([]dtos.ResBook, error)
	Insert(newBook Book) int64
	SelectByID(bookID int) *dtos.ResBook
	Update(book Book) int64
	DeleteByID(bookID int) int64
}

type Usecase interface {
	FindAll(page, size int) []dtos.ResBook
	FindByID(bookID int) *dtos.ResBook
	Create(newBook dtos.InputBook, bookCover multipart.File) *dtos.AfterInsert
	Modify(bookData dtos.InputBook, bookID int) bool
	Remove(bookID int) bool
}

type Handler interface {
	GetBooks() echo.HandlerFunc
	BookDetails() echo.HandlerFunc
	CreateBook() echo.HandlerFunc
	UpdateBook() echo.HandlerFunc
	DeleteBook() echo.HandlerFunc
}
