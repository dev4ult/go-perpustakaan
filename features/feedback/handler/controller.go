package handler

import (
	"perpustakaan/helpers"
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
		member := ctx.QueryParam("member")
		priority := ctx.QueryParam("priority")

		if size == 0 {
			page = 1
			size = 10
		}

		feedbacks, message := ctl.service.FindAll(page, size, member, priority)

		if message != "" {
			return ctx.JSON(404, helpers.Response(message))
		}

		if feedbacks == nil {
			return ctx.JSON(404, helpers.Response("There is No Feedbacks!"))
		}

		return ctx.JSON(200, helpers.Response("Success!", map[string]any {
			"data": feedbacks,
		}))
	}
}

func (ctl *controller) FeedbackDetails() echo.HandlerFunc {
	return func (ctx echo.Context) error  {
		feedbackID, err := strconv.Atoi(ctx.Param("id"))

		if err != nil {
			return ctx.JSON(400, helpers.Response("Param must be provided in number!"))
		}

		feedback, message := ctl.service.FindByID(feedbackID)

		if feedback == nil {
			return ctx.JSON(404, helpers.Response(message))
		}

		return ctx.JSON(200, helpers.Response("Success!", map[string]any {
			"data": feedback,
		}))
	}
}

func (ctl *controller) CreateFeedback() echo.HandlerFunc {
	return func (ctx echo.Context) error  {
		input := ctx.Get("request").(*dtos.InputFeedback)
		var userID *int
		
		auth := ctx.Request().Header.Get("Authorization")
		if auth != "" {
			token := auth[len("Bearer "):]
			claims := helpers.ExtractToken(token, false)

			if claims == nil {
				return ctx.JSON(401, helpers.Response("token is invalid or expired!"))
			}

			uid := int(claims["user-id"].(float64))

			userID = &uid
		}

		if userID != nil {
			input.MemberID = userID 
		}

		feedback, message := ctl.service.Create(*input)

		if feedback == nil {
			return ctx.JSON(500, helpers.Response(message))
		}

		return ctx.JSON(200, helpers.Response("Success!", map[string]any {
			"data": feedback,
		}))
	}
}

func (ctl *controller) ReplyOnFeedback() echo.HandlerFunc {
	return func (ctx echo.Context) error {
		input := ctx.Get("request").(*dtos.InputReply)

		feedbackID, err := strconv.Atoi(ctx.Param("id"))

		if err != nil {
			return ctx.JSON(400, helpers.Response(err.Error()))
		}

		feedback, message := ctl.service.FindByID(feedbackID)

		if feedback == nil {
			return ctx.JSON(404, helpers.Response(message))
		}
		
		if feedback.Reply.Staff != "" {
			return ctx.JSON(409, helpers.Response("Feedback Already has A Reply!"))
		}

		userID := int(ctx.Get("user-id").(float64))

		input.StaffID = userID 

		staffReply, errMessage := ctl.service.AddAReply(*input, feedbackID)

		if staffReply == nil {
			return ctx.JSON(500, helpers.Response(errMessage))
		}

		feedback.Reply.Staff = staffReply.Staff
		feedback.Reply.Comment = staffReply.Comment
		
		return ctx.JSON(200, helpers.Response("Feedback Success Updated!", map[string]any {
			"data": feedback,
		}))
	}
}

func (ctl *controller) DeleteFeedback() echo.HandlerFunc {
	return func (ctx echo.Context) error  {
		feedbackID, err := strconv.Atoi(ctx.Param("id"))

		if err != nil {
			return ctx.JSON(400, helpers.Response("Param must be provided in number!"))
		}

		feedback, message := ctl.service.FindByID(feedbackID)

		if feedback == nil {
			return ctx.JSON(404, helpers.Response(message))
		}

		delete, deleteMessage := ctl.service.Remove(feedbackID)

		if !delete {
			return ctx.JSON(500, helpers.Response(deleteMessage))
		}

		return ctx.JSON(200, helpers.Response("Feedback Success Deleted!", nil))
	}
}
