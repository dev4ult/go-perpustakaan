package handler

import (
	helper "perpustakaan/helpers"
	"strconv"

	"perpustakaan/features/author"
	"perpustakaan/features/author/dtos"

	"github.com/labstack/echo/v4"
)

type controller struct {
	service author.Usecase
}

func New(service author.Usecase) author.Handler {
	return &controller {
		service: service,
	}
}


func (ctl *controller) GetAuthors() echo.HandlerFunc {
	return func (ctx echo.Context) error  {
		pagination := dtos.Pagination{}
		ctx.Bind(&pagination)
		
		page := pagination.Page
		size := pagination.Size

		if page <= 0 || size <= 0 {
			page = 1
			size = 10
		}
		
		authors, message := ctl.service.FindAll(page, size)
		
		if message != "" {
			return ctx.JSON(500, helper.Response(message))
		}

		if authors == nil {
			return ctx.JSON(404, helper.Response("There is No Authors!"))
		}

		return ctx.JSON(200, helper.Response("Success!", map[string]any {
			"data": authors,
		}))
	}
}


func (ctl *controller) AuthorDetails() echo.HandlerFunc {
	return func (ctx echo.Context) error  {
		authorID, err := strconv.Atoi(ctx.Param("id"))

		if err != nil {
			return ctx.JSON(400, helper.Response("Param must be provided in number!"))
		}

		author, message := ctl.service.FindByID(authorID)

		if author == nil {
			return ctx.JSON(404, helper.Response(message))
		}

		return ctx.JSON(200, helper.Response("Success!", map[string]any {
			"data": author,
		}))
	}
}

func (ctl *controller) CreateAuthor() echo.HandlerFunc {
	return func (ctx echo.Context) error  {
		input := ctx.Get("request").(*dtos.InputAuthor)

		author, message := ctl.service.Create(*input)

		if author == nil {
			return ctx.JSON(500, helper.Response(message))
		}

		return ctx.JSON(200, helper.Response("Success!", map[string]any {
			"data": author,
		}))
	}
}

func (ctl *controller) UpdateAuthor() echo.HandlerFunc {
	return func (ctx echo.Context) error {
		input := ctx.Get("request").(*dtos.InputAuthor)

		authorID, errParam := strconv.Atoi(ctx.Param("id"))

		if errParam != nil {
			return ctx.JSON(400, helper.Response(errParam.Error()))
		}

		author, message := ctl.service.FindByID(authorID)

		if author == nil {
			return ctx.JSON(404, helper.Response(message))
		}

		update, updateMessage := ctl.service.Modify(*input, authorID)

		if !update {
			return ctx.JSON(500, helper.Response(updateMessage))
		}

		return ctx.JSON(200, helper.Response("Author Success Updated!"))
	}
}

func (ctl *controller) DeleteAuthor() echo.HandlerFunc {
	return func (ctx echo.Context) error  {
		authorID, err := strconv.Atoi(ctx.Param("id"))

		if err != nil {
			return ctx.JSON(400, helper.Response("Param must be provided in number!"))
		}

		author, message := ctl.service.FindByID(authorID)

		if author == nil {
			return ctx.JSON(404, helper.Response(message))
		}

		delete, deleteMessage := ctl.service.Remove(authorID)

		if !delete {
			return ctx.JSON(500, helper.Response(deleteMessage))
		}

		return ctx.JSON(200, helper.Response("Author Success Deleted!", nil))
	}
}

func (ctl *controller) CreateAnAuthorship() echo.HandlerFunc {
	return func (ctx echo.Context) error  {
		input := ctx.Get("request").(*dtos.InputAuthorshipIDS)
	
		author, err := ctl.service.SetupAuthorship(*input)

		if author == nil {
			return ctx.JSON(500, helper.Response(err))
		}

		return ctx.JSON(200, helper.Response("Success!", map[string]any {
			"data": author,
		}))
	}
}

func (ctl *controller) DeleteAnAuthorship() echo.HandlerFunc {
	return func (ctx echo.Context) error  {
		authorshipID, err := strconv.Atoi(ctx.Param("id"))

		if err != nil {
			return ctx.JSON(400, helper.Response("Param must be provided in number!"))
		}

		if exist, _ := ctl.service.IsAuthorshipExistByID(authorshipID); !exist {
			return ctx.JSON(404, helper.Response("Authorship Not Found!"))
		}

		delete, message := ctl.service.RemoveAuthorship(authorshipID)

		if !delete {
			return ctx.JSON(500, helper.Response(message))
		}

		return ctx.JSON(200, helper.Response("Authorship Success Deleted!", nil))
	}
}