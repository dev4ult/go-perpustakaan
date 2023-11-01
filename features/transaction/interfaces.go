package transaction

import (
	"perpustakaan/features/member"
	"perpustakaan/features/transaction/dtos"

	"github.com/labstack/echo/v4"
)

type Repository interface {
	Paginate(page, size int) ([]Transaction, error)
	Insert(newTransaction Transaction) (int, error)
	SelectByID(transactionID int) (*Transaction, error)
	Update(transaction Transaction) (int, error)
	DeleteByID(transactionID int) (int, error)
	UpdateBatchTransactionDetail(items []dtos.FineItem, transactionID int) (bool, error)
	SelectAllFineItemOnMemberID(memberID int) ([]dtos.FineItem, error)
	SelectAllFineItemOnTransactionID(TransactionID int) ([]dtos.FineItem, error)
	SelectFineItemByIDAndMemberID(fineItemID, memberID int) (*dtos.FineItem, error)
	SelectMemberByID(memberID int) (*member.Member, error)
	SelectTransactionByOrderID(orderID string) (*Transaction, error)
	UpdateStatus(transactionID int, status string) (bool, error)
	UnsetTransactionIDs(transactionID int) (bool, error)
}

type Usecase interface {
	FindAll(page, size int) ([]dtos.ResTransaction, string)
	FindByID(transactionID int) (*dtos.ResTransaction, string)
	Create(newTransaction dtos.InputTransaction) (*dtos.ResTransaction, string)
	Modify(transactionData dtos.InputTransaction, transactionID int, orderID string, status string, paymentURL string) (bool, string)
	Remove(transactionID int) (bool, string)
	VerifyPayment(payload map[string]any) (bool, string)
}

type Handler interface {
	GetTransactions() echo.HandlerFunc
	TransactionDetails() echo.HandlerFunc
	CreateTransaction() echo.HandlerFunc
	UpdateTransaction() echo.HandlerFunc
	DeleteTransaction() echo.HandlerFunc
	Notification() echo.HandlerFunc
}
