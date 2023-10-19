package routes

import (
	"perpustakaan/config"
	"perpustakaan/features/book"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func Books(e *echo.Echo, handler book.Handler, cfg config.ServerConfig) {
	books := e.Group("/books")

	books.GET("", handler.GetBooks())
	books.POST("", handler.CreateBook(), echojwt.JWT([]byte(cfg.SIGN_KEY)))
	
	books.GET("/:id", handler.BookDetails())
	books.PUT("/:id", handler.UpdateBook(), echojwt.JWT([]byte(cfg.SIGN_KEY)))
	books.DELETE("/:id", handler.DeleteBook(), echojwt.JWT([]byte(cfg.SIGN_KEY)))
}