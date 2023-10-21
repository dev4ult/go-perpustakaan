package usecase

import (
	"perpustakaan/features/feedback"
	"perpustakaan/features/feedback/dtos"

	"github.com/labstack/gommon/log"
	"github.com/mashingan/smapping"
)

type service struct {
	model feedback.Repository
}

func New(model feedback.Repository) feedback.Usecase {
	return &service {
		model: model,
	}
}

func (svc *service) FindAll(page, size int) []dtos.ResFeedback {
	var feedbacks []dtos.ResFeedback

	feedbacksEnt := svc.model.Paginate(page, size)

	for _, feedback := range feedbacksEnt {
		var data dtos.ResFeedback

		if err := smapping.FillStruct(&data, smapping.MapFields(feedback)); err != nil {
			log.Error(err.Error())
		} 
		
		feedbacks = append(feedbacks, data)
	}

	return feedbacks
}

func (svc *service) FindByID(feedbackID int) *dtos.ResFeedback {
	res := dtos.ResFeedback{}
	feedback := svc.model.SelectByID(feedbackID)

	if feedback == nil {
		return nil
	}

	err := smapping.FillStruct(&res, smapping.MapFields(feedback))
	if err != nil {
		log.Error(err)
		return nil
	}

	return &res
}

func (svc *service) Create(newFeedback dtos.InputFeedback) *dtos.ResFeedback {
	feedback := feedback.Feedback{}
	
	err := smapping.FillStruct(&feedback, smapping.MapFields(newFeedback))
	if err != nil {
		log.Error(err)
		return nil
	}

	feedbackID := svc.model.Insert(feedback)

	if feedbackID == -1 {
		return nil
	}

	resFeedback := dtos.ResFeedback{}
	errRes := smapping.FillStruct(&resFeedback, smapping.MapFields(newFeedback))
	if errRes != nil {
		log.Error(errRes)
		return nil
	}

	return &resFeedback
}

func (svc *service) Modify(feedbackData dtos.InputFeedback, feedbackID int) bool {
	newFeedback := feedback.Feedback{}

	err := smapping.FillStruct(&newFeedback, smapping.MapFields(feedbackData))
	if err != nil {
		log.Error(err)
		return false
	}

	newFeedback.ID = feedbackID
	rowsAffected := svc.model.Update(newFeedback)

	if rowsAffected <= 0 {
		log.Error("There is No Feedback Updated!")
		return false
	}
	
	return true
}

func (svc *service) Remove(feedbackID int) bool {
	rowsAffected := svc.model.DeleteByID(feedbackID)

	if rowsAffected <= 0 {
		log.Error("There is No Feedback Deleted!")
		return false
	}

	return true
}