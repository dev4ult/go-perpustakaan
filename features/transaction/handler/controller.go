package handler

import (
	"encoding/json"
	"perpustakaan/helpers"
	helper "perpustakaan/helpers"
	"strconv"

	"perpustakaan/features/transaction"
	"perpustakaan/features/transaction/dtos"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type controller struct {
	service transaction.Usecase
}

func New(service transaction.Usecase) transaction.Handler {
	return &controller {
		service: service,
	}
}

var validate *validator.Validate

func (ctl *controller) GetTransactions() echo.HandlerFunc {
	return func (ctx echo.Context) error  {
		pagination := dtos.Pagination{}
		ctx.Bind(&pagination)
		
		page := pagination.Page
		size := pagination.Size

		if page <= 0 || size <= 0 {
			return ctx.JSON(400, helper.Response("Param must be provided in number!"))
		}

		transactions, message := ctl.service.FindAll(page, size)

		if message != "" {
			return ctx.JSON(500, helper.Response(message))
		}

		if transactions == nil {
			return ctx.JSON(404, helper.Response("There is No Transactions!"))
		}

		return ctx.JSON(200, helper.Response("Success!", map[string]any {
			"data": transactions,
		}))
	}
}


func (ctl *controller) TransactionDetails() echo.HandlerFunc {
	return func (ctx echo.Context) error  {
		transactionID, err := strconv.Atoi(ctx.Param("id"))

		if err != nil {
			return ctx.JSON(400, helper.Response("Param must be provided in number!"))
		}

		transaction, message := ctl.service.FindByID(transactionID)

		if transaction == nil {
			return ctx.JSON(404, helper.Response(message))
		}

		return ctx.JSON(200, helper.Response("Success!", map[string]any {
			"data": transaction,
		}))
	}
}

func (ctl *controller) CreateTransaction() echo.HandlerFunc {
	return func (ctx echo.Context) error  {
		input := ctx.Get("request").(*dtos.InputTransaction)

		transaction, message := ctl.service.Create(*input)

		if transaction == nil {
			return ctx.JSON(500, helper.Response(message))
		}

		return ctx.JSON(200, helper.Response("Success!", map[string]any {
			"data": transaction,
		}))
	}
}

func (ctl *controller) UpdateTransaction() echo.HandlerFunc {
	return func (ctx echo.Context) error {
		input := ctx.Get("request").(*dtos.InputTransaction)
		transactionID, errParam := strconv.Atoi(ctx.Param("id"))

		if errParam != nil {
			return ctx.JSON(400, helper.Response(errParam.Error()))
		}

		transaction, message := ctl.service.FindByID(transactionID)

		if transaction == nil {
			return ctx.JSON(404, helper.Response(message))
		}
		
		update, updateMessage := ctl.service.Modify(*input, transactionID, transaction.OrderID, transaction.Status, transaction.PaymentURL)

		if !update {
			return ctx.JSON(500, helper.Response(updateMessage))
		}

		return ctx.JSON(200, helper.Response("Transaction Success Updated!"))
	}
}

func (ctl *controller) DeleteTransaction() echo.HandlerFunc {
	return func (ctx echo.Context) error  {
		transactionID, err := strconv.Atoi(ctx.Param("id"))

		if err != nil {
			return ctx.JSON(400, helper.Response("Param must be provided in number!"))
		}

		transaction, message := ctl.service.FindByID(transactionID)

		if transaction == nil {
			return ctx.JSON(404, helper.Response(message))
		}

		delete, deleteMessage := ctl.service.Remove(transactionID)

		if !delete {
			return ctx.JSON(500, helper.Response(deleteMessage))
		}

		return ctx.JSON(200, helper.Response("Transaction Success Deleted!"))
	}
}

func (ctl *controller) Notification() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		var payload map[string]any

		if err := json.NewDecoder(ctx.Request().Body).Decode(&payload); err != nil {
			return ctx.JSON(400, helpers.Response("Error parsing data"))
		}

		verified, errMessage := ctl.service.VerifyPayment(payload)

		if !verified {
			return ctx.JSON(500, helpers.Response(errMessage))
		}

		return ctx.JSON(200, helpers.Response("Payment Verified"))
	}
}