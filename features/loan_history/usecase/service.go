package usecase

import (
	"perpustakaan/features/loan_history"
	"perpustakaan/features/loan_history/dtos"

	"github.com/labstack/gommon/log"
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

func (svc *service) FindAll(page, size int) []dtos.ResLoanHistory {
	loanHistories := svc.model.Paginate(page, size)

	return loanHistories
}

func (svc *service) FindByID(loanHistoryID int) *dtos.ResLoanHistory {
	loanHistory := svc.model.SelectByID(loanHistoryID)

	if loanHistory == nil {
		return nil
	}

	return loanHistory
}

func (svc *service) Create(newLoanHistory dtos.InputLoanHistory) *dtos.ResLoanHistory {
	loanHistory := loan_history.LoanHistory{}
	if err := smapping.FillStruct(&loanHistory, smapping.MapFields(newLoanHistory)); err != nil {
		log.Error(err)
		return nil
	}

	loanHistory.FineTypeID = 1
	resLoanHistory := svc.model.Insert(loanHistory)
	if resLoanHistory == nil {
		return nil
	}

	return resLoanHistory
}

func (svc *service) Modify(loanHistoryData dtos.InputLoanHistory, loanHistoryID int) bool {
	newLoanHistory := loan_history.LoanHistory{}

	err := smapping.FillStruct(&newLoanHistory, smapping.MapFields(loanHistoryData))
	if err != nil {
		log.Error(err)
		return false
	}

	newLoanHistory.ID = loanHistoryID
	rowsAffected := svc.model.Update(newLoanHistory)

	if rowsAffected <= 0 {
		log.Error("There is No Loan History Updated!")
		return false
	}
	
	return true
}

func (svc *service) ModifyStatus(status, loanHistoryID int) bool {
	rowsAffected := svc.model.UpdateStatus(status, loanHistoryID)

	if rowsAffected <= 0 {
		log.Error("There is No Loan Status Updated!")
		return false
	}
	
	return true
}

func (svc *service) Remove(loanHistoryID int) bool {
	rowsAffected := svc.model.DeleteByID(loanHistoryID)

	if rowsAffected <= 0 {
		log.Error("There is No Loan History Deleted!")
		return false
	}

	return true
}