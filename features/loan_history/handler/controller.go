package handler

import (
	helper "perpustakaan/helpers"
	"strconv"

	"perpustakaan/features/loan_history"
	"perpustakaan/features/loan_history/dtos"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type controller struct {
	service loan_history.Usecase
}

func New(service loan_history.Usecase) loan_history.Handler {
	return &controller {
		service: service,
	}
}

var validate *validator.Validate

func (ctl *controller) GetLoanHistories() echo.HandlerFunc {
	return func (ctx echo.Context) error  {
		pagination := dtos.Pagination{}
		ctx.Bind(&pagination)
		
		page := pagination.Page
		size := pagination.Size
		searchKey := ctx.QueryParam("member")

		if page <= 0 || size <= 0 {
			page = 1
			size = 10
		}

		loanHistories, message := ctl.service.FindAll(page, size, searchKey)

		if message != "" {
			return ctx.JSON(500, helper.Response(message))
		}

		if len(loanHistories) == 0 {
			return ctx.JSON(404, helper.Response("There is No Loan Histories!"))
		}

		return ctx.JSON(200, helper.Response("Success!", map[string]any {
			"data": loanHistories,
		}))
	}
}


func (ctl *controller) LoanHistoryDetails() echo.HandlerFunc {
	return func (ctx echo.Context) error  {
		loanHistoryID, err := strconv.Atoi(ctx.Param("id"))

		if err != nil {
			return ctx.JSON(400, helper.Response("Param must be provided in number!"))
		}

		loanHistory, message := ctl.service.FindByID(loanHistoryID)

		if loanHistory == nil {
			return ctx.JSON(404, helper.Response(message))
		}

		return ctx.JSON(200, helper.Response("Success!", map[string]any {
			"data": loanHistory,
		}))
	}
}

func (ctl *controller) CreateLoanHistory() echo.HandlerFunc {
	return func (ctx echo.Context) error  {
		input := ctx.Get("request").(*dtos.InputLoanHistory)

		loanHistory, message := ctl.service.Create(*input)

		if loanHistory == nil {
			return ctx.JSON(500, helper.Response(message))
		}

		return ctx.JSON(200, helper.Response("Success!", map[string]any {
			"data": loanHistory,
		}))
	}
}

func (ctl *controller) UpdateLoanHistory() echo.HandlerFunc {
	return func (ctx echo.Context) error {
		input := ctx.Get("request").(*dtos.InputLoanHistory)
		loanHistoryID, errParam := strconv.Atoi(ctx.Param("id"))

		if errParam != nil {
			return ctx.JSON(400, helper.Response(errParam.Error()))
		}

		loanHistory, message := ctl.service.FindByID(loanHistoryID)

		if loanHistory == nil {
			return ctx.JSON(404, helper.Response(message))
		}
		
		update, updateMessage := ctl.service.Modify(*input, loanHistoryID)

		if !update {
			return ctx.JSON(500, helper.Response(updateMessage))
		}

		return ctx.JSON(200, helper.Response("Loan History Success Updated!"))
	}
}

func (ctl *controller) UpdateLoanStatus() echo.HandlerFunc {
	return func (ctx echo.Context) error {
		input := ctx.Get("request").(*dtos.LoanStatus)
		loanHistoryID, errParam := strconv.Atoi(ctx.Param("id"))

		if errParam != nil {
			return ctx.JSON(400, helper.Response(errParam.Error()))
		}

		loanHistory, message := ctl.service.FindByID(loanHistoryID)

		if loanHistory == nil {
			return ctx.JSON(404, helper.Response(message))
		}
		
		patch, patchMessage := ctl.service.ModifyStatus(input.Status, loanHistoryID)

		if !patch {
			return ctx.JSON(500, helper.Response(patchMessage))
		}

		return ctx.JSON(200, helper.Response("Loan Status Success Updated!"))
	}
}

func (ctl *controller) DeleteLoanHistory() echo.HandlerFunc {
	return func (ctx echo.Context) error  {
		loanHistoryID, err := strconv.Atoi(ctx.Param("id"))

		if err != nil {
			return ctx.JSON(400, helper.Response("Param must be provided in number!"))
		}

		loanHistory, message := ctl.service.FindByID(loanHistoryID)

		if loanHistory == nil {
			return ctx.JSON(404, helper.Response(message))
		}

		delete, deleteMessage := ctl.service.Remove(loanHistoryID)

		if !delete {
			return ctx.JSON(500, helper.Response(deleteMessage))
		}

		return ctx.JSON(200, helper.Response("Loan History Success Deleted!", nil))
	}
}
