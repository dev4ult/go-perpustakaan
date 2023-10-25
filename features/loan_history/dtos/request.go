package dtos

type InputLoanHistory struct {
	StartToLoanAt 	string `json:"start-to-loan-at" form:"start-to-loan-at" validate:"required"`
	DueDate			string `json:"due-date" form:"due-date" validate:"required"`
	BookID     		int `json:"book-id" form:"book-id" validate:"required"`
	MemberID   		int `json:"member-id" form:"member-id" validate:"required"`
	FineTypeID 		*int `json:"fine-type-id" form:"fine-type-id"`
}

type UpdateLoanHistory struct {
	StartToLoanAt 	string `json:"start-to-loan-at" form:"start-to-loan-at" validate:"required"`
	DueDate			string `json:"due-date" form:"due-date" validate:"required"`
	BookID     		int `json:"book-id" form:"book-id" validate:"required"`
	LoanStatusID    int `json:"loan-status-id" form:"loan-status-id" validate:"required"`
	MemberID   		int `json:"member-id" form:"member-id" validate:"required"`
	FineTypeID 		*int `json:"fine-type-id" form:"fine-type-id"`
}

type LoanStatus struct {
	Status int `json:"status" form:"status" validate:"required"`
}

type Pagination struct {
	Page int `query:"page"`
	Size int `query:"size"`
}