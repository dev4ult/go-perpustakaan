package routes

import (
	"perpustakaan/features/transaction"
	"perpustakaan/features/transaction/dtos"
	m "perpustakaan/middlewares"

	"github.com/labstack/echo/v4"
)

func Transactions(e *echo.Echo, handler transaction.Handler) {
	transactions := e.Group("/transactions")
	transactions.Use(m.Authorization("librarian"))

	transactions.GET("", handler.GetTransactions())
	transactions.POST("", handler.CreateTransaction(), m.RequestValidation(dtos.InputTransaction{}))
	
	transactions.GET("/:id", handler.TransactionDetails())
	transactions.PUT("/:id", handler.UpdateTransaction(), m.RequestValidation(dtos.InputTransaction{}))
	transactions.DELETE("/:id", handler.DeleteTransaction())

	notification := e.Group("/notification")
	notification.POST("", handler.Notification())
}