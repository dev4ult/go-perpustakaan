package repository

import (
	"perpustakaan/features/loanHistory"

	"github.com/labstack/gommon/log"
	"gorm.io/gorm"
)

type model struct {
	db *gorm.DB
}

func New(db *gorm.DB) loanHistory.Repository {
	return &model {
		db: db,
	}
}

func (mdl *model) Paginate(page, size int) []loanHistory.LoanHistory {
	var loanHistorys []loanHistory.LoanHistory

	offset := (page - 1) * size

	result := mdl.db.Offset(offset).Limit(size).Find(&loanHistorys)
	
	if result.Error != nil {
		log.Error(result.Error)
		return nil
	}

	return loanHistorys
}

func (mdl *model) Insert(newLoanHistory loanHistory.LoanHistory) int64 {
	result := mdl.db.Create(&newLoanHistory)

	if result.Error != nil {
		log.Error(result.Error)
		return -1
	}

	return int64(newLoanHistory.ID)
}

func (mdl *model) SelectByID(loanHistoryID int) *loanHistory.LoanHistory {
	var loanHistory loanHistory.LoanHistory
	result := mdl.db.First(&loanHistory, loanHistoryID)

	if result.Error != nil {
		log.Error(result.Error)
		return nil
	}

	return &loanHistory
}

func (mdl *model) Update(loanHistory loanHistory.LoanHistory) int64 {
	result := mdl.db.Save(&loanHistory)

	if result.Error != nil {
		log.Error(result.Error)
	}

	return result.RowsAffected
}

func (mdl *model) DeleteByID(loanHistoryID int) int64 {
	result := mdl.db.Delete(&loanHistory.LoanHistory{}, loanHistoryID)
	
	if result.Error != nil {
		log.Error(result.Error)
		return 0
	}

	return result.RowsAffected
}