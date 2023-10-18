package repository

import (
	"perpustakaan/features/loan_history"

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

func (mdl *model) Paginate(page, size int) []loan_history.LoanHistory {
	var loanHistorys []loan_history.LoanHistory

	offset := (page - 1) * size

	result := mdl.db.Offset(offset).Limit(size).Find(&loanHistorys)
	
	if result.Error != nil {
		log.Error(result.Error)
		return nil
	}

	return loanHistorys
}

func (mdl *model) Insert(newLoanHistory loan_history.LoanHistory) int64 {
	result := mdl.db.Create(&newLoanHistory)

	if result.Error != nil {
		log.Error(result.Error)
		return -1
	}

	return int64(newLoanHistory.ID)
}

func (mdl *model) SelectByID(loanHistoryID int) *loan_history.LoanHistory {
	var loanHistory loan_history.LoanHistory
	result := mdl.db.First(&loanHistory, loanHistoryID)

	if result.Error != nil {
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

func (mdl *model) DeleteByID(loanHistoryID int) int64 {
	result := mdl.db.Delete(&loan_history.LoanHistory{}, loanHistoryID)
	
	if result.Error != nil {
		log.Error(result.Error)
		return 0
	}

	return result.RowsAffected
}