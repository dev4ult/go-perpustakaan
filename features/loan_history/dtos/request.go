package dtos

type InputLoanHistory struct {
	StartToLoanAt 	string `json:"start_to_loan_at" form:"start_to_loan_at" validate:"required"`
	DueDate			string `json:"due_date" form:"due_date" validate:"required"`
	BookID     		int `json:"book_id" form:"book_id" validate:"required"`
	MemberID   		int `json:"member_id" form:"member_id" validate:"required"`
	FineTypeID 		int `json:"status_id" form:"status_id"`
}

type LoanStatus struct {
	Status int `json:"status" form:"status" validate:"required"`
}

type Pagination struct {
	Page int `query:"page"`
	Size int `query:"size"`
}