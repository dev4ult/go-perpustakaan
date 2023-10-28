package utils

import (
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/coreapi"
	"github.com/midtrans/midtrans-go/snap"
)

func SnapClient(midtransServerKey string) snap.Client {
	var snapClient snap.Client
	snapClient.New(midtransServerKey, midtrans.Sandbox)
	return snapClient
}

func CoreAPIClient(midtransServerKey string) coreapi.Client {
	var coreAPIClient coreapi.Client
	coreAPIClient.New(midtransServerKey, midtrans.Sandbox)
	return coreAPIClient
}