package middlewares

import (
	"perpustakaan/helpers"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

func RequestValidation(request any) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {
			if err := ctx.Bind(request); err != nil {
				errorMap := helpers.ParseError(err.Error())
				return ctx.JSON(errorMap["code"].(int), helpers.Response(errorMap["message"].(string)))
			}

			var validate = validator.New()

			if err := validate.Struct(request); err != nil {
				var errMap = []map[string]string{} 
				for _, err := range err.(validator.ValidationErrors) {
					errMap = append(errMap, map[string]string {
						"field": err.Field(),
						"tag": err.ActualTag(),
					})
				}

				return ctx.JSON(400, helpers.Response("Missing Data Required!", map[string]any {
					"errors": errMap,
				}))
			}

			ctx.Set("request", request)

			return next(ctx)
		}
	}
}