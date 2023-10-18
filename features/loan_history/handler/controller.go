package handler

import (
	"perpustakaan/helpers"
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

func (ctl *controller) GetLoanHistorys() echo.HandlerFunc {
	return func (ctx echo.Context) error  {
		pagination := dtos.Pagination{}
		ctx.Bind(&pagination)
		
		page := pagination.Page
		size := pagination.Size

		if page <= 0 || size <= 0 {
			return ctx.JSON(400, helper.Response("Please provide query `page` and `size` in number!"))
		}

		loanHistorys := ctl.service.FindAll(page, size)

		if loanHistorys == nil {
			return ctx.JSON(404, helper.Response("There is No LoanHistorys!"))
		}

		return ctx.JSON(200, helper.Response("Success!", map[string]any {
			"data": loanHistorys,
		}))
	}
}


func (ctl *controller) LoanHistoryDetails() echo.HandlerFunc {
	return func (ctx echo.Context) error  {
		loanHistoryID, err := strconv.Atoi(ctx.Param("id"))

		if err != nil {
			return ctx.JSON(400, helper.Response(err.Error()))
		}

		loanHistory := ctl.service.FindByID(loanHistoryID)

		if loanHistory == nil {
			return ctx.JSON(404, helper.Response("LoanHistory Not Found!"))
		}

		return ctx.JSON(200, helper.Response("Success!", map[string]any {
			"data": loanHistory,
		}))
	}
}

func (ctl *controller) CreateLoanHistory() echo.HandlerFunc {
	return func (ctx echo.Context) error  {
		input := dtos.InputLoanHistory{}

		ctx.Bind(&input)

		validate = validator.New(validator.WithRequiredStructEnabled())

		err := validate.Struct(input)

		if err != nil {
			errMap := helpers.ErrorMapValidation(err)
			return ctx.JSON(400, helper.Response("Bad Request!", map[string]any {
				"error": errMap,
			}))
		}

		loanHistory := ctl.service.Create(input)

		if loanHistory == nil {
			return ctx.JSON(500, helper.Response("Something went Wrong!", nil))
		}

		return ctx.JSON(200, helper.Response("Success!", map[string]any {
			"data": loanHistory,
		}))
	}
}

func (ctl *controller) UpdateLoanHistory() echo.HandlerFunc {
	return func (ctx echo.Context) error {
		input := dtos.InputLoanHistory{}

		loanHistoryID, errParam := strconv.Atoi(ctx.Param("id"))

		if errParam != nil {
			return ctx.JSON(400, helper.Response(errParam.Error()))
		}

		loanHistory := ctl.service.FindByID(loanHistoryID)

		if loanHistory == nil {
			return ctx.JSON(404, helper.Response("LoanHistory Not Found!"))
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

		update := ctl.service.Modify(input, loanHistoryID)

		if !update {
			return ctx.JSON(500, helper.Response("Something Went Wrong!"))
		}

		return ctx.JSON(200, helper.Response("LoanHistory Success Updated!"))
	}
}

func (ctl *controller) DeleteLoanHistory() echo.HandlerFunc {
	return func (ctx echo.Context) error  {
		loanHistoryID, err := strconv.Atoi(ctx.Param("id"))

		if err != nil {
			return ctx.JSON(400, helper.Response(err.Error()))
		}

		loanHistory := ctl.service.FindByID(loanHistoryID)

		if loanHistory == nil {
			return ctx.JSON(404, helper.Response("LoanHistory Not Found!"))
		}

		delete := ctl.service.Remove(loanHistoryID)

		if !delete {
			return ctx.JSON(500, helper.Response("Something Went Wrong!"))
		}

		return ctx.JSON(200, helper.Response("LoanHistory Success Deleted!", nil))
	}
}
