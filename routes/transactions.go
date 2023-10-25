package routes

import (
	"perpustakaan/features/transaction"

	"github.com/labstack/echo/v4"
)

func Transactions(e *echo.Echo, handler transaction.Handler) {
	transactions := e.Group("/transactions")

	transactions.GET("", handler.GetTransactions())
	transactions.POST("", handler.CreateTransaction())
	
	transactions.GET("/:id", handler.TransactionDetails())
	transactions.PUT("/:id", handler.UpdateTransaction())
	transactions.DELETE("/:id", handler.DeleteTransaction())
}