package handler

import (
	"perpustakaan/helpers"
	helper "perpustakaan/helpers"
	"strconv"

	"perpustakaan/features/publisher"
	"perpustakaan/features/publisher/dtos"

	"github.com/labstack/echo/v4"
)

type controller struct {
	service publisher.Usecase
}

func New(service publisher.Usecase) publisher.Handler {
	return &controller {
		service: service,
	}
}

func (ctl *controller) GetPublishers() echo.HandlerFunc {
	return func (ctx echo.Context) error  {
		pagination := dtos.Pagination{}
		ctx.Bind(&pagination)
		
		page := pagination.Page
		size := pagination.Size

		if page <= 0 || size <= 0 {
			return ctx.JSON(400, helper.Response("Param must be provided in number!"))
		}

		publishers := ctl.service.FindAll(page, size)

		if publishers == nil {
			return ctx.JSON(404, helper.Response("There is No Publishers!"))
		}

		return ctx.JSON(200, helper.Response("Success!", map[string]any {
			"data": publishers,
		}))
	}
}


func (ctl *controller) PublisherDetails() echo.HandlerFunc {
	return func (ctx echo.Context) error  {
		publisherID, err := strconv.Atoi(ctx.Param("id"))

		if err != nil {
			return ctx.JSON(400, helper.Response("Param must be provided in number!"))
		}

		publisher := ctl.service.FindByID(publisherID)

		if publisher == nil {
			return ctx.JSON(404, helper.Response("Publisher Not Found!"))
		}

		return ctx.JSON(200, helper.Response("Success!", map[string]any {
			"data": publisher,
		}))
	}
}

func (ctl *controller) CreatePublisher() echo.HandlerFunc {
	return func (ctx echo.Context) error  {
		input := dtos.InputPublisher{}

		ctx.Bind(&input)

		if err := helpers.ValidateRequest(input); err != nil {
			errMap := helpers.ErrorMapValidation(err)
			return ctx.JSON(400, helper.Response("Missing Data Required!", map[string]any {
				"error": errMap,
			}))
		}

		publisher := ctl.service.Create(input)

		if publisher == nil {
			return ctx.JSON(500, helper.Response("Something went Wrong!", nil))
		}

		return ctx.JSON(200, helper.Response("Success!", map[string]any {
			"data": publisher,
		}))
	}
}

func (ctl *controller) UpdatePublisher() echo.HandlerFunc {
	return func (ctx echo.Context) error {
		input := dtos.InputPublisher{}

		publisherID, errParam := strconv.Atoi(ctx.Param("id"))

		if errParam != nil {
			return ctx.JSON(400, helper.Response(errParam.Error()))
		}

		publisher := ctl.service.FindByID(publisherID)

		if publisher == nil {
			return ctx.JSON(404, helper.Response("Publisher Not Found!"))
		}
		
		ctx.Bind(&input)
		
		if err := helpers.ValidateRequest(input); err != nil {
			errMap := helpers.ErrorMapValidation(err)
			return ctx.JSON(400, helper.Response("Missing Data Required!", map[string]any {
				"error": errMap,
			}))
		}

		update := ctl.service.Modify(input, publisherID)

		if !update {
			return ctx.JSON(500, helper.Response("Something Went Wrong!"))
		}

		return ctx.JSON(200, helper.Response("Publisher Success Updated!"))
	}
}

func (ctl *controller) DeletePublisher() echo.HandlerFunc {
	return func (ctx echo.Context) error  {
		publisherID, err := strconv.Atoi(ctx.Param("id"))

		if err != nil {
			return ctx.JSON(400, helper.Response("Param must be provided in number!"))
		}

		publisher := ctl.service.FindByID(publisherID)

		if publisher == nil {
			return ctx.JSON(404, helper.Response("Publisher Not Found!"))
		}

		delete := ctl.service.Remove(publisherID)

		if !delete {
			return ctx.JSON(500, helper.Response("Something Went Wrong!"))
		}

		return ctx.JSON(200, helper.Response("Publisher Success Deleted!", nil))
	}
}
