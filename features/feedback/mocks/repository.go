// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	feedback "perpustakaan/features/feedback"
	dtos "perpustakaan/features/feedback/dtos"

	mock "github.com/stretchr/testify/mock"
)

// Repository is an autogenerated mock type for the Repository type
type Repository struct {
	mock.Mock
}

// DeleteByID provides a mock function with given fields: feedbackID
func (_m *Repository) DeleteByID(feedbackID int) (int, error) {
	ret := _m.Called(feedbackID)

	var r0 int
	var r1 error
	if rf, ok := ret.Get(0).(func(int) (int, error)); ok {
		return rf(feedbackID)
	}
	if rf, ok := ret.Get(0).(func(int) int); ok {
		r0 = rf(feedbackID)
	} else {
		r0 = ret.Get(0).(int)
	}

	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(feedbackID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Insert provides a mock function with given fields: newFeedback
func (_m *Repository) Insert(newFeedback feedback.Feedback) (*dtos.ResFeedback, error) {
	ret := _m.Called(newFeedback)

	var r0 *dtos.ResFeedback
	var r1 error
	if rf, ok := ret.Get(0).(func(feedback.Feedback) (*dtos.ResFeedback, error)); ok {
		return rf(newFeedback)
	}
	if rf, ok := ret.Get(0).(func(feedback.Feedback) *dtos.ResFeedback); ok {
		r0 = rf(newFeedback)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*dtos.ResFeedback)
		}
	}

	if rf, ok := ret.Get(1).(func(feedback.Feedback) error); ok {
		r1 = rf(newFeedback)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// InsertReplyForAFeedback provides a mock function with given fields: reply
func (_m *Repository) InsertReplyForAFeedback(reply feedback.FeedbackReply) (*dtos.StaffReply, error) {
	ret := _m.Called(reply)

	var r0 *dtos.StaffReply
	var r1 error
	if rf, ok := ret.Get(0).(func(feedback.FeedbackReply) (*dtos.StaffReply, error)); ok {
		return rf(reply)
	}
	if rf, ok := ret.Get(0).(func(feedback.FeedbackReply) *dtos.StaffReply); ok {
		r0 = rf(reply)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*dtos.StaffReply)
		}
	}

	if rf, ok := ret.Get(1).(func(feedback.FeedbackReply) error); ok {
		r1 = rf(reply)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Paginate provides a mock function with given fields: page, size, member, priority
func (_m *Repository) Paginate(page int, size int, member string, priority string) ([]dtos.FeedbackJoinReply, error) {
	ret := _m.Called(page, size, member, priority)

	var r0 []dtos.FeedbackJoinReply
	var r1 error
	if rf, ok := ret.Get(0).(func(int, int, string, string) ([]dtos.FeedbackJoinReply, error)); ok {
		return rf(page, size, member, priority)
	}
	if rf, ok := ret.Get(0).(func(int, int, string, string) []dtos.FeedbackJoinReply); ok {
		r0 = rf(page, size, member, priority)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]dtos.FeedbackJoinReply)
		}
	}

	if rf, ok := ret.Get(1).(func(int, int, string, string) error); ok {
		r1 = rf(page, size, member, priority)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SelectByID provides a mock function with given fields: feedbackID
func (_m *Repository) SelectByID(feedbackID int) (*dtos.FeedbackWithReply, error) {
	ret := _m.Called(feedbackID)

	var r0 *dtos.FeedbackWithReply
	var r1 error
	if rf, ok := ret.Get(0).(func(int) (*dtos.FeedbackWithReply, error)); ok {
		return rf(feedbackID)
	}
	if rf, ok := ret.Get(0).(func(int) *dtos.FeedbackWithReply); ok {
		r0 = rf(feedbackID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*dtos.FeedbackWithReply)
		}
	}

	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(feedbackID)
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