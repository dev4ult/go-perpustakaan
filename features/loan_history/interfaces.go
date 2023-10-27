package loan_history

import (
	"perpustakaan/features/loan_history/dtos"

	"github.com/labstack/echo/v4"
)

type Repository interface {
	Paginate(page, size int) []dtos.ResLoanHistory
	Insert(newLoanHistory LoanHistory) *dtos.ResLoanHistory
	SelectByID(loanHistoryID int) *dtos.ResLoanHistory
	Update(loanHistory LoanHistory) int64
	UpdateStatus(status, loanHistoryID int) int64
	DeleteByID(loanHistoryID int) int64
}

type Usecase interface {
	FindAll(page, size int) []dtos.ResLoanHistory
	FindByID(loanHistoryID int) *dtos.ResLoanHistory
	Create(newLoanHistory dtos.InputLoanHistory) *dtos.ResLoanHistory
	Modify(loanHistoryData dtos.InputLoanHistory, loanHistoryID int) bool
	ModifyStatus(status, loanHistoryID int) bool
	Remove(loanHistoryID int) bool
}

type Handler interface {
	GetLoanHistories() echo.HandlerFunc
	LoanHistoryDetails() echo.HandlerFunc
	CreateLoanHistory() echo.HandlerFunc
	UpdateLoanHistory() echo.HandlerFunc
	UpdateLoanStatus() echo.HandlerFunc
	DeleteLoanHistory() echo.HandlerFunc
}
