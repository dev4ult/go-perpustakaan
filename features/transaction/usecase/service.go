package usecase

import (
	"perpustakaan/features/transaction"
	"perpustakaan/features/transaction/dtos"

	"github.com/labstack/gommon/log"
	"github.com/mashingan/smapping"
)

type service struct {
	model transaction.Repository
}

func New(model transaction.Repository) transaction.Usecase {
	return &service {
		model: model,
	}
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

func (svc *service) Create(newTransaction dtos.InputTransaction) *dtos.ResTransaction {
	transaction := transaction.Transaction{}
	
	err := smapping.FillStruct(&transaction, smapping.MapFields(newTransaction))
	if err != nil {
		log.Error(err)
		return nil
	}

	transactionID := svc.model.Insert(transaction)

	if transactionID == -1 {
		return nil
	}

	resTransaction := dtos.ResTransaction{}
	errRes := smapping.FillStruct(&resTransaction, smapping.MapFields(newTransaction))
	if errRes != nil {
		log.Error(errRes)
		return nil
	}

	return &resTransaction
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