package dtos

type ResTransaction struct {
	Note string `json:"note"`
	Status string `json:"status"`
	PaymentURL string `json:"payment-url"`
	Fines []FineItem `json:"fine-item" gorm:"-"`
}

type FineItem struct {
	ID int `json:"id"`
	Name string `json:"name"`
	Status string `json:"status"`
	Amount int64 `json:"amount"`
}
