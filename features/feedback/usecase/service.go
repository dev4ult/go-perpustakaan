package usecase

import (
	"perpustakaan/features/feedback"
	"perpustakaan/features/feedback/dtos"
	"perpustakaan/helpers"

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
	feedbacks := svc.model.Paginate(page, size)

	return feedbacks
}

func (svc *service) FindByID(feedbackID int) *dtos.ResFeedback {
	feedback := svc.model.SelectByID(feedbackID)

	if feedback == nil {
		log.Error("Feedback Not Found!")
		return nil
	}

	return feedback
}

func (svc *service) Create(newFeedback dtos.InputFeedback) *dtos.ResFeedback {
	feedback := feedback.Feedback{}
	
	err := smapping.FillStruct(&feedback, smapping.MapFields(newFeedback))
	if err != nil {
		log.Error(err)
		return nil
	}
	
	feedback.PriorityStatus = helpers.GetPrediction(newFeedback.Comment)

	fb := svc.model.Insert(feedback)

	if fb == nil {
		return nil
	}

	return fb
}

func (svc *service) AddAReply(replyData dtos.InputReply, feedbackID int) *dtos.FeedbackWithReply {
	feedbackReply := feedback.FeedbackReply{}
	
	if err := smapping.FillStruct(&feedbackReply, smapping.MapFields(replyData)); err != nil {
		log.Error(err)
		return nil
	}

	feedbackReply.FeedbackID = feedbackID
	feedbackReply.LibrarianID = replyData.StaffID
	staffReply := svc.model.InsertReplyForAFeedback(feedbackReply)

	if staffReply == nil {
		log.Error("There is No Reply Added to The Feedback!")
		return nil
	}

	feedbackWithReply := dtos.FeedbackWithReply{}

	if err := smapping.FillStruct(&feedbackWithReply, smapping.MapFields(replyData)); err != nil {
		log.Error(err)
		return nil
	}

	feedbackWithReply.Reply.Staff = staffReply.Staff
	feedbackWithReply.Reply.Comment = staffReply.Comment
	
	return &feedbackWithReply
}

func (svc *service) Remove(feedbackID int) bool {
	rowsAffected := svc.model.DeleteByID(feedbackID)

	if rowsAffected <= 0 {
		log.Error("There is No Feedback Deleted!")
		return false
	}

	return true
}