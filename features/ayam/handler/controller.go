package handler

import (
	"perpustakaan/helpers"
	helper "perpustakaan/helpers"
	"strconv"

	"perpustakaan/features/ayam"
	"perpustakaan/features/ayam/dtos"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type controller struct {
	service ayam.Usecase
}

func New(service ayam.Usecase) ayam.Handler {
	return &controller {
		service: service,
	}
}

var validate *validator.Validate

func (ctl *controller) GetAyams() echo.HandlerFunc {
	return func (ctx echo.Context) error  {
		pagination := dtos.Pagination{}
		ctx.Bind(&pagination)
		
		page := pagination.Page
		size := pagination.Size

		if page <= 0 || size <= 0 {
			return ctx.JSON(400, helper.Response("Please provide query `page` and `size` in number!"))
		}

		ayams := ctl.service.FindAll(page, size)

		if ayams == nil {
			return ctx.JSON(404, helper.Response("There is No Ayams!"))
		}

		return ctx.JSON(200, helper.Response("Success!", map[string]any {
			"data": ayams,
		}))
	}
}


func (ctl *controller) AyamDetails() echo.HandlerFunc {
	return func (ctx echo.Context) error  {
		ayamID, err := strconv.Atoi(ctx.Param("id"))

		if err != nil {
			return ctx.JSON(400, helper.Response(err.Error()))
		}

		ayam := ctl.service.FindByID(ayamID)

		if ayam == nil {
			return ctx.JSON(404, helper.Response("Ayam Not Found!"))
		}

		return ctx.JSON(200, helper.Response("Success!", map[string]any {
			"data": ayam,
		}))
	}
}

func (ctl *controller) CreateAyam() echo.HandlerFunc {
	return func (ctx echo.Context) error  {
		input := dtos.InputAyam{}

		ctx.Bind(&input)

		validate = validator.New(validator.WithRequiredStructEnabled())

		err := validate.Struct(input)

		if err != nil {
			errMap := helpers.ErrorMapValidation(err)
			return ctx.JSON(400, helper.Response("Bad Request!", map[string]any {
				"error": errMap,
			}))
		}

		ayam := ctl.service.Create(input)

		if ayam == nil {
			return ctx.JSON(500, helper.Response("Something went Wrong!", nil))
		}

		return ctx.JSON(200, helper.Response("Success!", map[string]any {
			"data": ayam,
		}))
	}
}

func (ctl *controller) UpdateAyam() echo.HandlerFunc {
	return func (ctx echo.Context) error {
		input := dtos.InputAyam{}

		ayamID, errParam := strconv.Atoi(ctx.Param("id"))

		if errParam != nil {
			return ctx.JSON(400, helper.Response(errParam.Error()))
		}

		ayam := ctl.service.FindByID(ayamID)

		if ayam == nil {
			return ctx.JSON(404, helper.Response("Ayam Not Found!"))
		}
		
		ctx.Bind(&input)

		validate = validator.New(validator.WithRequiredStructEnabled())
		err := validate.Struct(input)

		if err != nil {
			errMap := helpers.ErrorMapValidation(err)
			return ctx.JSON(400, helper.Response("Bad Request!", map[string]any {
				"error": errMap,
			}))
		}

		update := ctl.service.Modify(input, ayamID)

		if !update {
			return ctx.JSON(500, helper.Response("Something Went Wrong!"))
		}

		return ctx.JSON(200, helper.Response("Ayam Success Updated!"))
	}
}

func (ctl *controller) DeleteAyam() echo.HandlerFunc {
	return func (ctx echo.Context) error  {
		ayamID, err := strconv.Atoi(ctx.Param("id"))

		if err != nil {
			return ctx.JSON(400, helper.Response(err.Error()))
		}

		ayam := ctl.service.FindByID(ayamID)

		if ayam == nil {
			return ctx.JSON(404, helper.Response("Ayam Not Found!"))
		}

		delete := ctl.service.Remove(ayamID)

		if !delete {
			return ctx.JSON(500, helper.Response("Something Went Wrong!"))
		}

		return ctx.JSON(200, helper.Response("Ayam Success Deleted!", nil))
	}
}
