package repository

import (
	"perpustakaan/features/feedback"

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

func (mdl *model) Paginate(page, size int) []feedback.Feedback {
	var feedbacks []feedback.Feedback

	offset := (page - 1) * size

	result := mdl.db.Offset(offset).Limit(size).Find(&feedbacks)
	
	if result.Error != nil {
		log.Error(result.Error)
		return nil
	}

	return feedbacks
}

func (mdl *model) Insert(newFeedback feedback.Feedback) int64 {
	result := mdl.db.Create(&newFeedback)

	if result.Error != nil {
		log.Error(result.Error)
		return -1
	}

	return int64(newFeedback.ID)
}

func (mdl *model) SelectByID(feedbackID int) *feedback.Feedback {
	var feedback feedback.Feedback
	result := mdl.db.First(&feedback, feedbackID)

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