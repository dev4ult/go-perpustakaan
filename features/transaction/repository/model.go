package repository

import (
	"perpustakaan/features/member"
	"perpustakaan/features/transaction"
	"perpustakaan/features/transaction/dtos"

	"github.com/labstack/gommon/log"
	"gorm.io/gorm"
)

type model struct {
	db *gorm.DB
}

func New(db *gorm.DB) transaction.Repository {
	return &model {
		db: db,
	}
}

func (mdl *model) Paginate(page, size int) []transaction.Transaction {
	var transactions []transaction.Transaction

	offset := (page - 1) * size

	result := mdl.db.Offset(offset).Limit(size).Find(&transactions)
	
	if result.Error != nil {
		log.Error(result.Error)
		return nil
	}

	return transactions
}

func (mdl *model) Insert(newTransaction transaction.Transaction) int64 {
	result := mdl.db.Create(&newTransaction)

	if result.Error != nil {
		log.Error(result.Error)
		return -1
	}

	return int64(newTransaction.ID)
}

func (mdl *model) SelectByID(transactionID int) *transaction.Transaction {
	var transaction transaction.Transaction
	result := mdl.db.First(&transaction, transactionID)

	if result.Error != nil {
		log.Error(result.Error)
		return nil
	}

	return &transaction
}

func (mdl *model) Update(transaction transaction.Transaction) int64 {
	result := mdl.db.Save(&transaction)

	if result.Error != nil {
		log.Error(result.Error)
	}

	return result.RowsAffected
}

func (mdl *model) DeleteByID(transactionID int) int64 {
	result := mdl.db.Delete(&transaction.Transaction{}, transactionID)
	
	if result.Error != nil {
		log.Error(result.Error)
		return 0
	}

	return result.RowsAffected
}

func (mdl *model) UpdateBatchTransactionDetail(items []dtos.FineItem, transactionID int64) bool {
	for _, item := range items {
		if result := mdl.db.Table("loan_histories").Where("id = ?", item.ID).Update("transaction_id", transactionID); result.Error != nil {
			return false
		}
	}


	return true
}

func (mdl *model) SelectAllFineItemOnMemberID(memberID int) []dtos.FineItem {
	var fineItems = []dtos.FineItem{} 

	if result := mdl.db.Table("loan_histories").
	Select("loan_histories.id, fine_types.status, books.title as name, fine_types.fine_cost as amount").
	Where("loan_histories.member_id = ?", memberID).
	Where("fine_types.fine_cost IS NOT NULL").
	Where("loan_histories.transaction_id IS NULL").
	Where("loan_histories.deleted_at IS NULL").
	Joins("LEFT JOIN books ON books.id = loan_histories.book_id").
	Joins("LEFT JOIN fine_types ON fine_types.id = loan_histories.fine_type_id").
	Find(&fineItems); result.Error != nil {
		log.Error(result.Error)
		return nil
	}

	return fineItems
}

func (mdl *model) SelectAllFineItemOnTransactionID(transactionID int) []dtos.FineItem {
	var fineItems = []dtos.FineItem{} 

	if result := mdl.db.Table("loan_histories").
	Select("loan_histories.id, fine_types.status, books.title as name, fine_types.fine_cost as amount").
	Where("loan_histories.transaction_id = ?", transactionID).
	Where("fine_types.fine_cost IS NOT NULL").
	Where("loan_histories.transaction_id IS NOT NULL").
	Where("loan_histories.deleted_at IS NULL").
	Joins("LEFT JOIN books ON books.id = loan_histories.book_id").
	Joins("LEFT JOIN fine_types ON fine_types.id = loan_histories.fine_type_id").
	Find(&fineItems); result.Error != nil {
		log.Error(result.Error)
		return nil
	}

	return fineItems
}

func (mdl *model) SelectFineItemByIDAndMemberID(fineItemID, memberID int) *dtos.FineItem {
	var fineItem = dtos.FineItem{} 

	if result := mdl.db.Table("loan_histories").
	Select("loan_histories.id, fine_types.status, books.title as name, fine_types.fine_cost as amount").
	Where("loan_histories.id = ?", fineItemID).
	Where("loan_histories.member_id = ?", memberID).
	Where("fine_types.fine_cost IS NOT NULL").
	Where("loan_histories.transaction_id IS NULL").
	Where("loan_histories.deleted_at IS NULL").
	Joins("LEFT JOIN books ON books.id = loan_histories.book_id").
	Joins("LEFT JOIN fine_types ON fine_types.id = loan_histories.fine_type_id").
	First(&fineItem); result.Error != nil {
		log.Error(result.Error)
		return nil
	}

	return &fineItem
}

func (mdl *model) SelectMemberByID(memberID int) *member.Member {
	var member member.Member
	
	if result := mdl.db.Table("members").Where("id = ?", memberID).First(&member); result.Error != nil {
		log.Error(result.Error.Error())
		return nil
	}

	return &member
}

func (mdl *model) SelectTransactionByOrderID(orderID string) *transaction.Transaction {
	var transaction transaction.Transaction
	
	if result := mdl.db.Where("order_id = ?", orderID).First(&transaction); result.Error != nil {
		log.Error(result.Error.Error())
		return nil
	}

	return &transaction
}

func (mdl *model) UpdateStatus(transactionID int, status string) bool {
	if result := mdl.db.Table("transactions").Where("id = ?", transactionID).Update("status", status); result.Error != nil {
		log.Error(result.Error.Error())
		return false
	}

	return true
}