package main

import (
	"perpustakaan/features"
	"perpustakaan/routes"

	"github.com/labstack/echo/v4"
)

var (
	bookHandler = features.BookHandler()
)

func main() {
	e := echo.New()

	routes.Books(e, bookHandler)
	// e.GET("/", func(c echo.Context) error {
	// 	return c.JSON(200, map[string]any {
	// 		"message": "Hello World",
	// 	})
	// })

	e.Logger.Fatal(e.Start(":8000"))
}