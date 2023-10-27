package loan_history

import (
	"perpustakaan/features/book"
	"perpustakaan/features/member"
	"perpustakaan/features/transaction"

	"gorm.io/gorm"
)

type LoanHistory struct {
	gorm.Model

	ID         		int `gorm:"type:int(11)"`
	StartToLoanAt  	string `gorm:"type:date"`
	DueDate   		string `gorm:"type:date"`
	BookID     		int `gorm:"type:int(11)"`
	MemberID   		int `gorm:"type:int(11)"`
	FineTypeID 		int `gorm:"type:int(11)"`
	TransactionID	*int `gorm:"type:int(11)"`

	Book     book.Book
	Member   member.Member
	FineType FineType
	Transaction transaction.Transaction
}

type FineType struct {
	gorm.Model

	ID int `gorm:"type:int(11)"`
	Status string `gorm:"type:varchar(255)"`
	FineCost int `gorm:"type:bigint"`
}