package helpers

import (
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/snap"
)

func (h *helper) CreatePaymentLink(snapClient snap.Client, orderID string, totalPrice int64, items []midtrans.ItemDetails, customer midtrans.CustomerDetails) (*snap.Response, error) {
	snapRequest := &snap.Request{
		TransactionDetails: midtrans.TransactionDetails{
			OrderID: orderID,
			GrossAmt: totalPrice,
		},
		CustomerDetail: &customer,
		EnabledPayments: snap.AllSnapPaymentType,
		Items: &items,
	}

	snapResponse, err := snapClient.CreateTransaction(snapRequest)

	if err != nil {
		return nil, err
	}

	return snapResponse, nil
}