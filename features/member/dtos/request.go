package dtos

type InputMember struct {
	FullName string `json:"full-name" form:"full-name" validate:"required"`
	CredentialNumber string `json:"credential-number" form:"credential-number" validate:"required"`
	Email string `json:"email" form:"email" validate:"required"`
	Password string `json:"password" form:"password" validate:"required"`
	PhoneNumber string `json:"phone-number" form:"phone-number" validate:"required"`
	Address string `json:"address" form:"address" validate:"required"`
}

type Pagination struct {
	Page int `query:"page"`
	Size int `query:"size"`
}