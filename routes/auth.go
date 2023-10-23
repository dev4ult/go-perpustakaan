package routes

import (
	"perpustakaan/config"
	"perpustakaan/features/auth"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func Auths(e *echo.Echo, handler auth.Handler, cfg config.ServerConfig) {
	e.POST("/login", handler.Login())
	e.POST("/refresh", handler.Refresh(), echojwt.JWT([]byte(cfg.REFRESH_KEY)))
	e.POST("/staff", handler.StaffRegistration())
}