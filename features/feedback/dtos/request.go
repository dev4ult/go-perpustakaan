package dtos

type InputFeedback struct {
	MemberID int `json:"user-id" validate:"required"`
	Comment string `json:"name" form:"name" validate:"required"`
}

type Pagination struct {
	Page int `query:"page"`
	Size int `query:"size"`
}