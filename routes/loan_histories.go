package routes

import (
	"perpustakaan/features/loan_history"
	"perpustakaan/features/loan_history/dtos"
	m "perpustakaan/middlewares"

	"github.com/labstack/echo/v4"
)

func LoanHistories(e *echo.Echo, handler loan_history.Handler) {
	loanHistories := e.Group("/loan-histories")
	loanHistories.Use(m.Authorization("librarian"))

	loanHistories.GET("", handler.GetLoanHistories())
	loanHistories.POST("", handler.CreateLoanHistory(), m.RequestValidation(&dtos.InputLoanHistory{}))
	
	loanHistories.GET("/:id", handler.LoanHistoryDetails())
	loanHistories.PUT("/:id", handler.UpdateLoanHistory(), m.RequestValidation(&dtos.InputLoanHistory{}))
	loanHistories.PATCH("/:id", handler.UpdateLoanStatus(), m.RequestValidation(&dtos.LoanStatus{}))
	loanHistories.DELETE("/:id", handler.DeleteLoanHistory())
}