package feedback

import (
	"perpustakaan/features/feedback/dtos"

	"github.com/labstack/echo/v4"
)

type Repository interface {
	Paginate(page, size int) []dtos.ResFeedback
	Insert(newFeedback Feedback) *dtos.ResFeedback
	SelectByID(feedbackID int) *dtos.ResFeedback
	Update(feedback Feedback) int64
	DeleteByID(feedbackID int) int64
}

type Usecase interface {
	FindAll(page, size int) []dtos.ResFeedback
	FindByID(feedbackID int) *dtos.ResFeedback
	Create(newFeedback dtos.InputFeedback) *dtos.ResFeedback
	Modify(feedbackData dtos.InputFeedback, feedbackID int) bool
	Remove(feedbackID int) bool
}

type Handler interface {
	GetFeedbacks() echo.HandlerFunc
	FeedbackDetails() echo.HandlerFunc
	CreateFeedback() echo.HandlerFunc
	UpdateFeedback() echo.HandlerFunc
	DeleteFeedback() echo.HandlerFunc
}
