package helpers

import (
	"errors"
	"perpustakaan/config"
	"perpustakaan/utils"
)

func CheckTransaction(orderID string) (string, error) {
	cfg := config.LoadServerConfig()
	response, err := utils.CoreAPIClient(cfg.MT_SERVER_KEY).CheckTransaction(orderID)

	if err != nil {
		return "", err
	}
	
	if response != nil {
		var status = ""

		// 5. Do set transaction status based on response from check transaction status
		if response.TransactionStatus == "capture" {
			if response.FraudStatus == "challenge" {
				// TODO set transaction status on your database to 'challenge'
				// e.g: 'Payment status challenged. Please take action on your Merchant Administration Portal
				status = "Challenge"
			} else if response.FraudStatus == "accept" {
				// TODO set transaction status on your database to 'success'
				status = "Success"
			}
		} else if response.TransactionStatus == "settlement" {
			// TODO set transaction status on your databaase to 'success'
			status = "Success"
		} else if response.TransactionStatus == "deny" {
			// TODO you can ignore 'deny', because most of the time it allows payment retries
			// and later can become success
		} else if response.TransactionStatus == "cancel" || response.TransactionStatus == "expire" {
			// TODO set transaction status on your databaase to 'failure'
			status = "Failed"
		} else if response.TransactionStatus == "pending" {
			// TODO set transaction status on your databaase to 'pending' / waiting payment
			status = "Pending"
		}

		return status, nil
	}

	return "", errors.New("Unknown Response Status From Transaction!")
}