package loanHistory

import (
	"perpustakaan/features/loanHistory/dtos"

	"github.com/labstack/echo/v4"
)

type Repository interface {
	Paginate(page, size int) []LoanHistory
	Insert(newLoanHistory LoanHistory) int64
	SelectByID(loanHistoryID int) *LoanHistory
	Update(loanHistory LoanHistory) int64
	DeleteByID(loanHistoryID int) int64
}

type Usecase interface {
	FindAll(page, size int) []dtos.ResLoanHistory
	FindByID(loanHistoryID int) *dtos.ResLoanHistory
	Create(newLoanHistory dtos.InputLoanHistory) *dtos.ResLoanHistory
	Modify(loanHistoryData dtos.InputLoanHistory, loanHistoryID int) bool
	Remove(loanHistoryID int) bool
}

type Handler interface {
	GetLoanHistorys() echo.HandlerFunc
	LoanHistoryDetails() echo.HandlerFunc
	CreateLoanHistory() echo.HandlerFunc
	UpdateLoanHistory() echo.HandlerFunc
	DeleteLoanHistory() echo.HandlerFunc
}
