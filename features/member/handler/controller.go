package handler

import (
	"perpustakaan/helpers"
	"strconv"

	"perpustakaan/features/member"
	"perpustakaan/features/member/dtos"

	"github.com/labstack/echo/v4"
)

type controller struct {
	service member.Usecase
}

func New(service member.Usecase) member.Handler {
	return &controller {
		service: service,
	}
}

func (ctl *controller) GetMembers() echo.HandlerFunc {
	return func (ctx echo.Context) error  {
		pagination := dtos.Pagination{}
		ctx.Bind(&pagination)
		
		page := pagination.Page
		size := pagination.Size
		email := ctx.QueryParam("email")
		credentialNumber := ctx.QueryParam("cred-number")

		if size == 0 {
			page = 1
			size = 10
		}

		members, message := ctl.service.FindAll(page, size, email, credentialNumber)


		if message != "" {
			return ctx.JSON(500, helpers.Response(message))
		}

		if members == nil {
			return ctx.JSON(404, helpers.Response("There is No Members!"))
		}

		return ctx.JSON(200, helpers.Response("Success!", map[string]any {
			"data": members,
		}))
	}
}


func (ctl *controller) MemberDetails() echo.HandlerFunc {
	return func (ctx echo.Context) error  {
		memberID, err := strconv.Atoi(ctx.Param("id"))

		if err != nil {
			return ctx.JSON(400, helpers.Response("Param must be provided in number!"))
		}

		member, message := ctl.service.FindByID(memberID)

		if member == nil {
			return ctx.JSON(404, helpers.Response(message))
		}

		return ctx.JSON(200, helpers.Response("Success!", map[string]any {
			"data": member,
		}))
	}
}

func (ctl *controller) CreateMember() echo.HandlerFunc {
	return func (ctx echo.Context) error  {
		input := ctx.Get("request").(*dtos.InputMember)

		member, message := ctl.service.Create(*input)

		if member == nil {
			return ctx.JSON(500, helpers.Response(message))
		}

		return ctx.JSON(200, helpers.Response("Success!", map[string]any {
			"data": member,
		}))
	}
}

func (ctl *controller) UpdateMember() echo.HandlerFunc {
	return func (ctx echo.Context) error {
		input := ctx.Get("request").(*dtos.InputMember)

		memberID, errParam := strconv.Atoi(ctx.Param("id"))

		if errParam != nil {
			return ctx.JSON(400, helpers.Response(errParam.Error()))
		}

		member, message := ctl.service.FindByID(memberID)

		if member == nil {
			return ctx.JSON(404, helpers.Response(message))
		}
		
		update, messageUpdate := ctl.service.Modify(*input, memberID)

		if !update {
			return ctx.JSON(500, helpers.Response(messageUpdate))
		}

		return ctx.JSON(200, helpers.Response("Member Success Updated!"))
	}
}

func (ctl *controller) DeleteMember() echo.HandlerFunc {
	return func (ctx echo.Context) error  {
		memberID, err := strconv.Atoi(ctx.Param("id"))

		if err != nil {
			return ctx.JSON(400, helpers.Response("Param must be provided in number!"))
		}

		member, message := ctl.service.FindByID(memberID)

		if member == nil {
			return ctx.JSON(404, helpers.Response(message))
		}

		delete, messageDelete := ctl.service.Remove(memberID)

		if !delete {
			return ctx.JSON(500, helpers.Response(messageDelete))
		}

		return ctx.JSON(200, helpers.Response("Member Success Deleted!", nil))
	}
}
