package feedback

import (
	"perpustakaan/features/feedback/dtos"

	"github.com/labstack/echo/v4"
)

type Repository interface {
	Paginate(page int, size int, member string, priority string) ([]dtos.FeedbackJoinReply, error)
	Insert(newFeedback Feedback) (*dtos.ResFeedback, error)
	SelectByID(feedbackID int) (*dtos.FeedbackWithReply, error)
	InsertReplyForAFeedback(reply FeedbackReply) (*dtos.StaffReply, error)
	DeleteByID(feedbackID int) (int, error)
}

type Usecase interface {
	FindAll(page int, size int, member string, priority string) ([]dtos.FeedbackWithReply, string)
	FindByID(feedbackID int) (*dtos.FeedbackWithReply, string)
	Create(newFeedback dtos.InputFeedback) (*dtos.FeedbackWithReply, string)
	AddAReply(replyData dtos.InputReply, feedbackID int) (*dtos.StaffReply, string)
	Remove(feedbackID int) (bool, string)
}

type Handler interface {
	GetFeedbacks() echo.HandlerFunc
	FeedbackDetails() echo.HandlerFunc
	CreateFeedback() echo.HandlerFunc
	ReplyOnFeedback() echo.HandlerFunc
	DeleteFeedback() echo.HandlerFunc
}
