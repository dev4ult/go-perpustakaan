package handler

import (
	"perpustakaan/helpers"
	helper "perpustakaan/helpers"
	"strconv"

	"perpustakaan/features/feedback"
	"perpustakaan/features/feedback/dtos"

	"github.com/labstack/echo/v4"
)

type controller struct {
	service feedback.Usecase
}

func New(service feedback.Usecase) feedback.Handler {
	return &controller {
		service: service,
	}
}

func (ctl *controller) GetFeedbacks() echo.HandlerFunc {
	return func (ctx echo.Context) error  {
		pagination := dtos.Pagination{}
		ctx.Bind(&pagination)
		
		page := pagination.Page
		size := pagination.Size

		if page <= 0 || size <= 0 {
			return ctx.JSON(400, helper.Response("Param must be provided in number!"))
		}

		feedbacks := ctl.service.FindAll(page, size)

		if feedbacks == nil {
			return ctx.JSON(404, helper.Response("There is No Feedbacks!"))
		}

		return ctx.JSON(200, helper.Response("Success!", map[string]any {
			"data": feedbacks,
		}))
	}
}

func (ctl *controller) FeedbackDetails() echo.HandlerFunc {
	return func (ctx echo.Context) error  {
		feedbackID, err := strconv.Atoi(ctx.Param("id"))

		if err != nil {
			return ctx.JSON(400, helper.Response("Param must be provided in number!"))
		}

		feedback := ctl.service.FindByID(feedbackID)

		if feedback == nil {
			return ctx.JSON(404, helper.Response("Feedback Not Found!"))
		}

		return ctx.JSON(200, helper.Response("Success!", map[string]any {
			"data": feedback,
		}))
	}
}

func (ctl *controller) CreateFeedback() echo.HandlerFunc {
	return func (ctx echo.Context) error  {
		authorization := ctx.Request().Header.Get("Authorization")
		input := dtos.InputFeedback{}

		ctx.Bind(&input)

		if err := helpers.ValidateRequest(input); err != nil {
			errMap := helpers.ErrorMapValidation(err)
			return ctx.JSON(400, helper.Response("Missing Data Required!", map[string]any {
				"errors": errMap,
			}))
		}

		if authorization != "" {
			token := authorization[len("Bearer "):]
			claims := helpers.ExtractToken(token)

			if claims == nil {
				return ctx.JSON(401, helpers.Response("Invalid Token Given!"))
			}

			userID := int(claims["user-id"].(float64))
			if role := claims["role"]; role == "librarian" {
				return ctx.JSON(401, helpers.Response("Member Only Authorization!"))
			}

			input.MemberID = &userID 
		}

		feedback := ctl.service.Create(input)

		if feedback == nil {
			return ctx.JSON(500, helper.Response("Something Went Wrong!", nil))
		}

		return ctx.JSON(200, helper.Response("Success!", map[string]any {
			"data": feedback,
		}))
	}
}

func (ctl *controller) ReplyOnFeedback() echo.HandlerFunc {
	return func (ctx echo.Context) error {
		authorization := ctx.Request().Header.Get("Authorization")

		input := dtos.InputReply{}

		feedbackID, err := strconv.Atoi(ctx.Param("id"))

		if err != nil {
			return ctx.JSON(400, helper.Response(err.Error()))
		}

		feedback := ctl.service.FindByID(feedbackID)

		if feedback == nil {
			return ctx.JSON(404, helper.Response("Feedback Not Found!"))
		}
		
		ctx.Bind(&input)
		
		if err := helpers.ValidateRequest(input); err != nil {
			errMap := helpers.ErrorMapValidation(err)
			return ctx.JSON(400, helper.Response("Missing Data Required!", map[string]any {
				"error": errMap,
			}))
		}

		if authorization == "" {
			return ctx.JSON(401, helper.Response("Missing Token for Authorization!"))
		}

		token := authorization[len("Bearer "):]
		claims := helpers.ExtractToken(token)

		if claims == nil {
			return ctx.JSON(401, helpers.Response("Invalid Token Given!"))
		}

		userID := int(claims["user-id"].(float64))
		if role := claims["role"]; role == "member" {
			return ctx.JSON(401, helpers.Response("Role is not Recognized for this Feature!"))
		}

		input.StaffID = userID 

		staffReply := ctl.service.AddAReply(input, feedbackID)

		if staffReply == nil {
			return ctx.JSON(500, helper.Response("Something Went Wrong!"))
		}
		
		return ctx.JSON(200, helper.Response("Feedback Success Updated!", map[string]any {
			"data": staffReply,
		}))
	}
}

func (ctl *controller) DeleteFeedback() echo.HandlerFunc {
	return func (ctx echo.Context) error  {
		feedbackID, err := strconv.Atoi(ctx.Param("id"))

		if err != nil {
			return ctx.JSON(400, helper.Response("Param must be provided in number!"))
		}

		feedback := ctl.service.FindByID(feedbackID)

		if feedback == nil {
			return ctx.JSON(404, helper.Response("Feedback Not Found!"))
		}

		delete := ctl.service.Remove(feedbackID)

		if !delete {
			return ctx.JSON(500, helper.Response("Something Went Wrong!"))
		}

		return ctx.JSON(200, helper.Response("Feedback Success Deleted!", nil))
	}
}
