package routes

import (
	"perpustakaan/features/publisher"

	"github.com/labstack/echo/v4"
)

func Publishers(e *echo.Echo, handler publisher.Handler) {
	publishers := e.Group("/publishers")

	publishers.GET("", handler.GetPublishers())
	publishers.POST("", handler.CreatePublisher())
	
	publishers.GET("/:id", handler.PublisherDetails())
	publishers.PUT("/:id", handler.UpdatePublisher())
	publishers.DELETE("/:id", handler.DeletePublisher())
}