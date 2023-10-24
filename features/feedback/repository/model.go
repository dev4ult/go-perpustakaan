package repository

import (
	"perpustakaan/features/auth"
	"perpustakaan/features/feedback"
	"perpustakaan/features/feedback/dtos"
	"perpustakaan/features/member"

	"github.com/labstack/gommon/log"
	"gorm.io/gorm"
)

type model struct {
	db *gorm.DB
}

func New(db *gorm.DB) feedback.Repository {
	return &model {
		db: db,
	}
}

func (mdl *model) Paginate(page, size int) []dtos.FeedbackJoinReply {
	var feedbacks []dtos.FeedbackJoinReply

	offset := (page - 1) * size

	result := mdl.db.Table("feedbacks").
	Select("feedbacks.comment, feedbacks.priority_status, members.full_name as member, feedback_replies.comment as reply, librarians.full_name as staff").
		Joins("LEFT JOIN members ON members.id = feedbacks.member_id").
		Joins("LEFT JOIN feedback_replies ON feedbacks.id = feedback_replies.feedback_id").
		Joins("LEFT JOIN librarians ON librarians.id = feedback_replies.librarian_id").
		Offset(offset).Limit(size).Find(&feedbacks)
	
	if result.Error != nil {
		log.Error(result.Error)
		return nil
	}

	return feedbacks
}

func (mdl *model) Insert(newFeedback feedback.Feedback) *dtos.ResFeedback {
	if result := mdl.db.Create(&newFeedback); result.Error != nil {
		log.Error(result.Error)
		return nil
	}
	
	var feedback = dtos.ResFeedback{
		Comment: newFeedback.Comment,
		PriorityStatus: newFeedback.PriorityStatus,
	}

	feedback.Member = "Anonymous"
	if newFeedback.MemberID != nil {
		var member member.Member

		if result := mdl.db.Table("members").Where("id = ?", newFeedback.MemberID).First(&member); result.Error == nil {
			feedback.Member = member.FullName
		} else {
			log.Error(result.Error.Error())
		}

	}

	return &feedback
}

func (mdl *model) SelectByID(feedbackID int) *dtos.FeedbackWithReply {
	var fb feedback.Feedback
	
	if result := mdl.db.First(&fb, feedbackID); result.Error != nil {
		log.Error(result.Error)
		return nil
	}

	var feedbackWithReply = dtos.FeedbackWithReply{
		Comment: fb.Comment,
		PriorityStatus: fb.PriorityStatus,
	}

	feedbackWithReply.Member = "Anonymous"
	
	if fb.MemberID != nil {
		var member member.Member

		if result := mdl.db.Table("members").Where("id = ?", fb.MemberID).First(&member); result.Error == nil {
			feedbackWithReply.Member = member.FullName
		} else {
			log.Error(result.Error.Error())
		}
	}

	var reply feedback.FeedbackReply
	
	if result := mdl.db.Table("feedback_replies").Where("feedback_id = ?", feedbackID).First(&reply); result.Error == nil {
		var librarian auth.Librarian
		if result := mdl.db.Table("librarians").Where("id = ?", reply.LibrarianID).First(&librarian); result.Error == nil {
			feedbackWithReply.Reply = dtos.StaffReply{
				Staff: librarian.FullName,
				Comment: reply.Comment,
			}
		}
	}

	return &feedbackWithReply
}

func (mdl *model) InsertReplyForAFeedback(reply feedback.FeedbackReply) *dtos.StaffReply {
	if result := mdl.db.Create(&reply); result.Error != nil {
		log.Error(result.Error)
	}

	var staff auth.Librarian
	if result := mdl.db.First(&staff, reply.LibrarianID); result.Error != nil {
		log.Error(result.Error)
	}

	return &dtos.StaffReply{
		Staff: staff.FullName,
		Comment: reply.Comment,
	} 
}

func (mdl *model) DeleteByID(feedbackID int) int64 {
	result := mdl.db.Delete(&feedback.Feedback{}, feedbackID)
	
	if result.Error != nil {
		log.Error(result.Error)
		return 0
	}

	return result.RowsAffected
}