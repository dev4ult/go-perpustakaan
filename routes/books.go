package routes

import (
	"perpustakaan/features/book"
	"perpustakaan/features/book/dtos"
	m "perpustakaan/middlewares"

	"github.com/labstack/echo/v4"
)

func Books(e *echo.Echo, handler book.Handler) {
	books := e.Group("/books")

	books.GET("", handler.GetBooks())
	books.POST("", handler.CreateBook(), m.Authorization("librarian"), m.RequestValidation(&dtos.InputBook{}))
	
	books.GET("/:id", handler.BookDetails())
	books.PUT("/:id", handler.UpdateBook(), m.Authorization("librarian"), m.RequestValidation(&dtos.InputBook{}))
	books.DELETE("/:id", handler.DeleteBook(), m.Authorization("librarian"))
}