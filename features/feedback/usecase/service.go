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

func (svc *service) FindAll(page, size int) []dtos.FeedbackWithReply {
	feedbacks := svc.model.Paginate(page, size)

	var feedbackWithReply []dtos.FeedbackWithReply 
	for _, feedback := range feedbacks {
		if feedback.Member == "" {
			feedback.Member = "Anonymous"
		}
		
		feedbackWithReply = append(feedbackWithReply, dtos.FeedbackWithReply{
			Member: feedback.Member,
			Comment: feedback.Comment,
			PriorityStatus: feedback.PriorityStatus,
			Reply: dtos.StaffReply{
				Staff: feedback.Staff,
				Comment: feedback.Reply,
			},
		})
	}

	return feedbackWithReply
}

func (svc *service) FindByID(feedbackID int) *dtos.FeedbackWithReply {
	feedback := svc.model.SelectByID(feedbackID)

	if feedback == nil {
		log.Error("Feedback Not Found!")
		return nil
	}

	return feedback
}

func (svc *service) Create(newFeedback dtos.InputFeedback) *dtos.FeedbackWithReply {
	feedback := feedback.Feedback{}
	
	if err := smapping.FillStruct(&feedback, smapping.MapFields(newFeedback)); err != nil {
		log.Error(err.Error())
		return nil
	}
	
	feedback.PriorityStatus = helpers.GetPrediction(newFeedback.Comment)

	resFeedback := svc.model.Insert(feedback)
	if resFeedback == nil {
		return nil
	}

	feedbackWithReply := dtos.FeedbackWithReply{}

	if err := smapping.FillStruct(&feedbackWithReply, smapping.MapFields(resFeedback)); err != nil {
		log.Error(err.Error())
		return nil
	}

	return &feedbackWithReply
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