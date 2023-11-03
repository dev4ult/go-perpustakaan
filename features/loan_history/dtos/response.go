package dtos

type ResLoanHistory struct {
	StartToLoanAt string `json:"start_to_loan_at"`
	DueDate	string `json:"due_date"`
	Status string `json:"status"`
	FullName string `json:"full_name"`
	CredentialNumber string `json:"credential_number"`
	Title string `json:"title"`
	CoverImage string `json:"cover_img"`
	Summary string `json:"summary"`
	TransactionStatus string `json:"transaction_status"`
}
