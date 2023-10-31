package usecase

import (
	"fmt"
	"perpustakaan/config"
	"perpustakaan/features/transaction"
	"perpustakaan/features/transaction/dtos"
	"perpustakaan/helpers"
	"perpustakaan/utils"
	"strings"

	"github.com/google/uuid"
	"github.com/labstack/gommon/log"
	"github.com/mashingan/smapping"
	"github.com/midtrans/midtrans-go"
)

type service struct {
	model transaction.Repository
}

func New(model transaction.Repository) transaction.Usecase {
	return &service {
		model: model,
	}
}

func (svc *service) FindAll(page, size int) ([]dtos.ResTransaction, string) {
	var res []dtos.ResTransaction

	transactions, err := svc.model.Paginate(page, size)

	if err != nil {
		return nil, err.Error()
	}

	for _, transaction := range transactions {
		var data dtos.ResTransaction

		if err := smapping.FillStruct(&data, smapping.MapFields(transaction)); err != nil {
			return nil, err.Error()
		} 
		
		res = append(res, data)
	}

	return res, ""
}

func (svc *service) FindByID(transactionID int) (*dtos.ResTransaction, string) {
	var res dtos.ResTransaction
	transaction, err := svc.model.SelectByID(transactionID)

	if err != nil {
		return nil, err.Error()
	}
	
	if err := smapping.FillStruct(&res, smapping.MapFields(transaction)); err != nil {
		return nil, err.Error()
	}

	fineItems, err := svc.model.SelectAllFineItemOnTransactionID(transactionID)
	
	if err != nil {
		return nil, err.Error()
	}

	res.Fines = fineItems

	return &res, ""
}

func (svc *service) Create(newTransaction dtos.InputTransaction) (*dtos.ResTransaction, string) {
	var fineItems []dtos.FineItem

	member, err := svc.model.SelectMemberByID(newTransaction.MemberID)
	
	if err != nil {
		return nil, err.Error()
	}

	if len(newTransaction.LoanIDS) > 0 {
		for _, loanID := range newTransaction.LoanIDS {
			loanHistory, _ := svc.model.SelectFineItemByIDAndMemberID(loanID, newTransaction.MemberID) 

			if loanHistory == nil {
				return nil, "Unknown Loan History for this Member or Loan Status might not already be Changed from Checked-Out or On-Hold!"
			}

			fineItems = append(fineItems, *loanHistory)
		}
	} else {
		fineItems, _ = svc.model.SelectAllFineItemOnMemberID(newTransaction.MemberID)

		if fineItems == nil || len(fineItems) == 0 {
			return nil, "Member does not have any Fines yet!"
		}
	}

	var resTransaction = dtos.ResTransaction{
		Note: newTransaction.Note,
		Status: "Pending",
	}

	var midtransItems []midtrans.ItemDetails
	var totalPrice int64 = 0
	for _, item := range fineItems {
		midtransItems = append(midtransItems, midtrans.ItemDetails{
			ID: fmt.Sprintf("BOOK-%d", item.ID),
			Name: item.Name,
			Price: item.Amount,
			Qty: 1,
		})
		totalPrice += item.Amount
	}

	cfg := config.LoadServerConfig()
	snapClient := utils.SnapClient(cfg.MT_SERVER_KEY)

	customer := midtrans.CustomerDetails{
		FName: member.FullName,
		Email: member.Email,
		Phone: member.PhoneNumber,
	}

	randomID := uuid.New()

	orderID := fmt.Sprintf("LOAN-%s", strings.ToUpper(randomID.String()))
	
	snapRequest, err := helpers.CreatePaymentLink(snapClient, orderID, totalPrice, midtransItems, customer)

	if err != nil {
		return nil, err.Error()
	}

	resTransaction.PaymentURL = snapRequest.RedirectURL
	
	transaction := transaction.Transaction{}
	if err := smapping.FillStruct(&transaction, smapping.MapFields(newTransaction)); err != nil {
		log.Error(err)
		return nil, "Something Went Wrong!"
	}
	transaction.Status = "Pending"
	transaction.PaymentURL = snapRequest.RedirectURL
	transaction.OrderID = orderID

	transactionID, err := svc.model.Insert(transaction)
	if err != nil {
		return nil, err.Error()
	}

	if _, err := svc.model.UpdateBatchTransactionDetail(fineItems, transactionID); err != nil {
		return nil, "Error Set Transaction ID for Loan History!"
	}

	resTransaction.Fines = fineItems

	return &resTransaction, ""
}

func (svc *service) Modify(transactionData dtos.InputTransaction, transactionID int, orderID string, status string, paymentURL string) (bool, string) {
	var fineItems []dtos.FineItem
	
	if len(transactionData.LoanIDS) > 0 {
		_, err := svc.model.UnsetTransactionIDs(transactionID)

		if err != nil {
			return false, err.Error()
		}

		for _, loanID := range transactionData.LoanIDS {
			loanHistory, _ := svc.model.SelectFineItemByIDAndMemberID(loanID, transactionData.MemberID) 

			if loanHistory == nil {
				return false, "Unknown Loan History for this Member or Loan Status might not already be Changed from Checked-Out or On-Hold!"
			}

			fineItems = append(fineItems, *loanHistory)
		}
	} else {
		fineItems, _ = svc.model.SelectAllFineItemOnMemberID(transactionData.MemberID)

		if fineItems == nil || len(fineItems) == 0 {
			return false, "Member does not have any Fines yet!"
		}
	}

	var newTransaction transaction.Transaction
	
	if err := smapping.FillStruct(&newTransaction, smapping.MapFields(transactionData)); err != nil {
		return false, err.Error()
	}

	newTransaction.ID = transactionID
	newTransaction.OrderID = orderID
	newTransaction.Status = status
	newTransaction.PaymentURL = paymentURL
	_, err := svc.model.Update(newTransaction)

	if err != nil {
		return false, err.Error()
	}

	if _, err := svc.model.UpdateBatchTransactionDetail(fineItems, transactionID); err != nil {
		return false, "Error Set Transaction ID for Loan History!"
	}
	
	return true, ""
}

func (svc *service) Remove(transactionID int) (bool, string) {
	_, err := svc.model.DeleteByID(transactionID)

	if err != nil {
		return false, err.Error()
	}

	return true, ""
}

func (svc *service) VerifyPayment(payload map[string]any) (bool, string) {
	orderID, exist := payload["order_id"].(string)
	if !exist {
		return false, "Invalid Notification!"
	}

	status, err := svc.model.GetTransactionStatusByOrderID(orderID)
	if err != nil {
		return false, err.Error()
	}

	transaction, err := svc.model.SelectTransactionByOrderID(orderID)
	if err != nil {
		return false, "Transaction Not Found!"
	}
	
	if _, err := svc.model.UpdateStatus(transaction.ID, status); err != nil {
		return false, "Update Transaction Status Failed!"
	}

	return true, ""
}