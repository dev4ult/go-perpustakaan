package repository

import (
	"perpustakaan/features/feedback"
	"perpustakaan/features/feedback/dtos"
	"perpustakaan/features/member"

	"github.com/labstack/gommon/log"
	"github.com/mashingan/smapping"
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

func (mdl *model) Paginate(page, size int) []dtos.ResFeedback {
	var feedbacks []dtos.ResFeedback

	offset := (page - 1) * size

	result := mdl.db.Table("feedbacks").Select("feedbacks.*, members.full_name as user").Joins("LEFT JOIN members ON members.id = feedbacks.member_id").Offset(offset).Limit(size).Find(&feedbacks)
	
	if result.Error != nil {
		log.Error(result.Error)
		return nil
	}

	return feedbacks
}

func (mdl *model) Insert(newFeedback feedback.Feedback) *dtos.ResFeedback {
	result := mdl.db.Create(&newFeedback)

	if result.Error != nil {
		log.Error(result.Error)
		return nil
	}
	
	var feedback dtos.ResFeedback
	err := smapping.FillStruct(&feedback, smapping.MapFields(newFeedback))
	if err != nil {
		log.Error(err.Error())
		return nil
	}

	feedback.User = "Anonymous"
	if newFeedback.MemberID != nil {
		var member member.Member
		result = mdl.db.Table("members").Where("id = ?", newFeedback.MemberID).First(&member)

		if result.Error == nil {
			feedback.User = member.FullName
		} else {
			log.Error(result.Error.Error())
		}

	}

	return &feedback
}

func (mdl *model) SelectByID(feedbackID int) *dtos.ResFeedback {
	var feedback dtos.ResFeedback
	result := mdl.db.Table("feedbacks").Select("feedbacks.*, members.full_name as user").Joins("LEFT JOIN members ON members.id = feedbacks.member_id").First(&feedback)

	if result.Error != nil {
		log.Error(result.Error)
		return nil
	}

	return &feedback
}


func (mdl *model) Update(feedback feedback.Feedback) int64 {
	result := mdl.db.Save(&feedback)

	if result.Error != nil {
		log.Error(result.Error)
	}

	return result.RowsAffected
}

func (mdl *model) DeleteByID(feedbackID int) int64 {
	result := mdl.db.Delete(&feedback.Feedback{}, feedbackID)
	
	if result.Error != nil {
		log.Error(result.Error)
		return 0
	}

	return result.RowsAffected
}