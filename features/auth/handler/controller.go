package handler

import (
	"perpustakaan/helpers"
	helper "perpustakaan/helpers"

	"perpustakaan/features/auth"
	"perpustakaan/features/auth/dtos"

	"github.com/labstack/echo/v4"
)

type controller struct {
	service auth.Usecase
}

func New(service auth.Usecase) auth.Handler {
	return &controller {
		service: service,
	}
}


func (ctl *controller) Login() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		var input dtos.InputLogin

		ctx.Bind(&input)

		if err := helpers.ValidateRequest(input); err != nil {
			errMap := helper.ErrorMapValidation(err)
			return ctx.JSON(400, helper.Response("Login Credential Can Not be Empty!", map[string]any {
				"error": errMap,
			}))
		}

		if (input.StaffID != "" && input.CredentialNumber != "") || (input.StaffID == "" && input.CredentialNumber == "") {
			return ctx.JSON(400, helper.Response("Choose between `credential-number` and `staff-id` for Login. Do not send both!"))
		}

		isLibrarian := false

		if input.StaffID != "" {
			isLibrarian = true
		}

		authorization := ctl.service.VerifyLogin(input.CredentialNumber, input.Password, isLibrarian)

		if authorization == nil {
			return ctx.JSON(404, helper.Response("Your credential or password does not Match!"))
		}
		
		return ctx.JSON(200, helper.Response("Success Login!", map[string]any {
			"data": authorization,
		}))
	}
}

func (ctl *controller) Refresh() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		refreshToken := ctx.Request().Header.Get("Authorization")
		var	authorization dtos.Authorization
		
		ctx.Bind(&authorization)

		if err := helpers.ValidateRequest(authorization); err != nil {
			return ctx.JSON(400, helper.Response("Missing Access Token!"))
		}

		refreshToken = refreshToken[len("Bearer "):]

		newToken := ctl.service.RefreshToken(authorization.AccessToken, refreshToken)

		if newToken == nil {
			return ctx.JSON(500, helper.Response("Something Went Wrong!"))
		}

		return ctx.JSON(200, helper.Response("Token Refreshed!", map[string]any {
			"data": newToken,
		}))
	}
}

// func (ctl *controller) GetAuths() echo.HandlerFunc {
// 	return func (ctx echo.Context) error  {
// 		pagination := dtos.Pagination{}
// 		ctx.Bind(&pagination)
		
// 		page := pagination.Page
// 		size := pagination.Size

// 		if page <= 0 || size <= 0 {
// 			return ctx.JSON(400, helper.Response("Please provide query `page` and `size` in number!"))
// 		}

// 		auths := ctl.service.FindAll(page, size)

// 		if auths == nil {
// 			return ctx.JSON(404, helper.Response("There is No Auths!"))
// 		}

// 		return ctx.JSON(200, helper.Response("Success!", map[string]any {
// 			"data": auths,
// 		}))
// 	}
// }


// func (ctl *controller) AuthDetails() echo.HandlerFunc {
// 	return func (ctx echo.Context) error  {
// 		authID, err := strconv.Atoi(ctx.Param("id"))

// 		if err != nil {
// 			return ctx.JSON(400, helper.Response(err.Error()))
// 		}

// 		auth := ctl.service.FindByID(authID)

// 		if auth == nil {
// 			return ctx.JSON(404, helper.Response("Auth Not Found!"))
// 		}

// 		return ctx.JSON(200, helper.Response("Success!", map[string]any {
// 			"data": auth,
// 		}))
// 	}
// }

// func (ctl *controller) CreateAuth() echo.HandlerFunc {
// 	return func (ctx echo.Context) error  {
// 		input := dtos.InputAuth{}

// 		ctx.Bind(&input)

// 		validate = validator.New(validator.WithRequiredStructEnabled())

// 		err := validate.Struct(input)

// 		if err != nil {
// 			errMap := helpers.ErrorMapValidation(err)
// 			return ctx.JSON(400, helper.Response("Bad Request!", map[string]any {
// 				"error": errMap,
// 			}))
// 		}

// 		auth := ctl.service.Create(input)

// 		if auth == nil {
// 			return ctx.JSON(500, helper.Response("Something went Wrong!", nil))
// 		}

// 		return ctx.JSON(200, helper.Response("Success!", map[string]any {
// 			"data": auth,
// 		}))
// 	}
// }

// func (ctl *controller) UpdateAuth() echo.HandlerFunc {
// 	return func (ctx echo.Context) error {
// 		input := dtos.InputAuth{}

// 		authID, errParam := strconv.Atoi(ctx.Param("id"))

// 		if errParam != nil {
// 			return ctx.JSON(400, helper.Response(errParam.Error()))
// 		}

// 		auth := ctl.service.FindByID(authID)

// 		if auth == nil {
// 			return ctx.JSON(404, helper.Response("Auth Not Found!"))
// 		}
		
// 		ctx.Bind(&input)

// 		validate = validator.New(validator.WithRequiredStructEnabled())
// 		err := validate.Struct(input)

// 		if err != nil {
// 			errMap := helpers.ErrorMapValidation(err)
// 			return ctx.JSON(400, helper.Response("Bad Request!", map[string]any {
// 				"error": errMap,
// 			}))
// 		}

// 		update := ctl.service.Modify(input, authID)

// 		if !update {
// 			return ctx.JSON(500, helper.Response("Something Went Wrong!"))
// 		}

// 		return ctx.JSON(200, helper.Response("Auth Success Updated!"))
// 	}
// }

// func (ctl *controller) DeleteAuth() echo.HandlerFunc {
// 	return func (ctx echo.Context) error  {
// 		authID, err := strconv.Atoi(ctx.Param("id"))

// 		if err != nil {
// 			return ctx.JSON(400, helper.Response(err.Error()))
// 		}

// 		auth := ctl.service.FindByID(authID)

// 		if auth == nil {
// 			return ctx.JSON(404, helper.Response("Auth Not Found!"))
// 		}

// 		delete := ctl.service.Remove(authID)

// 		if !delete {
// 			return ctx.JSON(500, helper.Response("Something Went Wrong!"))
// 		}

// 		return ctx.JSON(200, helper.Response("Auth Success Deleted!", nil))
// 	}
// }
