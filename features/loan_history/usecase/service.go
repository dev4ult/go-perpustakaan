package usecase

import (
	"perpustakaan/features/loan_history"
	"perpustakaan/features/loan_history/dtos"

	"github.com/mashingan/smapping"
)

type service struct {
	model loan_history.Repository
}

func New(model loan_history.Repository) loan_history.Usecase {
	return &service {
		model: model,
	}
}

func (svc *service) FindAll(page int, size int, memberName string, status string) ([]dtos.ResLoanHistory, string) {
	loanHistories, err := svc.model.Paginate(page, size, memberName, status)

	if err != nil {
		return nil, err.Error()
	}

	return loanHistories, ""
}

func (svc *service) FindByID(loanHistoryID int) (*dtos.ResLoanHistory, string) {
	loanHistory, err := svc.model.SelectByID(loanHistoryID)

	if err != nil {
		return nil, err.Error()
	}

	return loanHistory, ""
}

func (svc *service) Create(newLoanHistory dtos.InputLoanHistory) (*dtos.ResLoanHistory, string) {
	var loanHistory loan_history.LoanHistory

	if err := smapping.FillStruct(&loanHistory, smapping.MapFields(newLoanHistory)); err != nil {
		return nil, err.Error()
	}

	loanHistory.FineTypeID = 1
	res, err := svc.model.Insert(loanHistory)

	if err != nil {
		return nil, err.Error()
	}

	return res, ""
}

func (svc *service) Modify(loanHistoryData dtos.InputLoanHistory, loanHistoryID int) (bool, string) {
	var newLoanHistory loan_history.LoanHistory
	
	if err := smapping.FillStruct(&newLoanHistory, smapping.MapFields(loanHistoryData)); err != nil {
		return false, err.Error()
	}

	newLoanHistory.ID = loanHistoryID
	_, err := svc.model.Update(newLoanHistory)

	if err != nil {
		return false, err.Error()
	}
	
	return true, ""
}

func (svc *service) ModifyStatus(status int, statusBefore string, loanHistoryID int) (bool, string) {
	_, err := svc.model.UpdateStatus(status, statusBefore, loanHistoryID)

	if err != nil {
		return false, err.Error()
	}
	
	return true, ""
}

func (svc *service) Remove(loanHistoryID int) (bool, string) {
	_, err := svc.model.DeleteByID(loanHistoryID)

	if err != nil {
		return false, err.Error()
	}

	return true, ""
}