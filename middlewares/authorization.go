package middlewares

import (
	"perpustakaan/helpers"

	"github.com/labstack/echo/v4"
)

func Authorization(role string) echo.MiddlewareFunc {
	return func (next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {
			auth := ctx.Request().Header.Get("Authorization")

			if auth == "" {
				return ctx.JSON(401, helpers.Response("There is no Authorization Token Given!"))
			}
			
			// token := auth[len("Bearer "):]

			// claims := helpers.ExtractToken(token)


			return next(ctx)
		}
	}
}