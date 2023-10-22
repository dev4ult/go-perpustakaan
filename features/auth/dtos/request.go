package dtos

type InputLogin struct {
	CredentialNumber string `json:"credential-number" form:"credential-number"`
	StaffID string `json:"staff-id" form:"staff-id"`
	Password string `json:"password" form:"password" validate:"required"`
}

type Authorization struct {
	AccessToken string `json:"access-token" form:"access-token"`
}