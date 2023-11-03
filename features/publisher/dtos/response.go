package dtos

type ResPublisher struct {
	Name string `json:"name"`
	Country string `json:"country"`
	City string `json:"city"`
	Address string `json:"address"`
	PostalCode int `json:"postal_code"`
	PhoneNumber string `json:"phone_number"`
	Email string `json:"email"`
}
