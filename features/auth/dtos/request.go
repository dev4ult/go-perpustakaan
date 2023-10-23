package dtos

type InputLogin struct {
	CredentialNumber string `json:"credential-number" form:"credential-number"`
	StaffID string `json:"staff-id" form:"staff-id"`
	Password string `json:"password" form:"password" validate:"required"`
}

type InputStaffRegistration struct {
	FullName string `json:"full-name" form:"full-name" validate:"required"`
	StaffID string `json:"staff-id" form:"staff-id" validate:"required"`
	NIK int `json:"nik" form:"nik" validate:"required"`
	PhoneNumber string `json:"phone-number" form:"phone-number" validate:"required"`
	Address string `json:"address" form:"address" validate:"required"`
	Email string `json:"email" form:"email" validate:"required"`
	Password string `json:"password" form:"password" validate:"required"`
}

type Authorization struct {
	AccessToken string `json:"access-token" form:"access-token"`
}