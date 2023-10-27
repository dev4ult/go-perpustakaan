package transaction

import (
	"perpustakaan/features/member"

	"gorm.io/gorm"
)

type Transaction struct {
	gorm.Model

	ID int `gorm:"type:int(11)"`
	OrderID string `gorm:"type:varchar(100)"`
	Note string `gorm:"type:varchar(255)"`
	Status string `gorm:"type:varchar(255)"`
	PaymentURL string `gorm:"type:varchar(255)"`
	MemberID int `gorm:"type:int(11)"`
	
	Member member.Member
}

