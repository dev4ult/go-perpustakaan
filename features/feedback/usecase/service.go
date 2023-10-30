package usecase

import (
	"perpustakaan/features/feedback"
	"perpustakaan/features/feedback/dtos"

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

func (svc *service) FindAll(page int, size int, member string, priority string) ([]dtos.FeedbackWithReply, string) {
	feedbacks, err := svc.model.Paginate(page, size, member, priority)

	if err != nil {
		return nil, err.Error()
	}

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

	return feedbackWithReply, ""
}

func (svc *service) FindByID(feedbackID int) (*dtos.FeedbackWithReply, string) {
	feedback, err := svc.model.SelectByID(feedbackID)

	if err != nil {
		return nil, err.Error()
	}

	return feedback, ""
}

func (svc *service) Create(newFeedback dtos.InputFeedback) (*dtos.FeedbackWithReply, string) {
	var feedback feedback.Feedback
	
	if err := smapping.FillStruct(&feedback, smapping.MapFields(newFeedback)); err != nil {
		return nil, err.Error()
	}
	
	res, err := svc.model.Insert(feedback)
	if err != nil {
		return nil, err.Error()
	}

	var feedbackWithReply dtos.FeedbackWithReply

	if err := smapping.FillStruct(&feedbackWithReply, smapping.MapFields(res)); err != nil {
		return nil, err.Error()
	}

	return &feedbackWithReply, ""
}

func (svc *service) AddAReply(replyData dtos.InputReply, feedbackID int) (*dtos.StaffReply, string) {
	var feedbackReply feedback.FeedbackReply
	
	if err := smapping.FillStruct(&feedbackReply, smapping.MapFields(replyData)); err != nil {
		return nil, err.Error()
	}

	feedbackReply.FeedbackID = feedbackID
	feedbackReply.LibrarianID = replyData.StaffID
	staffReply, err := svc.model.InsertReplyForAFeedback(feedbackReply)

	if err != nil {
		return nil, err.Error()
	}
	
	return staffReply, ""
}

func (svc *service) Remove(feedbackID int) (bool, string) {
	_, err := svc.model.DeleteByID(feedbackID)

	if err != nil {
		return false, err.Error()
	}

	return true, ""
}