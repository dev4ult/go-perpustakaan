package loan_history

import (
	"perpustakaan/features/loan_history/dtos"

	"github.com/labstack/echo/v4"
)

type Repository interface {
	Paginate(page int, size int, memberName string, status string) ([]dtos.ResLoanHistory, error)
	Insert(newLoanHistory LoanHistory)(*dtos.ResLoanHistory, error)
	SelectByID(loanHistoryID int) (*dtos.ResLoanHistory, error)
	Update(loanHistory LoanHistory) (int, error)
	UpdateStatus(status int, statusBefore string, loanHistoryID int) (int, error)
	DeleteByID(loanHistoryID int) (int, error)
}

type Usecase interface {
	FindAll(page int, size int, memberName string, status string) ([]dtos.ResLoanHistory, string)
	FindByID(loanHistoryID int) (*dtos.ResLoanHistory, string)
	Create(newLoanHistory dtos.InputLoanHistory) (*dtos.ResLoanHistory, string)
	Modify(loanHistoryData dtos.InputLoanHistory, loanHistoryID int)( bool, string)
	ModifyStatus(status int, statusBefore string, loanHistoryID int) (bool, string)
	Remove(loanHistoryID int) (bool, string)
}

type Handler interface {
	GetLoanHistories() echo.HandlerFunc
	LoanHistoryDetails() echo.HandlerFunc
	CreateLoanHistory() echo.HandlerFunc
	UpdateLoanHistory() echo.HandlerFunc
	UpdateLoanStatus() echo.HandlerFunc
	DeleteLoanHistory() echo.HandlerFunc
}
