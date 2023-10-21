package routes

import (
	"perpustakaan/features/feedback"

	"github.com/labstack/echo/v4"
)

func Feedbacks(e *echo.Echo, handler feedback.Handler) {
	feedbacks := e.Group("/feedbacks")

	feedbacks.GET("", handler.GetFeedbacks())
	feedbacks.POST("", handler.CreateFeedback())
	
	feedbacks.GET("/:id", handler.FeedbackDetails())
	feedbacks.PUT("/:id", handler.UpdateFeedback())
	feedbacks.DELETE("/:id", handler.DeleteFeedback())
}