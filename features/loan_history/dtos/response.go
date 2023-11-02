package dtos

type ResLoanHistory struct {
	StartToLoanAt string `json:"start-to-loan-at"`
	DueDate	string `json:"due-date"`
	Status string `json:"status"`
	FullName string `json:"full-name"`
	CredentialNumber string `json:"credential-number"`
	Title string `json:"title"`
	CoverImage string `json:"cover-img"`
	Summary string `json:"summary"`
	TransactionStatus string `json:"transaction-status"`
}
