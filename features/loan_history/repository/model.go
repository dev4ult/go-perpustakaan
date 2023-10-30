package repository

import (
	"perpustakaan/features/loan_history"
	"perpustakaan/features/loan_history/dtos"

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

func (mdl *model) Paginate(page int, size int, searchKey string) ([]dtos.ResLoanHistory, error) {
	var loanHistories []dtos.ResLoanHistory

	offset := (page - 1) * size
	memberName := "%" + searchKey + "%"
	
	if err := mdl.db.Table("loan_histories").
	Select("loan_histories.start_to_loan_at, loan_histories.due_date, fine_types.status, members.full_name, members.credential_number, books.title, books.cover_image, books.summary").
	Joins("LEFT JOIN members ON members.id = loan_histories.member_id").
	Joins("LEFT JOIN books ON books.id = loan_histories.book_id").
	Joins("LEFT JOIN fine_types ON fine_types.id = loan_histories.fine_type_id").
	Where("loan_histories.deleted_at IS NULL").
	Where("members.full_name LIKE ?", memberName).
	Offset(offset).Limit(size).Find(&loanHistories).Error; err != nil {
		return nil, err
	}

	return loanHistories, nil
}

func (mdl *model) Insert(newLoanHistory loan_history.LoanHistory) (*dtos.ResLoanHistory, error) {
	if err := mdl.db.Create(&newLoanHistory).Error; err != nil {
		return nil, err
	}
	
	loanHistory, err := mdl.SelectByID(newLoanHistory.ID)
	
	if err != nil {
		return nil, err
	}

	return loanHistory, nil
}

func (mdl *model) SelectByID(loanHistoryID int) (*dtos.ResLoanHistory, error) {
	var loanHistory = dtos.ResLoanHistory{} 

	if err := mdl.db.Table("loan_histories").
	Select("loan_histories.start_to_loan_at, loan_histories.due_date, fine_types.status, members.full_name, members.credential_number, books.title, books.cover_image, books.summary").Where("loan_histories.id = ?", loanHistoryID).Where("loan_histories.deleted_at IS NULL").
	Joins("LEFT JOIN members ON members.id = loan_histories.member_id").
	Joins("LEFT JOIN books ON books.id = loan_histories.book_id").
	Joins("LEFT JOIN fine_types ON fine_types.id = loan_histories.fine_type_id").
	First(&loanHistory).Error; err != nil {
		return nil, err
	}

	return &loanHistory, nil
}

func (mdl *model) Update(loanHistory loan_history.LoanHistory) (int, error) {
	result := mdl.db.Save(&loanHistory)

	if result.Error != nil {
		return 0, result.Error
	}

	return int(result.RowsAffected), nil
}

func (mdl *model) UpdateStatus(status, loanHistoryID int) (int, error) {
	result := mdl.db.Table("loan_histories").Where("id", loanHistoryID).Update("fine_type_id", status);

	if result.Error != nil {
		return 0, result.Error
	}

	return int(result.RowsAffected), nil
}

func (mdl *model) DeleteByID(loanHistoryID int) (int, error) {
	result := mdl.db.Delete(&loan_history.LoanHistory{}, loanHistoryID)
	
	if result.Error != nil {
		return 0, result.Error
	}

	return int(result.RowsAffected), nil
}