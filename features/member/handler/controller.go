package handler

import (
	"perpustakaan/helpers"
	helper "perpustakaan/helpers"
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

		if page <= 0 || size <= 0 {
			return ctx.JSON(400, helper.Response("Please provide query `page` and `size` in number!"))
		}

		members := ctl.service.FindAll(page, size)

		if members == nil {
			return ctx.JSON(404, helper.Response("There is No Members!"))
		}

		return ctx.JSON(200, helper.Response("Success!", map[string]any {
			"data": members,
		}))
	}
}


func (ctl *controller) MemberDetails() echo.HandlerFunc {
	return func (ctx echo.Context) error  {
		memberID, err := strconv.Atoi(ctx.Param("id"))

		if err != nil {
			return ctx.JSON(400, helper.Response("Please provide query `page` and `size` in number!"))
		}

		member := ctl.service.FindByID(memberID)

		if member == nil {
			return ctx.JSON(404, helper.Response("Member Not Found!"))
		}

		return ctx.JSON(200, helper.Response("Success!", map[string]any {
			"data": member,
		}))
	}
}

func (ctl *controller) CreateMember() echo.HandlerFunc {
	return func (ctx echo.Context) error  {
		input := dtos.InputMember{}

		ctx.Bind(&input)

		if err := helpers.ValidateRequest(input); err != nil {
			errMap := helpers.ErrorMapValidation(err)
			return ctx.JSON(400, helper.Response("Missing Data Required!", map[string]any {
				"error": errMap,
			}))
		}

		member := ctl.service.Create(input)

		if member == nil {
			return ctx.JSON(500, helper.Response("Something went Wrong!", nil))
		}

		return ctx.JSON(200, helper.Response("Success!", map[string]any {
			"data": member,
		}))
	}
}

func (ctl *controller) UpdateMember() echo.HandlerFunc {
	return func (ctx echo.Context) error {
		input := dtos.InputMember{}

		memberID, errParam := strconv.Atoi(ctx.Param("id"))

		if errParam != nil {
			return ctx.JSON(400, helper.Response(errParam.Error()))
		}

		member := ctl.service.FindByID(memberID)

		if member == nil {
			return ctx.JSON(404, helper.Response("Member Not Found!"))
		}
		
		ctx.Bind(&input)

		if err := helpers.ValidateRequest(input); err != nil {
			errMap := helpers.ErrorMapValidation(err)
			return ctx.JSON(400, helper.Response("Missing Data Required!", map[string]any {
				"error": errMap,
			}))
		}

		update := ctl.service.Modify(input, memberID)

		if !update {
			return ctx.JSON(500, helper.Response("Something Went Wrong!"))
		}

		return ctx.JSON(200, helper.Response("Member Success Updated!"))
	}
}

func (ctl *controller) DeleteMember() echo.HandlerFunc {
	return func (ctx echo.Context) error  {
		memberID, err := strconv.Atoi(ctx.Param("id"))

		if err != nil {
			return ctx.JSON(400, helper.Response("Please provide query `page` and `size` in number!"))
		}

		member := ctl.service.FindByID(memberID)

		if member == nil {
			return ctx.JSON(404, helper.Response("Member Not Found!"))
		}

		delete := ctl.service.Remove(memberID)

		if !delete {
			return ctx.JSON(500, helper.Response("Something Went Wrong!"))
		}

		return ctx.JSON(200, helper.Response("Member Success Deleted!", nil))
	}
}
