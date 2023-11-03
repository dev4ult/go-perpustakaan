package dtos

type ResTransaction struct {
	Note string `json:"note"`
	OrderID string `json:"order_id"`
	Status string `json:"status"`
	PaymentURL string `json:"payment_url"`
	Fines []FineItem `json:"fine_item" gorm:"_"`
}

type FineItem struct {
	ID int `json:"id"`
	Name string `json:"name"`
	Status string `json:"status"`
	Amount int64 `json:"amount"`
}
