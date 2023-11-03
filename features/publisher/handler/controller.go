package handler

import (
	"strconv"

	"perpustakaan/features/publisher"
	"perpustakaan/features/publisher/dtos"
	"perpustakaan/helpers"

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
		searchKey := ctx.QueryParam("name")

		if size == 0 {
			page = 1
			size = 10
		}

		publishers, message := ctl.service.FindAll(page, size, searchKey)
		
		if message != "" {
			return ctx.JSON(500, helpers.Response(message))
		}

		if len(publishers) == 0 {
			return ctx.JSON(404, helpers.Response("There is No Publishers!"))
		}

		return ctx.JSON(200, helpers.Response("Success!", map[string]any {
			"data": publishers,
		}))
	}
}


func (ctl *controller) PublisherDetails() echo.HandlerFunc {
	return func (ctx echo.Context) error  {
		publisherID, err := strconv.Atoi(ctx.Param("id"))

		if err != nil {
			return ctx.JSON(400, helpers.Response("Param must be provided in number!"))
		}

		publisher, message := ctl.service.FindByID(publisherID)

		if publisher == nil {
			return ctx.JSON(404, helpers.Response(message))
		}

		return ctx.JSON(200, helpers.Response("Success!", map[string]any {
			"data": publisher,
		}))
	}
}

func (ctl *controller) CreatePublisher() echo.HandlerFunc {
	return func (ctx echo.Context) error  {
		input := ctx.Get("request").(*dtos.InputPublisher)

		publisher, message := ctl.service.Create(*input)

		if publisher == nil {
			return ctx.JSON(500, helpers.Response(message))
		}

		return ctx.JSON(200, helpers.Response("Success!", map[string]any {
			"data": publisher,
		}))
	}
}

func (ctl *controller) UpdatePublisher() echo.HandlerFunc {
	return func (ctx echo.Context) error {
		input := ctx.Get("request").(*dtos.InputPublisher)

		publisherID, err := strconv.Atoi(ctx.Param("id"))

		if err != nil {
			return ctx.JSON(400, helpers.Response(err.Error()))
		}

		publisher, message := ctl.service.FindByID(publisherID)

		if publisher == nil {
			return ctx.JSON(404, helpers.Response(message))
		}
		
		update, updateMessage := ctl.service.Modify(*input, publisherID)

		if !update {
			return ctx.JSON(500, helpers.Response(updateMessage))
		}

		return ctx.JSON(200, helpers.Response("Publisher Success Updated!"))
	}
}

func (ctl *controller) DeletePublisher() echo.HandlerFunc {
	return func (ctx echo.Context) error  {
		publisherID, err := strconv.Atoi(ctx.Param("id"))

		if err != nil {
			return ctx.JSON(400, helpers.Response("Param must be provided in number!"))
		}

		publisher, message := ctl.service.FindByID(publisherID)

		if publisher == nil {
			return ctx.JSON(404, helpers.Response(message))
		}

		delete, deleteMessage := ctl.service.Remove(publisherID)

		if !delete {
			return ctx.JSON(500, helpers.Response(deleteMessage))
		}

		return ctx.JSON(200, helpers.Response("Publisher Success Deleted!", nil))
	}
}
