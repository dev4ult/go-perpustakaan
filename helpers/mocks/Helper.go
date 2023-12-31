// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	helpers "perpustakaan/helpers"

	midtrans "github.com/midtrans/midtrans-go"

	mock "github.com/stretchr/testify/mock"

	multipart "mime/multipart"

	snap "github.com/midtrans/midtrans-go/snap"
)

// Helper is an autogenerated mock type for the Helper type
type Helper struct {
	mock.Mock
}

// CheckTransaction provides a mock function with given fields: orderID
func (_m *Helper) CheckTransaction(orderID string) (string, error) {
	ret := _m.Called(orderID)

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (string, error)); ok {
		return rf(orderID)
	}
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(orderID)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(orderID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreatePaymentLink provides a mock function with given fields: orderID, totalPrice, items, customer
func (_m *Helper) CreatePaymentLink(orderID string, totalPrice int64, items []midtrans.ItemDetails, customer midtrans.CustomerDetails) (*snap.Response, error) {
	ret := _m.Called(orderID, totalPrice, items, customer)

	var r0 *snap.Response
	var r1 error
	if rf, ok := ret.Get(0).(func(string, int64, []midtrans.ItemDetails, midtrans.CustomerDetails) (*snap.Response, error)); ok {
		return rf(orderID, totalPrice, items, customer)
	}
	if rf, ok := ret.Get(0).(func(string, int64, []midtrans.ItemDetails, midtrans.CustomerDetails) *snap.Response); ok {
		r0 = rf(orderID, totalPrice, items, customer)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*snap.Response)
		}
	}

	if rf, ok := ret.Get(1).(func(string, int64, []midtrans.ItemDetails, midtrans.CustomerDetails) error); ok {
		r1 = rf(orderID, totalPrice, items, customer)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GenerateHash provides a mock function with given fields: password
func (_m *Helper) GenerateHash(password string) string {
	ret := _m.Called(password)

	var r0 string
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(password)
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// GenerateToken provides a mock function with given fields: id, role
func (_m *Helper) GenerateToken(id int, role string) *helpers.JSONWebToken {
	ret := _m.Called(id, role)

	var r0 *helpers.JSONWebToken
	if rf, ok := ret.Get(0).(func(int, string) *helpers.JSONWebToken); ok {
		r0 = rf(id, role)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*helpers.JSONWebToken)
		}
	}

	return r0
}

// GetPrediction provides a mock function with given fields: comment
func (_m *Helper) GetPrediction(comment string) string {
	ret := _m.Called(comment)

	var r0 string
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(comment)
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// UploadImage provides a mock function with given fields: folder, file
func (_m *Helper) UploadImage(folder string, file multipart.File) string {
	ret := _m.Called(folder, file)

	var r0 string
	if rf, ok := ret.Get(0).(func(string, multipart.File) string); ok {
		r0 = rf(folder, file)
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// VerifyHash provides a mock function with given fields: password, hashed
func (_m *Helper) VerifyHash(password string, hashed string) bool {
	ret := _m.Called(password, hashed)

	var r0 bool
	if rf, ok := ret.Get(0).(func(string, string) bool); ok {
		r0 = rf(password, hashed)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// NewHelper creates a new instance of Helper. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewHelper(t interface {
	mock.TestingT
	Cleanup(func())
}) *Helper {
	mock := &Helper{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
