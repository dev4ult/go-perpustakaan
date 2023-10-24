package routes

import (
	"perpustakaan/features/loan_history"

	"github.com/labstack/echo/v4"
)

func LoanHistories(e *echo.Echo, handler loan_history.Handler) {
	loanHistorys := e.Group("/loan_histories")

	loanHistorys.GET("", handler.GetLoanHistorys())
	loanHistorys.POST("", handler.CreateLoanHistory())
	
	loanHistorys.GET("/:id", handler.LoanHistoryDetails())
	loanHistorys.PUT("/:id", handler.UpdateLoanHistory())
	loanHistorys.DELETE("/:id", handler.DeleteLoanHistory())
}