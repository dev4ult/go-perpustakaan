package handler

import (
	"perpustakaan/helpers"
	helper "perpustakaan/helpers"
	"strconv"

	"perpustakaan/features/feedback"
	"perpustakaan/features/feedback/dtos"

	"github.com/go-playground/validator/v10"
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

var validate *validator.Validate

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
		input := dtos.InputFeedback{}

		ctx.Bind(&input)
		validate = validator.New(validator.WithRequiredStructEnabled())
		err := validate.Struct(input)

		if err != nil {
			errMap := helpers.ErrorMapValidation(err)
			return ctx.JSON(400, helper.Response("Missing Data Required!", map[string]any {
				"error": errMap,
			}))
		}

		feedback := ctl.service.Create(input)

		if feedback == nil {
			return ctx.JSON(500, helper.Response("Something went Wrong!", nil))
		}

		return ctx.JSON(200, helper.Response("Success!", map[string]any {
			"data": feedback,
		}))
	}
}


// Not sure

func (ctl *controller) UpdateFeedback() echo.HandlerFunc {
	return func (ctx echo.Context) error {
		input := dtos.InputFeedback{}

		feedbackID, errParam := strconv.Atoi(ctx.Param("id"))

		if errParam != nil {
			return ctx.JSON(400, helper.Response(errParam.Error()))
		}

		feedback := ctl.service.FindByID(feedbackID)

		if feedback == nil {
			return ctx.JSON(404, helper.Response("Feedback Not Found!"))
		}
		
		ctx.Bind(&input)

		validate = validator.New(validator.WithRequiredStructEnabled())
		err := validate.Struct(input)

		if err != nil {
			errMap := helpers.ErrorMapValidation(err)
			return ctx.JSON(400, helper.Response("Missing Data Required!", map[string]any {
				"error": errMap,
			}))
		}

		update := ctl.service.Modify(input, feedbackID)

		if !update {
			return ctx.JSON(500, helper.Response("Something Went Wrong!"))
		}

		return ctx.JSON(200, helper.Response("Feedback Success Updated!"))
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
