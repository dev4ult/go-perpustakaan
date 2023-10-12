package book

import "github.com/labstack/echo/v4"

type Repository interface {
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
