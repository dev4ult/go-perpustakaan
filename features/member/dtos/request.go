package dtos

type InputMember struct {
	FullName string `json:"full_name" form:"full_name" validate:"required"`
	CredentialNumber string `json:"credential_number" form:"credential_number" validate:"required"`
	Email string `json:"email" form:"email" validate:"required"`
	Password string `json:"password" form:"password" validate:"required"`
	PhoneNumber string `json:"phone_number" form:"phone_number" validate:"required"`
	Address string `json:"address" form:"address" validate:"required"`
}

type Pagination struct {
	Page int `query:"page"`
	Size int `query:"size"`
}