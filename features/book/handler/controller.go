package handler

import (
	"perpustakaan/features/book"
	helper "perpustakaan/helpers"

	"github.com/labstack/echo/v4"
)

type controller struct {
	service book.Usecase
}

func (ctl *controller) GetBooks() echo.HandlerFunc {
	return func (ctx echo.Context) error  {
		return ctx.JSON(200, helper.Response("Success", nil))
	}
}


func (ctl *controller) BookDetails() echo.HandlerFunc {
	return func (ctx echo.Context) error  {
		return ctx.JSON(200, helper.Response("Success", nil))
	}
}

func (ctl *controller) CreateBook() echo.HandlerFunc {
	return func (ctx echo.Context) error  {
		return ctx.JSON(200, helper.Response("Success", nil))
	}
}

func (ctl *controller) UpdateBook() echo.HandlerFunc {
	return func (ctx echo.Context) error  {
		return ctx.JSON(200, helper.Response("Success", nil))
	}
}

func (ctl *controller) DeleteBook() echo.HandlerFunc {
	return func (ctx echo.Context) error  {
		return ctx.JSON(200, helper.Response("Success", nil))
	}
}
