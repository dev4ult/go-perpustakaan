package middlewares

import (
	"perpustakaan/helpers"

	"github.com/labstack/echo/v4"
)

func Authorization(role string, options ...bool) echo.MiddlewareFunc {
	return func (next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {
			auth := ctx.Request().Header.Get("Authorization")

			if auth == "" {
				return ctx.JSON(401, helpers.Response("There is no Authorization Given in the Header!"))
			}

			isRefreshToken := false

			if len(options) > 1 {
				return ctx.JSON(500, helpers.Response("Too many argument for options in Authorization!"))
			} else if len(options) == 1 {
				isRefreshToken = options[0]
			}
			
			token := auth[len("Bearer "):]
			claims := helpers.ExtractToken(token, isRefreshToken)

			if claims == nil {
				return ctx.JSON(401, helpers.Response("Token is invalid or expired!"))
			}

			ctx.Set("user-id", claims["user-id"])
			ctx.Set("role", claims["role"])

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