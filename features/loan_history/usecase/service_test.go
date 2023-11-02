package usecase

import (
	"errors"
	"perpustakaan/features/loan_history"
	"perpustakaan/features/loan_history/dtos"
	"perpustakaan/features/loan_history/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindAll(t *testing.T) {
	var repository = mocks.NewRepository(t)
	var service = New(repository)

	var loanHistories = []dtos.ResLoanHistory{
		{
			StartToLoanAt: "03/03/23",
			DueDate: "04/04/23",	
			Status: "Checked Out",
			FullName: "Sarbin Sisarbin",
			CredentialNumber: "2107411026",
			Title: "Dark Gathering",
			CoverImage: "",
			Summary: "Lorem Ipsum",
		},
	}

	var page = 1
	var size = 10
	var memberName = ""
	var status = ""

	t.Run("Success", func(t *testing.T) {
		repository.On("Paginate", page, size, memberName, status).Return(loanHistories, nil).Once()
		
		result, message := service.FindAll(page, size, memberName, status)
		assert.NotNil(t, result)
		assert.Equal(t, loanHistories[0].FullName, result[0].FullName)
		assert.Empty(t, message)
		repository.AssertExpectations(t)
	})

	t.Run("Failed", func(t *testing.T) {
		repository.On("Paginate", page, size, memberName, status).Return(nil, errors.New("record not found")).Once()
		
		result, message := service.FindAll(page, size, memberName, status)
		assert.Nil(t, result)
		assert.NotEmpty(t, message)
		repository.AssertExpectations(t)
	})
}

func TestFindByID(t *testing.T) {
	var repository = mocks.NewRepository(t)
	var service = New(repository)

	var loanHistory = dtos.ResLoanHistory{
		StartToLoanAt: "03/03/23",
		DueDate: "04/04/23",	
		Status: "Checked Out",
		FullName: "Sarbin Sisarbin",
		CredentialNumber: "2107411026",
		Title: "Dark Gathering",
		CoverImage: "",
		Summary: "Lorem Ipsum",
	}

	var loanHistoryID = 1

	t.Run("Success", func(t *testing.T) {
		repository.On("SelectByID", loanHistoryID).Return(&loanHistory, nil).Once()
		
		result, message := service.FindByID(loanHistoryID)
		assert.NotNil(t, result)
		assert.Equal(t, loanHistory.FullName, result.FullName)
		assert.Empty(t, message)
		repository.AssertExpectations(t)
	})

	t.Run("Failed", func(t *testing.T) {
		repository.On("SelectByID", 0).Return(nil, errors.New("record not found")).Once()
		
		result, message := service.FindByID(0)
		assert.Nil(t, result)
		assert.NotEmpty(t, message)
		repository.AssertExpectations(t)
	})
}

func TestCreate(t *testing.T) {
	var repository = mocks.NewRepository(t)
	var service = New(repository)

	var loanHistory = loan_history.LoanHistory{
		StartToLoanAt: "2023-02-02",
		DueDate: "2023-03-03",
		BookID: 2,
		MemberID: 3,
		FineTypeID: 1,
	}

	var invalidEntity = loan_history.LoanHistory{
		FineTypeID: 1,
	}

	var response = dtos.ResLoanHistory{
		StartToLoanAt: "03/03/23",
		DueDate: "04/04/23",	
		Status: "Checked Out",
		FullName: "Sarbin Sisarbin",
		CredentialNumber: "2107411026",
		Title: "Dark Gathering",
		CoverImage: "",
		Summary: "Lorem Ipsum",
	}

	var input = dtos.InputLoanHistory{
		StartToLoanAt: "2023-02-02",
		DueDate: "2023-03-03",
		BookID: 2,
		MemberID: 3,
		FineTypeID: 1,
	}

	var emptyInput = dtos.InputLoanHistory{}

	t.Run("Success", func(t *testing.T) {
		repository.On("Insert", loanHistory).Return(&response, nil).Once()
		
		result, message := service.Create(input)
		assert.NotNil(t, result)
		assert.Equal(t, response.FullName, result.FullName)
		assert.Empty(t, message)
		repository.AssertExpectations(t)
	})

	t.Run("Failed", func(t *testing.T) {
		repository.On("Insert", invalidEntity).Return(nil, errors.New("record not found")).Once()
		
		result, message := service.Create(emptyInput)
		assert.Nil(t, result)
		assert.NotEmpty(t, message)
		repository.AssertExpectations(t)
	})
}

func TestModify(t *testing.T) {
	var repository = mocks.NewRepository(t)
	var service = New(repository)

	var loanHistory = loan_history.LoanHistory{
		StartToLoanAt: "2023-02-02",
		DueDate: "2023-03-03",
		BookID: 2,
		MemberID: 3,
		FineTypeID: 1,
	}

	var invalidEntity = loan_history.LoanHistory{}

	var input = dtos.InputLoanHistory{
		StartToLoanAt: "2023-02-02",
		DueDate: "2023-03-03",
		BookID: 2,
		MemberID: 3,
		FineTypeID: 1,
	}

	var emptyInput = dtos.InputLoanHistory{}

	var loanHistoryID = 1

	t.Run("Success", func(t *testing.T) {
		loanHistory.ID = loanHistoryID
		repository.On("Update", loanHistory).Return(1, nil).Once()
		
		result, message := service.Modify(input, loanHistoryID)
		assert.Equal(t, true, result)
		assert.Empty(t, message)
		repository.AssertExpectations(t)
	})

	t.Run("Failed", func(t *testing.T) {
		repository.On("Update", invalidEntity).Return(0, errors.New("record not found")).Once()
		
		result, message := service.Modify(emptyInput, 0)
		assert.Equal(t, false, result)
		assert.NotEmpty(t, message)
		repository.AssertExpectations(t)
	})
}

func TestModifyStatus(t *testing.T) {
	var repository = mocks.NewRepository(t)
	var service = New(repository)

	var status = 2
	var statusBefore = "On Hold"
	var invalidStatus = 999

	var loanHistoryID = 1

	t.Run("Success", func(t *testing.T) {
		repository.On("UpdateStatus", status, statusBefore, loanHistoryID).Return(1, nil).Once()
		
		result, message := service.ModifyStatus(status, statusBefore, loanHistoryID)
		assert.Equal(t, true, result)
		assert.Empty(t, message)
		repository.AssertExpectations(t)
	})

	t.Run("Failed", func(t *testing.T) {
		repository.On("UpdateStatus", invalidStatus, statusBefore, 0).Return(0, errors.New("record not found")).Once()
		
		result, message := service.ModifyStatus(invalidStatus, statusBefore, 0)
		assert.Equal(t, false, result)
		assert.NotEmpty(t, message)
		repository.AssertExpectations(t)
	})
}

func TestRemove(t *testing.T) {
	var repository = mocks.NewRepository(t)
	var service = New(repository)

	var loanHistoryID = 1

	t.Run("Success", func(t *testing.T) {
		repository.On("DeleteByID", loanHistoryID).Return(1, nil).Once()
		
		result, message := service.Remove(loanHistoryID)
		assert.Equal(t, true, result)
		assert.Empty(t, message)
		repository.AssertExpectations(t)
	})

	t.Run("Failed", func(t *testing.T) {
		repository.On("DeleteByID", 0).Return(0, errors.New("record not found")).Once()
		
		result, message := service.Remove(0)
		assert.Equal(t, false, result)
		assert.NotEmpty(t, message)
		repository.AssertExpectations(t)
	})
}