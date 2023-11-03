package routes

import (
	"perpustakaan/features/author"
	"perpustakaan/features/author/dtos"
	m "perpustakaan/middlewares"

	"github.com/labstack/echo/v4"
)

func Authors(e *echo.Echo, handler author.Handler) {
	authors := e.Group("/authors")
	authors.Use(m.Authorization("librarian"))

	authors.GET("", handler.GetAuthors())
	authors.POST("", handler.CreateAuthor(), m.RequestValidation(&dtos.InputAuthor{}))
	
	authors.GET("/:id", handler.AuthorDetails())
	authors.PUT("/:id", handler.UpdateAuthor(), m.RequestValidation(&dtos.InputAuthor{}))
	authors.DELETE("/:id", handler.DeleteAuthor())

	authorships := e.Group("/authorships")
	authorships.Use(m.Authorization("librarian"))
	authorships.POST("", handler.CreateAnAuthorship(), m.RequestValidation(&dtos.InputAuthorshipIDS{}))
	authorships.DELETE("/:id", handler.DeleteAnAuthorship())
}