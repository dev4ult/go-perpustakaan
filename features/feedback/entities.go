package feedback

import (
	"perpustakaan/features/auth"
	"perpustakaan/features/member"

	"gorm.io/gorm"
)

type Feedback struct {
	gorm.Model

	ID int `gorm:"type:int(11)"`
	Comment string `gorm:"type:text"`
	PriorityStatus string `gorm:"type:enum('high','medium','low')"`
	MemberID *int `gorm:"type:int(11)"`

	Member member.Member
}

type FeedbackReply struct {
	gorm.Model

	ID int `gorm:"type:int(11)"`
	Comment string `gorm:"type:text"`
	LibrarianID int `gorm:"type:int(11)"`
	FeedbackID int `gorm:"type:int(11)"`

	Feedback Feedback
	Librarian auth.Librarian
}
