package helpers

import (
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/snap"
)

type Helper interface {
	CheckTransaction(orderID string) (string, error)
	UploadImage(folder string, file any) (string, error)
	GetPrediction(comment string) string
	GenerateHash(password string) string
	GenerateToken(id int, role string) *JSONWebToken
	VerifyHash(password, hashed string) bool
	CreatePaymentLink(snapClient snap.Client, orderID string, totalPrice int64, items []midtrans.ItemDetails, customer midtrans.CustomerDetails) (*snap.Response, error)
}

type helper struct{}

func New() Helper {
	return &helper{}
}
