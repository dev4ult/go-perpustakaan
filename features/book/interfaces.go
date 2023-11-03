package book

import (
	"mime/multipart"
	"perpustakaan/features/book/dtos"

	"github.com/labstack/echo/v4"
)

type Repository interface {
	Paginate(page int, size int, searchKey string) ([]dtos.ResBook, error)
	Insert(newBook Book) (int, error)
	SelectByID(bookID int) (*dtos.ResBook, error)
	Update(book Book) (int, error)
	DeleteByID(bookID int) (int, error)
}

type Usecase interface {
	FindAll(page int, size int, searchKey string) ([]dtos.ResBook, string)
	FindByID(bookID int) (*dtos.ResBook, string)
	Create(newBook dtos.InputBook, bookCover multipart.File) (*dtos.AfterInsert, string)
	Modify(bookData dtos.InputBook, bookID int, bookCover multipart.File) (bool, string)
	Remove(bookID int) (bool, string)
}

type Handler interface {
	GetBooks() echo.HandlerFunc
	BookDetails() echo.HandlerFunc
	CreateBook() echo.HandlerFunc
	UpdateBook() echo.HandlerFunc
	DeleteBook() echo.HandlerFunc
}
