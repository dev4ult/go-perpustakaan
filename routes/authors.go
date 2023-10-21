package routes

import (
	"perpustakaan/features/author"

	"github.com/labstack/echo/v4"
)

func Authors(e *echo.Echo, handler author.Handler) {
	authors := e.Group("/authors")

	authors.GET("", handler.GetAuthors())
	authors.POST("", handler.CreateAuthor())
	
	authors.GET("/:id", handler.AuthorDetails())
	authors.PUT("/:id", handler.UpdateAuthor())
	authors.DELETE("/:id", handler.DeleteAuthor())

	authorships := e.Group("/authorships")
	authorships.POST("", handler.CreateAnAuthorship())
	authorships.DELETE("/:id", handler.DeleteAnAuthorship())
}