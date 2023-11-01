package handler

import (
	"perpustakaan/helpers"

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
		input := ctx.Get("request").(*dtos.InputLogin)

		if (input.StaffID != "" && input.CredentialNumber != "") || (input.StaffID == "" && input.CredentialNumber == "") {
			return ctx.JSON(400, helpers.Response("Choose between `credential-number` and `staff-id` for Login. Do not send both!"))
		}

		isLibrarian := false
		credential := input.CredentialNumber

		if input.StaffID != "" {
			isLibrarian = true
			credential = input.StaffID
		}

		authorization, message := ctl.service.VerifyLogin(credential, input.Password, isLibrarian)

		if message != "" {
			return ctx.JSON(500, helpers.Response(message))
		}

		if authorization == nil {
			return ctx.JSON(404, helpers.Response("Your credential or password does not Match!"))
		}
		
		return ctx.JSON(200, helpers.Response("Success Login!", map[string]any {
			"data": authorization,
		}))
	}
}

func (ctl *controller) StaffRegistration() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		input := ctx.Get("request").(*dtos.InputStaffRegistration)

		librarian, _ := ctl.service.FindLibrarianByStaffID(input.StaffID)
		
		if librarian != nil {
			return ctx.JSON(409, helpers.Response("Staff ID already Registered!"))
		}

		resLibrarian, message := ctl.service.RegisterAStaff(*input)

		if message != "" {
			return ctx.JSON(500, helpers.Response(message))
		}

		return ctx.JSON(200, helpers.Response("test", map[string]any {
			"data": resLibrarian,
		}))
	}
}

func (ctl *controller) Refresh() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		input := ctx.Get("request").(*dtos.Authorization)
		userID := int(ctx.Get("user-id").(float64))
		role := ctx.Get("role").(string)

		accessTokenClaim := helpers.ExtractToken(input.AccessToken, false)
	
		if accessTokenClaim != nil {
			return ctx.JSON(500, helpers.Response("Access Token Still Valid"))
		}

		newToken, message := ctl.service.RefreshToken(userID, role)

		if message != "" {
			return ctx.JSON(500, helpers.Response(message))
		}

		return ctx.JSON(200, helpers.Response("Token Refreshed!", map[string]any {
			"data": newToken,
		}))
	}
}