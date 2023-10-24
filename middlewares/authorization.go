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
				return ctx.JSON(401, helpers.Response("There is no Authorization Given in the Header!"))
			}
			
			token := auth[len("Bearer "):]
			claims := helpers.ExtractToken(token)

			if role == "all" {
				return next(ctx)
			}
			
			if role != claims["role"] {
				return ctx.JSON(401, helpers.Response("This user can't access for this Endpoint!"))
			}
			
			return next(ctx)

		}
	}
}