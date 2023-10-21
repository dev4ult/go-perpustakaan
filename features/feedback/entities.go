package feedback

import (
	"perpustakaan/features/member"

	"gorm.io/gorm"
)

type Feedback struct {
	gorm.Model

	ID int `gorm:"type:int(11)"`
	Comment string `gorm:"type:text"`
	PriorityStatus string `gorm:"type:enum(high,medium,low)"`
	MemberID string `gorm:"type:int(11)"`

	Member member.Member
}

