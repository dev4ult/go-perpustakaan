package dtos

type InputTransaction struct {
	Note string `json:"note" form:"note"`
	MemberID int `json:"member_id" form:"member_id" validate:"required"`
	LoanIDS []int `json:"loan_ids" form:"loan_ids"`
}

type Pagination struct {
	Page int `query:"page"`
	Size int `query:"size"`
}