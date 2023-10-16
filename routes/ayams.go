package routes

import (
	"perpustakaan/features/ayam"

	"github.com/labstack/echo/v4"
)

func Ayams(e *echo.Echo, handler ayam.Handler) {
	ayams := e.Group("/ayams")

	ayams.GET("", handler.GetAyams())
	ayams.POST("", handler.CreateAyam())
	
	ayams.GET("/:id", handler.AyamDetails())
	ayams.PUT("/:id", handler.UpdateAyam())
	ayams.DELETE("/:id", handler.DeleteAyam())
}