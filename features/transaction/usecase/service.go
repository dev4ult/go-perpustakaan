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

func (svc *service) VerifyPayment(payload map[string]any) {
}

func (svc *service) FindAll(page, size int) []dtos.ResTransaction {
	var transactions []dtos.ResTransaction

	transactionsEnt := svc.model.Paginate(page, size)

	for _, transaction := range transactionsEnt {
		var data dtos.ResTransaction

		if err := smapping.FillStruct(&data, smapping.MapFields(transaction)); err != nil {
			log.Error(err.Error())
		} 
		
		transactions = append(transactions, data)
	}

	return transactions
}

func (svc *service) FindByID(transactionID int) *dtos.ResTransaction {
	res := dtos.ResTransaction{}
	transaction := svc.model.SelectByID(transactionID)

	if transaction == nil {
		return nil
	}

	err := smapping.FillStruct(&res, smapping.MapFields(transaction))
	if err != nil {
		log.Error(err)
		return nil
	}

	return &res
}

func (svc *service) Create(newTransaction dtos.InputTransaction) (*dtos.ResTransaction, string) {
	var fineItems []dtos.FineItem

	member := svc.model.SelectMemberByID(newTransaction.MemberID)
	
	if member == nil {
		return nil, "Unknown Member ID"
	}

	if len(newTransaction.LoanIDS) > 0 {
		for _, loanID := range newTransaction.LoanIDS {
			loanHistory := svc.model.SelectLoanHistoryByIDAndMemberID(loanID, newTransaction.MemberID) 

			if loanHistory == nil {
				return nil, "Unknown Loan History for this Member or Loan Status might not already be Changed from Checked-Out or On-Hold!"
			}
			fineItems = append(fineItems, *loanHistory)
		}
	} else {
		fineItems = svc.model.SelectAllLoanHistoryOnMemberID(newTransaction.MemberID)

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

	transactionID := svc.model.Insert(transaction)
	if transactionID == -1 {
		return nil, "Error When Creating a Transaction!"
	}

	setTransactionID := svc.model.UpdateBatchTransactionDetail(fineItems, transactionID)

	if !setTransactionID {
		return nil, "Error Set Transaction ID for Loan History!"
	}

	resTransaction.Fines = fineItems

	return &resTransaction, ""
}

func (svc *service) Modify(transactionData dtos.InputTransaction, transactionID int) bool {
	newTransaction := transaction.Transaction{}

	err := smapping.FillStruct(&newTransaction, smapping.MapFields(transactionData))
	if err != nil {
		log.Error(err)
		return false
	}

	newTransaction.ID = transactionID
	rowsAffected := svc.model.Update(newTransaction)

	if rowsAffected <= 0 {
		log.Error("There is No Transaction Updated!")
		return false
	}
	
	return true
}

func (svc *service) Remove(transactionID int) bool {
	rowsAffected := svc.model.DeleteByID(transactionID)

	if rowsAffected <= 0 {
		log.Error("There is No Transaction Deleted!")
		return false
	}

	return true
}
