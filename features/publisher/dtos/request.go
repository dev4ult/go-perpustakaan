package dtos

type InputPublisher struct {
	Name string `json:"name" form:"name" validate:"required"`
	Country string `json:"country" form:"country" validate:"required"`
	City string `json:"city" form:"city" validate:"required"`
	Address string `json:"address" form:"address" validate:"required"`
	PostalCode int `json:"postal_code" form:"postal_code" validate:"required"`
	PhoneNumber string `json:"phone_number" form:"phone_number" validate:"required"`
	Email string `json:"email" form:"email" validate:"required"`
}

type Pagination struct {
	Page int `query:"page"`
	Size int `query:"size"`
}