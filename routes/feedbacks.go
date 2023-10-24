package routes

import (
	"perpustakaan/features/feedback"
	"perpustakaan/features/feedback/dtos"
	m "perpustakaan/middlewares"

	"github.com/labstack/echo/v4"
)

func Feedbacks(e *echo.Echo, handler feedback.Handler) {
	feedbacks := e.Group("/feedbacks")

	feedbacks.GET("", handler.GetFeedbacks())
	feedbacks.POST("", handler.CreateFeedback(), m.Authorization("member"), m.RequestValidation(dtos.InputFeedback{}))
	
	feedbacks.GET("/:id", handler.FeedbackDetails())
	feedbacks.POST("/:id", handler.ReplyOnFeedback(), m.Authorization("librarian"), m.RequestValidation(dtos.InputReply{}))
	feedbacks.DELETE("/:id", handler.DeleteFeedback(), m.Authorization("librarian"))
}