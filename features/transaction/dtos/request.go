package dtos

type InputTransaction struct {
	Note string `json:"note" form:"note"`
	MemberID int `json:"member-id" form:"member-id" validate:"required"`
	LoanIDS []int `json:"loan-ids" form:"loan-ids"`
}

type Pagination struct {
	Page int `query:"page"`
	Size int `query:"size"`
}