package transaction

import (
	"perpustakaan/features/member"
	"perpustakaan/features/transaction/dtos"

	"github.com/labstack/echo/v4"
)

type Repository interface {
	Paginate(page, size int) []Transaction
	Insert(newTransaction Transaction) int64
	SelectByID(transactionID int) *Transaction
	Update(transaction Transaction) int64
	DeleteByID(transactionID int) int64
	UpdateBatchTransactionDetail(items []dtos.FineItem, transactionID int64) bool
	SelectAllFineItemOnMemberID(memberID int) []dtos.FineItem
	SelectAllFineItemOnTransactionID(TransactionID int) []dtos.FineItem
	SelectFineItemByIDAndMemberID(fineItemID, memberID int) *dtos.FineItem
	SelectMemberByID(memberID int) *member.Member
	SelectTransactionByOrderID(orderID string) *Transaction
	UpdateStatus(transactionID int, status string) bool
}

type Usecase interface {
	FindAll(page, size int) []dtos.ResTransaction
	FindByID(transactionID int) *dtos.ResTransaction
	Create(newTransaction dtos.InputTransaction) (*dtos.ResTransaction, string)
	Modify(transactionData dtos.InputTransaction, transactionID int) bool
	Remove(transactionID int) bool
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