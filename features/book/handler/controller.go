package handler

import (
	"perpustakaan/helpers"
	helper "perpustakaan/helpers"
	"strconv"

	"perpustakaan/features/book"
	"perpustakaan/features/book/dtos"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type controller struct {
	service book.Usecase
}

func New(service book.Usecase) book.Handler {
	return &controller {
		service: service,
	}
}

var validate *validator.Validate

func (ctl *controller) GetBooks() echo.HandlerFunc {
	return func (ctx echo.Context) error  {
		pagination := dtos.Pagination{}
		ctx.Bind(&pagination)
		
		page := pagination.Page
		size := pagination.Size

		if page <= 0 || size <= 0 {
			page = 1
			size = 10
		}

		books := ctl.service.FindAll(page, size)

		if books == nil {
			return ctx.JSON(404, helper.Response("There is No Books!"))
		}

		return ctx.JSON(200, helper.Response("Success!", map[string]any {
			"data": books,
		}))
	}
}


func (ctl *controller) BookDetails() echo.HandlerFunc {
	return func (ctx echo.Context) error  {
		bookID, err := strconv.Atoi(ctx.Param("id"))

		if err != nil {
			return ctx.JSON(400, helper.Response("Param must be provided in number!"))
		}

		book := ctl.service.FindByID(bookID)

		if book == nil {
			return ctx.JSON(404, helper.Response("Book Not Found!"))
		}

		return ctx.JSON(200, helper.Response("Success!", map[string]any {
			"data": book,
		}))
	}
}

func (ctl *controller) CreateBook() echo.HandlerFunc {
	return func (ctx echo.Context) error  {
		input := dtos.InputBook{}

		formHeader, err := ctx.FormFile("cover-img")
		if err != nil {
			return ctx.JSON(400, helper.Response("Missing Cover Image as `cover-img` (Required!)"))
		}

		formFile, err := formHeader.Open()
		if err != nil {
			return ctx.JSON(500, helper.Response(err.Error()))
		}

		ctx.Bind(&input)
		validate = validator.New(validator.WithRequiredStructEnabled())
		errValidation := validate.Struct(input)

		if errValidation != nil {
			errMap := helpers.ErrorMapValidation(err)
			return ctx.JSON(400, helper.Response("Missing Data Required!", map[string]any {
				"errors": errMap,
			}))
		}

		book := ctl.service.Create(input, formFile)

		if book == nil {
			return ctx.JSON(500, helper.Response("Something Went Wrong!", nil))
		}

		return ctx.JSON(201, helper.Response("Success!", map[string]any {
			"data": book,
		}))
	}
}

func (ctl *controller) UpdateBook() echo.HandlerFunc {
	return func (ctx echo.Context) error {
		input := dtos.InputBook{}

		bookID, errParam := strconv.Atoi(ctx.Param("id"))

		if errParam != nil {
			return ctx.JSON(400, helper.Response("Param must be provided in number!"))
		}

		book := ctl.service.FindByID(bookID)

		if book == nil {
			return ctx.JSON(404, helper.Response("Book Not Found!"))
		}
		
		ctx.Bind(&input)

		validate = validator.New(validator.WithRequiredStructEnabled())
		err := validate.Struct(input)

		if err != nil {
			errMap := helpers.ErrorMapValidation(err)
			return ctx.JSON(400, helper.Response("Missing Data Required!", map[string]any {
				"error": errMap,
			}))
		}

		update := ctl.service.Modify(input, bookID)

		if !update {
			return ctx.JSON(500, helper.Response("Something Went Wrong!"))
		}

		return ctx.JSON(200, helper.Response("Book Success Updated!"))
	}
}

func (ctl *controller) DeleteBook() echo.HandlerFunc {
	return func (ctx echo.Context) error  {
		bookID, err := strconv.Atoi(ctx.Param("id"))

		if err != nil {
			return ctx.JSON(400, helper.Response("Param must be provided in number!"))
		}

		book := ctl.service.FindByID(bookID)

		if book == nil {
			return ctx.JSON(404, helper.Response("Book Not Found!"))
		}

		delete := ctl.service.Remove(bookID)

		if !delete {
			return ctx.JSON(500, helper.Response("Something Went Wrong!"))
		}

		return ctx.JSON(200, helper.Response("Book Success Deleted!"))
	}
}
