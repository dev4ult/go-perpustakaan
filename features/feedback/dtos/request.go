package dtos

type InputFeedback struct {
	MemberID *int `json:"user-id"`
	Comment string `json:"comment" form:"comment" validate:"required"`
}

type Prediction struct {
	Label string  `json:"label"`
	Score float64 `json:"score"`
}

type Pagination struct {
	Page int `query:"page"`
	Size int `query:"size"`
}