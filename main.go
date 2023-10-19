package main

import (
	"perpustakaan/features"
	"perpustakaan/routes"

	"github.com/labstack/echo/v4"
)

var (
	bookHandler = features.BookHandler()
	memberHandler = features.MemberHandler()
)

func main() {
	e := echo.New()

	routes.Books(e, bookHandler)
	routes.Members(e, memberHandler)
	
	e.Logger.Fatal(e.Start(":8000"))
}