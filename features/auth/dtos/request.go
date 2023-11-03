package dtos

type InputLogin struct {
	CredentialNumber string `json:"credential_number" form:"credential_number"`
	StaffID string `json:"staff_id" form:"staff_id"`
	Password string `json:"password" form:"password" validate:"required"`
}

type InputStaffRegistration struct {
	FullName string `json:"full_name" form:"full_name" validate:"required"`
	StaffID string `json:"staff_id" form:"staff_id" validate:"required"`
	NIK int `json:"nik" form:"nik" validate:"required"`
	PhoneNumber string `json:"phone_number" form:"phone_number" validate:"required"`
	Address string `json:"address" form:"address" validate:"required"`
	Email string `json:"email" form:"email" validate:"required"`
	Password string `json:"password" form:"password" validate:"required"`
}

type Authorization struct {
	AccessToken string `json:"access_token" form:"access_token"`
}