package handler

import (
	helper "perpustakaan/helpers"
	"strconv"

	"perpustakaan/features/_blueprint"
	"perpustakaan/features/_blueprint/dtos"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type controller struct {
	service _blueprint.Usecase
}

func New(service _blueprint.Usecase) _blueprint.Handler {
	return &controller {
		service: service,
	}
}

var validate *validator.Validate

func (ctl *controller) GetPlaceholders() echo.HandlerFunc {
	return func (ctx echo.Context) error  {
		pagination := dtos.Pagination{}
		ctx.Bind(&pagination)
		
		page := pagination.Page
		size := pagination.Size

		if page <= 0 || size <= 0 {
			return ctx.JSON(400, helper.Response("Param must be provided in number!"))
		}

		placeholders := ctl.service.FindAll(page, size)

		if placeholders == nil {
			return ctx.JSON(404, helper.Response("There is No Placeholders!"))
		}

		return ctx.JSON(200, helper.Response("Success!", map[string]any {
			"data": placeholders,
		}))
	}
}


func (ctl *controller) PlaceholderDetails() echo.HandlerFunc {
	return func (ctx echo.Context) error  {
		placeholderID, err := strconv.Atoi(ctx.Param("id"))

		if err != nil {
			return ctx.JSON(400, helper.Response("Param must be provided in number!"))
		}

		placeholder := ctl.service.FindByID(placeholderID)

		if placeholder == nil {
			return ctx.JSON(404, helper.Response("Placeholder Not Found!"))
		}

		return ctx.JSON(200, helper.Response("Success!", map[string]any {
			"data": placeholder,
		}))
	}
}

func (ctl *controller) CreatePlaceholder() echo.HandlerFunc {
	return func (ctx echo.Context) error  {
		input := ctx.Get("request").(*dtos.InputPlaceholder)

		placeholder := ctl.service.Create(*input)

		if placeholder == nil {
			return ctx.JSON(500, helper.Response("Something went Wrong!", nil))
		}

		return ctx.JSON(200, helper.Response("Success!", map[string]any {
			"data": placeholder,
		}))
	}
}

func (ctl *controller) UpdatePlaceholder() echo.HandlerFunc {
	return func (ctx echo.Context) error {
		input := ctx.Get("request").(*dtos.InputPlaceholder)
		placeholderID, errParam := strconv.Atoi(ctx.Param("id"))

		if errParam != nil {
			return ctx.JSON(400, helper.Response(errParam.Error()))
		}

		placeholder := ctl.service.FindByID(placeholderID)

		if placeholder == nil {
			return ctx.JSON(404, helper.Response("Placeholder Not Found!"))
		}
		
		update := ctl.service.Modify(*input, placeholderID)

		if !update {
			return ctx.JSON(500, helper.Response("Something Went Wrong!"))
		}

		return ctx.JSON(200, helper.Response("Placeholder Success Updated!"))
	}
}

func (ctl *controller) DeletePlaceholder() echo.HandlerFunc {
	return func (ctx echo.Context) error  {
		placeholderID, err := strconv.Atoi(ctx.Param("id"))

		if err != nil {
			return ctx.JSON(400, helper.Response("Param must be provided in number!"))
		}

		placeholder := ctl.service.FindByID(placeholderID)

		if placeholder == nil {
			return ctx.JSON(404, helper.Response("Placeholder Not Found!"))
		}

		delete := ctl.service.Remove(placeholderID)

		if !delete {
			return ctx.JSON(500, helper.Response("Something Went Wrong!"))
		}

		return ctx.JSON(200, helper.Response("Placeholder Success Deleted!", nil))
	}
}
