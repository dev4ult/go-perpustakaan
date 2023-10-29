package handler

import (
	"perpustakaan/helpers"
	"strconv"

	"perpustakaan/features/book"
	"perpustakaan/features/book/dtos"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

type controller struct {
	service book.Usecase
}

func New(service book.Usecase) book.Handler {
	return &controller {
		service: service,
	}
}

func (ctl *controller) GetBooks() echo.HandlerFunc {
	return func (ctx echo.Context) error  {
		pagination := dtos.Pagination{}
		ctx.Bind(&pagination)
		
		page := pagination.Page
		size := pagination.Size
		searchKey := ctx.QueryParam("title")

		if page <= 0 || size <= 0 {
			page = 1
			size = 10
		}

		books, message := ctl.service.FindAll(page, size, searchKey)

		if len(books) == 0 {
			return ctx.JSON(404, helpers.Response("There Is No Books!"))
		}

		if message != "" {
			return ctx.JSON(500, helpers.Response(message))
		}

		return ctx.JSON(200, helpers.Response("Success!", map[string]any {
			"data": books,
		}))
	}
}


func (ctl *controller) BookDetails() echo.HandlerFunc {
	return func (ctx echo.Context) error  {
		bookID, err := strconv.Atoi(ctx.Param("id"))

		if err != nil {
			return ctx.JSON(400, helpers.Response("Param must be provided in number!"))
		}

		book, message := ctl.service.FindByID(bookID)

		if book == nil {
			return ctx.JSON(404, helpers.Response(message))
		}

		return ctx.JSON(200, helpers.Response("Success!", map[string]any {
			"data": book,
		}))
	}
}

func (ctl *controller) CreateBook() echo.HandlerFunc {
	return func (ctx echo.Context) error  {
		input := ctx.Get("request").(*dtos.InputBook)

		formHeader, err := ctx.FormFile("cover-img")
		if err != nil {
			return ctx.JSON(400, helpers.Response("Missing Cover Image as `cover-img` (Required!)"))
		}

		formFile, err := formHeader.Open()
		if err != nil {
			return ctx.JSON(500, helpers.Response(err.Error()))
		}

		book, message := ctl.service.Create(*input, formFile)

		if book == nil {
			log.Error(message)
			return ctx.JSON(500, helpers.Response(message))
		}

		return ctx.JSON(201, helpers.Response("Success!", map[string]any {
			"data": book,
		}))
	}
}

func (ctl *controller) UpdateBook() echo.HandlerFunc {
	return func (ctx echo.Context) error {
		input := ctx.Get("request").(*dtos.InputBook)

		bookID, errParam := strconv.Atoi(ctx.Param("id"))

		if errParam != nil {
			return ctx.JSON(400, helpers.Response("Param must be provided in number!"))
		}

		book, message := ctl.service.FindByID(bookID)

		if book == nil {
			return ctx.JSON(404, helpers.Response(message))
		}
		

		update, errMessage := ctl.service.Modify(*input, bookID)

		if !update {
			return ctx.JSON(500, helpers.Response(errMessage))
		}

		return ctx.JSON(200, helpers.Response("Book Success Updated!"))
	}
}

func (ctl *controller) DeleteBook() echo.HandlerFunc {
	return func (ctx echo.Context) error  {
		bookID, err := strconv.Atoi(ctx.Param("id"))

		if err != nil {
			return ctx.JSON(400, helpers.Response("Param must be provided in number!"))
		}

		book, message := ctl.service.FindByID(bookID)

		if book == nil {
			return ctx.JSON(404, helpers.Response(message))
		}

		delete, errMessage := ctl.service.Remove(bookID)

		if !delete {
			return ctx.JSON(500, helpers.Response(errMessage))
		}

		return ctx.JSON(200, helpers.Response("Book Success Deleted!"))
	}
}
