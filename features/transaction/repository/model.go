package repository

import (
	"perpustakaan/features/member"
	"perpustakaan/features/transaction"
	"perpustakaan/features/transaction/dtos"
	"perpustakaan/helpers"

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

func (mdl *model) Paginate(page, size int) ([]transaction.Transaction, error) {
	var transactions []transaction.Transaction

	offset := (page - 1) * size

	if err := mdl.db.Offset(offset).Limit(size).Find(&transactions).Error; err != nil {
		return nil, err
	}

	return transactions, nil
}

func (mdl *model) Insert(newTransaction transaction.Transaction) (int, error) {
	if err := mdl.db.Create(&newTransaction).Error; err != nil {
		return 0, err
	}

	return newTransaction.ID, nil
}

func (mdl *model) SelectByID(transactionID int) (*transaction.Transaction, error) {
	var transaction transaction.Transaction

	if err := mdl.db.First(&transaction, transactionID).Error; err != nil {
		return nil, err
	}

	return &transaction, nil
}

func (mdl *model) Update(transaction transaction.Transaction) (int, error) {
	result := mdl.db.Save(&transaction)
	
	if result.Error != nil {
		return 0, result.Error
	}

	return int(result.RowsAffected), nil
}

func (mdl *model) DeleteByID(transactionID int) (int, error) {
	result := mdl.db.Delete(&transaction.Transaction{}, transactionID)
	
	if result.Error != nil {
		return 0, result.Error
	}

	return int(result.RowsAffected), nil
}

func (mdl *model) UpdateBatchTransactionDetail(items []dtos.FineItem, transactionID int) (bool, error) {
	for _, item := range items {
		if err := mdl.db.Table("loan_histories").Where("id = ?", item.ID).Update("transaction_id", transactionID).Error; err != nil {
			return false, err
		}
	}

	return true, nil
}

func (mdl *model) SelectAllFineItemOnMemberID(memberID int) ([]dtos.FineItem, error) {
	var fineItems []dtos.FineItem

	if err := mdl.db.Table("loan_histories").
	Select("loan_histories.id, fine_types.status, books.title as name, fine_types.fine_cost as amount").
	Where("loan_histories.member_id = ?", memberID).
	Where("fine_types.fine_cost IS NOT NULL").
	Where("loan_histories.transaction_id IS NULL").
	Where("loan_histories.deleted_at IS NULL").
	Joins("LEFT JOIN books ON books.id = loan_histories.book_id").
	Joins("LEFT JOIN fine_types ON fine_types.id = loan_histories.fine_type_id").
	Find(&fineItems).Error; err != nil {
		return nil, err
	}

	return fineItems, nil
}

func (mdl *model) SelectAllFineItemOnTransactionID(transactionID int) ([]dtos.FineItem, error) {
	var fineItems []dtos.FineItem

	if err := mdl.db.Table("loan_histories").
	Select("loan_histories.id, fine_types.status, books.title as name, fine_types.fine_cost as amount").
	Where("loan_histories.transaction_id = ?", transactionID).
	Where("fine_types.fine_cost IS NOT NULL").
	Where("loan_histories.transaction_id IS NOT NULL").
	Where("loan_histories.deleted_at IS NULL").
	Joins("LEFT JOIN books ON books.id = loan_histories.book_id").
	Joins("LEFT JOIN fine_types ON fine_types.id = loan_histories.fine_type_id").
	Find(&fineItems).Error; err != nil {
		return nil, err
	}

	return fineItems, nil
}

func (mdl *model) SelectFineItemByIDAndMemberID(fineItemID, memberID int) (*dtos.FineItem, error) {
	var fineItem dtos.FineItem

	if err := mdl.db.Table("loan_histories").
	Select("loan_histories.id, fine_types.status, books.title as name, fine_types.fine_cost as amount").
	Where("loan_histories.id = ?", fineItemID).
	Where("loan_histories.member_id = ?", memberID).
	Where("fine_types.fine_cost IS NOT NULL").
	Where("loan_histories.transaction_id IS NULL").
	Where("loan_histories.deleted_at IS NULL").
	Joins("LEFT JOIN books ON books.id = loan_histories.book_id").
	Joins("LEFT JOIN fine_types ON fine_types.id = loan_histories.fine_type_id").
	First(&fineItem).Error; err != nil {
		return nil, err
	}

	return &fineItem, nil
}

func (mdl *model) SelectMemberByID(memberID int) (*member.Member, error) {
	var member member.Member
	
	if err := mdl.db.Table("members").Where("id = ?", memberID).First(&member).Error; err != nil {
		return nil, err
	}

	return &member, nil
}

func (mdl *model) SelectTransactionByOrderID(orderID string) (*transaction.Transaction, error) {
	var transaction transaction.Transaction
	
	if err := mdl.db.Where("order_id = ?", orderID).First(&transaction).Error; err != nil {
		return nil, err
	}

	return &transaction, nil
}

func (mdl *model) UpdateStatus(transactionID int, status string) (bool, error) {
	if err := mdl.db.Table("transactions").Where("id = ?", transactionID).Update("status", status).Error; err != nil {
		return false, err
	}

	return true, nil
}

func (mdl *model) UnsetTransactionIDs(transactionID int) (bool, error) {
	if err := mdl.db.Table("loan_histories").Where("transaction_id = ?", transactionID).Update("transaction_id", "NULL").Error; err != nil {
		return false, err
	}
	
	return true, nil
}

func (mdl *model) GetTransactionStatusByOrderID(orderID string) (string, error) {
	status, err := helpers.CheckTransaction(orderID)
	if err != nil {
		return "", err
	}
	
	return status, nil
}