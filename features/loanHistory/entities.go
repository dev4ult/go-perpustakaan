package loanHistory

import (
	"perpustakaan/features/book"
	"perpustakaan/features/member"

	"gorm.io/gorm"
)

type LoanHistory struct {
	gorm.Model

	ID         int `gorm:"type:int(11)"`
	BookID     int `gorm:"type:int(11)"`
	MemberID   int `gorm:"type:int(11)"`
	FineTypeID int `gorm:"type:int(11)"`

	Book     book.Book
	Member   member.Member
	FineType book.FineType
}