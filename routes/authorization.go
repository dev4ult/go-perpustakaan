package routes

import (
	"perpustakaan/features/auth"

	"github.com/labstack/echo/v4"
)

func Auths(e *echo.Echo, handler auth.Handler) {
	auths := e.Group("/auths")

	auths.GET("", handler.GetAuths())
	auths.POST("", handler.CreateAuth())
	
	auths.GET("/:id", handler.AuthDetails())
	auths.PUT("/:id", handler.UpdateAuth())
	auths.DELETE("/:id", handler.DeleteAuth())
}