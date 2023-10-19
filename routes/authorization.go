package routes

import (
	"perpustakaan/features/auth"

	"github.com/labstack/echo/v4"
)

func Auths(e *echo.Echo, handler auth.Handler) {
	e.POST("/login", handler.Login())
}