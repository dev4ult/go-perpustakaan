package feedback

import (
	"perpustakaan/features/feedback/dtos"

	"github.com/labstack/echo/v4"
)

type Repository interface {
	Paginate(page, size int) []dtos.ResFeedback
	Insert(newFeedback Feedback) *dtos.ResFeedback
	SelectByID(feedbackID int) *dtos.ResFeedback
	InsertReplyForAFeedback(reply FeedbackReply) *dtos.StaffReply
	DeleteByID(feedbackID int) int64
}

type Usecase interface {
	FindAll(page, size int) []dtos.ResFeedback
	FindByID(feedbackID int) *dtos.ResFeedback
	Create(newFeedback dtos.InputFeedback) *dtos.ResFeedback
	AddAReply(replyData dtos.InputReply, feedbackID int) *dtos.FeedbackWithReply
	Remove(feedbackID int) bool
}

type Handler interface {
	GetFeedbacks() echo.HandlerFunc
	FeedbackDetails() echo.HandlerFunc
	CreateFeedback() echo.HandlerFunc
	ReplyOnFeedback() echo.HandlerFunc
	DeleteFeedback() echo.HandlerFunc
}
