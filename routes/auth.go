package routes

import (
	"perpustakaan/config"
	"perpustakaan/features/auth"
	"perpustakaan/features/auth/dtos"
	m "perpustakaan/middlewares"

	"github.com/labstack/echo/v4"
)

func Auths(e *echo.Echo, handler auth.Handler, cfg config.ServerConfig) {
	e.POST("/login", handler.Login(), m.RequestValidation(dtos.InputLogin{}))
	e.POST("/refresh", handler.Refresh(), m.Authorization("all"), m.RequestValidation(dtos.Authorization{}))
	e.POST("/staff", handler.StaffRegistration(), m.RequestValidation(dtos.InputStaffRegistration{}))
}