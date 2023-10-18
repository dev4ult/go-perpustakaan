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
	var loanHistorys []dtos.ResLoanHistory

	loanHistorysEnt := svc.model.Paginate(page, size)

	for _, loanHistory := range loanHistorysEnt {
		var data dtos.ResLoanHistory

		if err := smapping.FillStruct(&data, smapping.MapFields(loanHistory)); err != nil {
			log.Error(err.Error())
		} 
		
		loanHistorys = append(loanHistorys, data)
	}

	return loanHistorys
}

func (svc *service) FindByID(loanHistoryID int) *dtos.ResLoanHistory {
	res := dtos.ResLoanHistory{}
	loanHistory := svc.model.SelectByID(loanHistoryID)

	if loanHistory == nil {
		return nil
	}

	err := smapping.FillStruct(&res, smapping.MapFields(loanHistory))
	if err != nil {
		log.Error(err)
		return nil
	}

	return &res
}

func (svc *service) Create(newLoanHistory dtos.InputLoanHistory) *dtos.ResLoanHistory {
	loanHistory := loan_history.LoanHistory{}
	
	err := smapping.FillStruct(&loanHistory, smapping.MapFields(newLoanHistory))
	if err != nil {
		log.Error(err)
		return nil
	}

	loanHistoryID := svc.model.Insert(loanHistory)

	if loanHistoryID == -1 {
		return nil
	}

	resLoanHistory := dtos.ResLoanHistory{}
	errRes := smapping.FillStruct(&resLoanHistory, smapping.MapFields(newLoanHistory))
	if errRes != nil {
		log.Error(errRes)
		return nil
	}

	return &resLoanHistory
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
		log.Error("There is No LoanHistory Updated!")
		return false
	}
	
	return true
}

func (svc *service) Remove(loanHistoryID int) bool {
	rowsAffected := svc.model.DeleteByID(loanHistoryID)

	if rowsAffected <= 0 {
		log.Error("There is No LoanHistory Deleted!")
		return false
	}

	return true
}