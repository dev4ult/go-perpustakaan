package dtos

type InputLogin struct {
	CredentialNumber string `json:"credential-number" form:"credential-number"`
	StaffID string `json:"staff-id" form:"staff-id"`
	Password string `json:"password" form:"password"`
}