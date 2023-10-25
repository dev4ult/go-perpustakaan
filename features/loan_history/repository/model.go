package repository

import (
	"perpustakaan/features/loan_history"
	"perpustakaan/features/loan_history/dtos"

	"github.com/labstack/gommon/log"
	"gorm.io/gorm"
)

type model struct {
	db *gorm.DB
}

func New(db *gorm.DB) loan_history.Repository {
	return &model {
		db: db,
	}
}

func (mdl *model) Paginate(page, size int) []dtos.ResLoanHistory {
	var loanHistories []dtos.ResLoanHistory
	offset := (page - 1) * size
	
	if result := mdl.db.Table("loan_histories").
	Select("loan_histories.start_to_loan_at, loan_histories.due_date, loan_statuses.name as status, members.full_name, members.credential_number, books.title, books.cover_image, books.summary").
	Joins("LEFT JOIN members ON members.id = loan_histories.member_id").
	Joins("LEFT JOIN books ON books.id = loan_histories.book_id").
	Joins("LEFT JOIN loan_statuses ON loan_statuses.id = loan_histories.loan_status_id").Where("loan_histories.deleted_at IS NULL").
	Offset(offset).Limit(size).Find(&loanHistories); result.Error != nil {
		log.Error(result.Error)
		return nil
	}

	return loanHistories
}

func (mdl *model) Insert(newLoanHistory loan_history.LoanHistory) *dtos.ResLoanHistory {
	if result := mdl.db.Create(&newLoanHistory); result.Error != nil {
		log.Error(result.Error)
		return nil
	}
	
	loanHistory := mdl.SelectByID(newLoanHistory.ID)
	
	if loanHistory == nil {
		log.Error("Loan History Not Found!")
		return nil
	}

	return loanHistory
}

func (mdl *model) SelectByID(loanHistoryID int) *dtos.ResLoanHistory {
	var loanHistory = dtos.ResLoanHistory{} 

	if result := mdl.db.Table("loan_histories").
	Select("loan_histories.start_to_loan_at, loan_histories.due_date, loan_statuses.name as status, members.full_name, members.credential_number, books.title, books.cover_image, books.summary").Where("loan_histories.id = ?", loanHistoryID).Where("loan_histories.deleted_at IS NULL").
	Joins("LEFT JOIN members ON members.id = loan_histories.member_id").
	Joins("LEFT JOIN books ON books.id = loan_histories.book_id").
	Joins("LEFT JOIN loan_statuses ON loan_statuses.id = loan_histories.loan_status_id").
	First(&loanHistory); result.Error != nil {
		log.Error(result.Error)
		return nil
	}

	return &loanHistory
}

func (mdl *model) Update(loanHistory loan_history.LoanHistory) int64 {
	result := mdl.db.Save(&loanHistory)

	if result.Error != nil {
		log.Error(result.Error)
	}

	return result.RowsAffected
}

func (mdl *model) UpdateStatus(status, loanHistoryID int) int64 {
	result := mdl.db.Table("loan_histories").Where("id", loanHistoryID).Update("loan_status_id", status);
	if result.Error != nil {
		log.Error(result.Error)
	}

	return result.RowsAffected
}

func (mdl *model) DeleteByID(loanHistoryID int) int64 {
	result := mdl.db.Delete(&loan_history.LoanHistory{}, loanHistoryID)
	
	if result.Error != nil {
		log.Error(result.Error)
		return 0
	}

	return result.RowsAffected
}