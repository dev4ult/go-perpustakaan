// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	loan_history "perpustakaan/features/loan_history"
	dtos "perpustakaan/features/loan_history/dtos"

	mock "github.com/stretchr/testify/mock"
)

// Repository is an autogenerated mock type for the Repository type
type Repository struct {
	mock.Mock
}

// DeleteByID provides a mock function with given fields: loanHistoryID
func (_m *Repository) DeleteByID(loanHistoryID int) (int, error) {
	ret := _m.Called(loanHistoryID)

	var r0 int
	var r1 error
	if rf, ok := ret.Get(0).(func(int) (int, error)); ok {
		return rf(loanHistoryID)
	}
	if rf, ok := ret.Get(0).(func(int) int); ok {
		r0 = rf(loanHistoryID)
	} else {
		r0 = ret.Get(0).(int)
	}

	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(loanHistoryID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Insert provides a mock function with given fields: newLoanHistory
func (_m *Repository) Insert(newLoanHistory loan_history.LoanHistory) (*dtos.ResLoanHistory, error) {
	ret := _m.Called(newLoanHistory)

	var r0 *dtos.ResLoanHistory
	var r1 error
	if rf, ok := ret.Get(0).(func(loan_history.LoanHistory) (*dtos.ResLoanHistory, error)); ok {
		return rf(newLoanHistory)
	}
	if rf, ok := ret.Get(0).(func(loan_history.LoanHistory) *dtos.ResLoanHistory); ok {
		r0 = rf(newLoanHistory)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*dtos.ResLoanHistory)
		}
	}

	if rf, ok := ret.Get(1).(func(loan_history.LoanHistory) error); ok {
		r1 = rf(newLoanHistory)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Paginate provides a mock function with given fields: page, size, memberName, status
func (_m *Repository) Paginate(page int, size int, memberName string, status string) ([]dtos.ResLoanHistory, error) {
	ret := _m.Called(page, size, memberName, status)

	var r0 []dtos.ResLoanHistory
	var r1 error
	if rf, ok := ret.Get(0).(func(int, int, string, string) ([]dtos.ResLoanHistory, error)); ok {
		return rf(page, size, memberName, status)
	}
	if rf, ok := ret.Get(0).(func(int, int, string, string) []dtos.ResLoanHistory); ok {
		r0 = rf(page, size, memberName, status)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]dtos.ResLoanHistory)
		}
	}

	if rf, ok := ret.Get(1).(func(int, int, string, string) error); ok {
		r1 = rf(page, size, memberName, status)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SelectByID provides a mock function with given fields: loanHistoryID
func (_m *Repository) SelectByID(loanHistoryID int) (*dtos.ResLoanHistory, error) {
	ret := _m.Called(loanHistoryID)

	var r0 *dtos.ResLoanHistory
	var r1 error
	if rf, ok := ret.Get(0).(func(int) (*dtos.ResLoanHistory, error)); ok {
		return rf(loanHistoryID)
	}
	if rf, ok := ret.Get(0).(func(int) *dtos.ResLoanHistory); ok {
		r0 = rf(loanHistoryID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*dtos.ResLoanHistory)
		}
	}

	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(loanHistoryID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: loanHistory
func (_m *Repository) Update(loanHistory loan_history.LoanHistory) (int, error) {
	ret := _m.Called(loanHistory)

	var r0 int
	var r1 error
	if rf, ok := ret.Get(0).(func(loan_history.LoanHistory) (int, error)); ok {
		return rf(loanHistory)
	}
	if rf, ok := ret.Get(0).(func(loan_history.LoanHistory) int); ok {
		r0 = rf(loanHistory)
	} else {
		r0 = ret.Get(0).(int)
	}

	if rf, ok := ret.Get(1).(func(loan_history.LoanHistory) error); ok {
		r1 = rf(loanHistory)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateStatus provides a mock function with given fields: status, statusBefore, loanHistoryID
func (_m *Repository) UpdateStatus(status int, statusBefore string, loanHistoryID int) (int, error) {
	ret := _m.Called(status, statusBefore, loanHistoryID)

	var r0 int
	var r1 error
	if rf, ok := ret.Get(0).(func(int, string, int) (int, error)); ok {
		return rf(status, statusBefore, loanHistoryID)
	}
	if rf, ok := ret.Get(0).(func(int, string, int) int); ok {
		r0 = rf(status, statusBefore, loanHistoryID)
	} else {
		r0 = ret.Get(0).(int)
	}

	if rf, ok := ret.Get(1).(func(int, string, int) error); ok {
		r1 = rf(status, statusBefore, loanHistoryID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewRepository creates a new instance of Repository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *Repository {
	mock := &Repository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
