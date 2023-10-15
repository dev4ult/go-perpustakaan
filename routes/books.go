package routes

import (
	"perpustakaan/features/book"

	"github.com/labstack/echo/v4"
)

func Books(e *echo.Echo, handler book.Handler) {
	books := e.Group("/books")

	books.GET("", handler.GetBooks())
	books.GET("/:id", handler.BookDetails())
	books.POST("", handler.CreateBook())
}