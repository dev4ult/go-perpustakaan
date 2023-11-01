package repository

import (
	"perpustakaan/features/auth"
	"perpustakaan/features/feedback"
	"perpustakaan/features/feedback/dtos"
	"perpustakaan/features/member"

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

func (mdl *model) Paginate(page int, size int, member string, priority string) ([]dtos.FeedbackJoinReply, error) {
	var feedbacks []dtos.FeedbackJoinReply

	offset := (page - 1) * size
	memberName := "%" + member + "%"
	priorityStatus := "%" + priority + "%"
	
	if err := mdl.db.Table("feedbacks").
	Select("feedbacks.comment, feedbacks.priority_status, members.full_name as member, feedback_replies.comment as reply, librarians.full_name as staff").
		Joins("LEFT JOIN members ON members.id = feedbacks.member_id").
		Joins("LEFT JOIN feedback_replies ON feedbacks.id = feedback_replies.feedback_id").
		Joins("LEFT JOIN librarians ON librarians.id = feedback_replies.librarian_id").
		Where("members.full_name IS NULL OR members.full_name LIKE ?", memberName).
		Where("feedbacks.priority_status LIKE ?", priorityStatus).
		Offset(offset).Limit(size).Find(&feedbacks).Error; err != nil {
		return nil, err
	}

	return feedbacks, nil
}

func (mdl *model) Insert(newFeedback feedback.Feedback) (*dtos.ResFeedback, error) {
	if err := mdl.db.Create(&newFeedback).Error; err != nil {
		return nil, err
	}
	
	var feedback = dtos.ResFeedback{
		Comment: newFeedback.Comment,
		PriorityStatus: newFeedback.PriorityStatus,
	}

	feedback.Member = "Anonymous"
	if newFeedback.MemberID != nil {
		var member member.Member

		if err := mdl.db.Table("members").Where("id = ?", newFeedback.MemberID).First(&member).Error; err == nil {
			feedback.Member = member.FullName
		} else {
			return nil, err
		}

	}

	return &feedback, nil
}

func (mdl *model) SelectByID(feedbackID int) (*dtos.FeedbackWithReply, error) {
	var fb feedback.Feedback
	
	if err := mdl.db.First(&fb, feedbackID).Error; err != nil {
		return nil, err
	}

	var feedbackWithReply = dtos.FeedbackWithReply{
		Comment: fb.Comment,
		PriorityStatus: fb.PriorityStatus,
	}

	feedbackWithReply.Member = "Anonymous"
	if fb.MemberID != nil {
		var member member.Member

		if err := mdl.db.Table("members").Where("id = ?", fb.MemberID).First(&member).Error; err == nil {
			feedbackWithReply.Member = member.FullName
		} else {
			return nil, err
		}
	}

	var reply feedback.FeedbackReply
	
	if result := mdl.db.Table("feedback_replies").Where("feedback_id = ?", feedbackID).First(&reply); result.Error == nil {
		var librarian auth.Librarian
		if err := mdl.db.Table("librarians").Where("id = ?", reply.LibrarianID).First(&librarian).Error; err == nil {
			feedbackWithReply.Reply = dtos.StaffReply{
				Staff: librarian.FullName,
				Comment: reply.Comment,
			}
		}
	}

	return &feedbackWithReply, nil
}

func (mdl *model) InsertReplyForAFeedback(reply feedback.FeedbackReply) (*dtos.StaffReply, error) {
	if err := mdl.db.Create(&reply).Error; err != nil {
		return nil, err
	}

	var staff auth.Librarian
	if err := mdl.db.First(&staff, reply.LibrarianID).Error; err != nil {
		return nil, err
	}

	return &dtos.StaffReply{
		Staff: staff.FullName,
		Comment: reply.Comment,
	}, nil
}

func (mdl *model) DeleteByID(feedbackID int) (int, error) {
	result := mdl.db.Delete(&feedback.Feedback{}, feedbackID)
	
	if result.Error != nil {
		return 0, result.Error
	}

	return int(result.RowsAffected), nil
}