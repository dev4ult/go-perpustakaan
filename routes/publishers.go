package routes

import (
	"perpustakaan/features/publisher"
	"perpustakaan/features/publisher/dtos"
	m "perpustakaan/middlewares"

	"github.com/labstack/echo/v4"
)

func Publishers(e *echo.Echo, handler publisher.Handler) {
	publishers := e.Group("/publishers")
	publishers.Use(m.Authorization("librarian"))

	publishers.GET("", handler.GetPublishers())
	publishers.POST("", handler.CreatePublisher(), m.RequestValidation(&dtos.InputPublisher{}))
	
	publishers.GET("/:id", handler.PublisherDetails())
	publishers.PUT("/:id", handler.UpdatePublisher(), m.RequestValidation(&dtos.InputPublisher{}))
	publishers.DELETE("/:id", handler.DeletePublisher())
}