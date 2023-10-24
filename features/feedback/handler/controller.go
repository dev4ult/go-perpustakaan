package handler

import (
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
		input := ctx.Get("request").(*dtos.InputFeedback)
		userID := int(ctx.Get("user-id").(float64))

		input.MemberID = &userID 

		feedback := ctl.service.Create(*input)

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
		input := ctx.Get("request").(*dtos.InputReply)

		feedbackID, err := strconv.Atoi(ctx.Param("id"))

		if err != nil {
			return ctx.JSON(400, helper.Response(err.Error()))
		}

		feedback := ctl.service.FindByID(feedbackID)

		if feedback == nil {
			return ctx.JSON(404, helper.Response("Feedback Not Found!"))
		}
		
		if feedback.Reply.Staff != "" {
			return ctx.JSON(409, helper.Response("Feedback Already has A Reply!"))
		}

		userID := int(ctx.Get("user-id").(float64))

		input.StaffID = userID 

		staffReply := ctl.service.AddAReply(*input, feedbackID)

		if staffReply == nil {
			return ctx.JSON(500, helper.Response("Something Went Wrong!"))
		}

		feedback.Reply.Staff = staffReply.Staff
		feedback.Reply.Comment = staffReply.Comment
		
		return ctx.JSON(200, helper.Response("Feedback Success Updated!", map[string]any {
			"data": feedback,
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
