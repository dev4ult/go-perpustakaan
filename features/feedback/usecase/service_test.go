package usecase

import (
	"errors"
	"perpustakaan/features/feedback"
	"perpustakaan/features/feedback/dtos"
	"perpustakaan/features/feedback/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindAll(t *testing.T) {
	var repository = mocks.NewRepository(t)
	var service = New(repository)

	var feedbackJoinReply = []dtos.FeedbackJoinReply{
		{
			Member: "Sarbin Sisarbin",
			Comment: "Performa Lambat",	
			PriorityStatus: "high",
			Staff: "Nibras Alyassar",		
			Reply: "Terimakasih atas feedbacknya, kami akan usahakan untuk perbaiki",		
		},
	}

	var page = 1
	var size = 10
	var memberName = ""
	var priority = ""

	t.Run("Success", func(t *testing.T) {
		repository.On("Paginate", page, size, memberName, priority).Return(feedbackJoinReply, nil).Once()

		result, message := service.FindAll(page, size, memberName, priority)
		assert.Equal(t, feedbackJoinReply[0].Member, result[0].Member)
		assert.Empty(t, message)
		repository.AssertExpectations(t)
	})

	t.Run("Failed", func(t *testing.T) {
		repository.On("Paginate", page, size, memberName, priority).Return(nil, errors.New("record not found")).Once()

		result, message := service.FindAll(page, size, memberName, priority)
		assert.Nil(t, result)
		assert.NotEmpty(t, message)
		repository.AssertExpectations(t)
	})
}

func TestFindByID(t *testing.T) {
	var repository = mocks.NewRepository(t)
	var service = New(repository)

	var feedbackWithReply = dtos.FeedbackWithReply{
		Member: "Sarbin Sisarbin",
		Comment: "Performa Lambat",	
		PriorityStatus: "high",
		Reply: dtos.StaffReply{},		
	}

	var feedbackID = 1

	t.Run("Success", func(t *testing.T) {
		repository.On("SelectByID", feedbackID).Return(&feedbackWithReply, nil).Once()

		result, message := service.FindByID(feedbackID)
		assert.Equal(t, feedbackWithReply.Member, result.Member)
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

	var response = dtos.ResFeedback{
		Member: "Anonymous",
		Comment: "Performa Lambat",	
		PriorityStatus: "high",
	}

	var input = dtos.InputFeedback{
		Comment: "Performa Lambat",
	}

	var emptyInput = dtos.InputFeedback{}

	var validFeedback = feedback.Feedback{
		Comment: "Performa Lambat",
	}

	var invalidFeedback = feedback.Feedback{}

	t.Run("Success", func(t *testing.T) {
		repository.On("Insert", validFeedback).Return(&response, nil).Once()

		result, message := service.Create(input)
		assert.Equal(t, response.Member, result.Member)
		assert.Empty(t, message)
		repository.AssertExpectations(t)
	})

	t.Run("Failed", func(t *testing.T) {
		repository.On("Insert", invalidFeedback).Return(nil, errors.New("record not found")).Once()

		result, message := service.Create(emptyInput)
		assert.Nil(t, result)
		assert.NotEmpty(t, message)
		repository.AssertExpectations(t)
	})
}

func TestAddAReply(t *testing.T) {
	var repository = mocks.NewRepository(t)
	var service = New(repository)

	var reply = dtos.InputReply{
		StaffID: 1,
		Comment: "Terimakasih atas Feedbacknya",
	}

	var emptyReply = dtos.InputReply{}

	var feedbackReply = feedback.FeedbackReply{
		Comment: "Terimakasih atas Feedbacknya", 
		LibrarianID: 1,
		FeedbackID: 1,
	}

	var invalidFeedbackReply = feedback.FeedbackReply{}

	var staffReply = dtos.StaffReply{
		Staff: "Nibras",
		Comment: "Terimakasih atas Feedbacknya",
	}

	var feedbackID = 1

	t.Run("Success", func(t *testing.T) {
		repository.On("InsertReplyForAFeedback", feedbackReply).Return(&staffReply, nil).Once()

		result, message := service.AddAReply(reply, feedbackID)
		assert.Equal(t, staffReply.Staff, result.Staff)
		assert.Empty(t, message)
		repository.AssertExpectations(t)
	})

	t.Run("Failed", func(t *testing.T) {
		repository.On("InsertReplyForAFeedback", invalidFeedbackReply).Return(nil, errors.New("record not found")).Once()

		result, message := service.AddAReply(emptyReply, 0)
		assert.Nil(t, result)
		assert.NotEmpty(t, message)
		repository.AssertExpectations(t)
	})
}

func TestRemove(t *testing.T) {
	var repository = mocks.NewRepository(t)
	var service = New(repository)

	var feedbackID = 1

	t.Run("Success", func(t *testing.T) {
		repository.On("DeleteByID", feedbackID).Return(1, nil).Once()

		result, message := service.Remove(feedbackID)
		assert.Equal(t, true, result)
		assert.Empty(t, message)
		repository.AssertExpectations(t)
	})

	t.Run("Failed", func(t *testing.T) {
		repository.On("DeleteByID", 0).Return(0, errors.New("error when delete")).Once()

		result, message := service.Remove(0)
		assert.Equal(t, false, result)
		assert.NotEmpty(t, message)
		repository.AssertExpectations(t)
	})
}