package utils

import (
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/snap"
)

func SnapClient(midtransServerKey string) snap.Client {
	var snapClient snap.Client
	snapClient.New(midtransServerKey, midtrans.Sandbox)
	return snapClient
}