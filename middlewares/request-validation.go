package middlewares

import (
	"fmt"
	"perpustakaan/helpers"
	"reflect"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

func RequestValidation(request any) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {
			requestType := reflect.TypeOf(request)
			newRequest := reflect.New(requestType).Interface()
			validate := validator.New()

			if err := ctx.Bind(&newRequest); err != nil {
				return ctx.JSON(500, helpers.Response(fmt.Sprintf("Error Binding Request : %s", err.Error())))
			}

			if err := validate.Struct(newRequest); err != nil {
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
			
			ctx.Set("request", newRequest)

			return next(ctx)
		}
	}
}