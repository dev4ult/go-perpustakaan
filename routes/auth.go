package routes

import (
	"perpustakaan/config"
	"perpustakaan/features/auth"
	"perpustakaan/features/auth/dtos"
	m "perpustakaan/middlewares"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func Auths(e *echo.Echo, handler auth.Handler, cfg config.ServerConfig) {
	e.POST("/login", handler.Login(), m.RequestValidation(dtos.InputLogin{}))
	e.POST("/refresh", handler.Refresh(), echojwt.JWT([]byte(cfg.REFRESH_KEY)))
	e.POST("/staff", handler.StaffRegistration(), m.RequestValidation(dtos.InputStaffRegistration{}))
}