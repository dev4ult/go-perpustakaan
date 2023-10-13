package book

import (
	"github.com/labstack/echo/v4"
)

type Repository interface {
	Paginate(page, size int) []Book
	Insert(newBook Book) int64
	SelectByID(bookID int) Book
	Update(book Book) int64
	DeleteByID(bookID int) int64
}

type Usecase interface {
}

type Handler interface {
	GetBooks() echo.HandlerFunc
	BookDetails() echo.HandlerFunc
	CreateBook() echo.HandlerFunc
	UpdateBook() echo.HandlerFunc
	DeleteBook() echo.HandlerFunc
}
