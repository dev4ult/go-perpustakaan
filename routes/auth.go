package routes

import (
	"perpustakaan/features/auth"
	"perpustakaan/features/auth/dtos"
	m "perpustakaan/middlewares"

	"github.com/labstack/echo/v4"
)

func Auths(e *echo.Echo, handler auth.Handler) {
	e.POST("/login", handler.Login(), m.RequestValidation(dtos.InputLogin{}))
	e.POST("/refresh", handler.Refresh(), m.Authorization("all", true), m.RequestValidation(dtos.Authorization{}))
	e.POST("/staff", handler.StaffRegistration(), m.RequestValidation(dtos.InputStaffRegistration{}))
}