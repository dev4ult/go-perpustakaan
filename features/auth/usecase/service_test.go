package usecase

import (
	"errors"
	"perpustakaan/features/auth"
	"perpustakaan/features/auth/dtos"
	"perpustakaan/features/auth/mocks"
	"perpustakaan/helpers"
	helperMocks "perpustakaan/helpers/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVerifyLogin(t *testing.T) {
	var repository = mocks.NewRepository(t)
	var helper = helperMocks.NewHelper(t)
	var service = New(repository, helper)

	var credential = "2107411026"
	var password = "nibras123"
	var isLibrarian = true
	var role = "librarian"

	var librarian = auth.Librarian{
		ID: 1, 
		FullName: "Nibras Alyassar",
		StaffID: "2107411026",
		NIK: 12309128740,
		PhoneNumber: "080000000000",
		Address: "dirumah",
		Email: "nibras@example.com", 
		Password: "randomgeneratedhash", 
	}

	var token = helpers.JSONWebToken{
		AccessToken: "random-access-token",
		RefreshToken: "random-refresh-token",
	}

	t.Run("Success", func(t *testing.T) {
		repository.On("SelectLibrarianByStaffID", credential).Return(&librarian, nil).Once()
		helper.On("VerifyHash", password, librarian.Password).Return(true, nil).Once()
		helper.On("GenerateToken", librarian.ID, role).Return(&token)
		
		result, message := service.VerifyLogin(credential, password, isLibrarian)
		assert.Empty(t, message)
		assert.Equal(t, librarian.FullName, result.FullName)
		assert.Equal(t, token.AccessToken, result.AccessToken)
		helper.AssertExpectations(t)
		repository.AssertExpectations(t)
	})

	t.Run("Failed : Password Does Not Match", func(t *testing.T) {
		repository.On("SelectLibrarianByStaffID", credential).Return(&librarian, nil).Once()
		helper.On("VerifyHash", password, librarian.Password).Return(false, nil).Once()
		
		result, message := service.VerifyLogin(credential, password, isLibrarian)
		assert.NotEmpty(t, message)
		assert.Nil(t, result)
		helper.AssertExpectations(t)
		repository.AssertExpectations(t)
	})

	t.Run("Failed : Librarian Not Found", func(t *testing.T) {
		repository.On("SelectLibrarianByStaffID", credential).Return(nil, errors.New("record not found")).Once()
		
		result, message := service.VerifyLogin(credential, password, isLibrarian)
		assert.NotEmpty(t, message)
		assert.Nil(t, result)
		repository.AssertExpectations(t)
	})

	t.Run("Failed : Member Not Found", func(t *testing.T) {
		repository.On("SelectMemberByCredentialNumber", credential).Return(nil, errors.New("record not found")).Once()
		
		result, message := service.VerifyLogin(credential, password, false)
		assert.NotEmpty(t, message)
		assert.Nil(t, result)
		repository.AssertExpectations(t)
	})
}

func TestFindLibrarianByStaffID(t *testing.T) {
	var repository = mocks.NewRepository(t)
	var helper = helperMocks.NewHelper(t)
	var service = New(repository, helper)

	var librarian = auth.Librarian{
		ID: 1, 
		FullName: "Nibras Alyassar",
		StaffID: "2107411026",
		NIK: 12309128740,
		PhoneNumber: "080000000000",
		Address: "dirumah",
		Email: "nibras@example.com", 
		Password: "randomgeneratedhash", 
	}

	var staffID = "2107411026"

	t.Run("Success", func(t *testing.T) {
		repository.On("SelectLibrarianByStaffID", staffID).Return(&librarian, nil).Once()
		
		result, message := service.FindLibrarianByStaffID(staffID)
		assert.Empty(t, message)
		assert.Equal(t, librarian.FullName, result.FullName)
		repository.AssertExpectations(t)
	})

	t.Run("Failed", func(t *testing.T) {
		repository.On("SelectLibrarianByStaffID", staffID).Return(nil, errors.New("record not found")).Once()
		
		result, message := service.FindLibrarianByStaffID(staffID)
		assert.Nil(t, result)
		assert.NotEmpty(t, message)
		repository.AssertExpectations(t)
	})
}

func TestRegisterAStaff(t *testing.T) {
	var repository = mocks.NewRepository(t)
	var helper = helperMocks.NewHelper(t)
	var service = New(repository, helper)

	var inputStaffForm = dtos.InputStaffRegistration{
		FullName: "Nibras Alyassar",
		StaffID: "2107411026",
		NIK: 12309128740,
		PhoneNumber: "080000000000",
		Address: "dirumah",
		Email: "nibras@example.com", 
		Password: "nibras123",
	}

	var librarian = auth.Librarian{
		FullName: "Nibras Alyassar",
		StaffID: "2107411026",
		NIK: 12309128740,
		PhoneNumber: "080000000000",
		Address: "dirumah",
		Email: "nibras@example.com", 
		Password: "randomgeneratedhash", 
	}

	t.Run("Success", func(t *testing.T) {
		helper.On("GenerateHash", inputStaffForm.Password).Return("randomgeneratedhash").Once()
		repository.On("InsertNewLibrarian", librarian).Return(&librarian, nil).Once()

		result, message := service.RegisterAStaff(inputStaffForm)
		assert.Empty(t, message)
		assert.Equal(t, librarian.FullName, result.FullName)
		helper.AssertExpectations(t)
		repository.AssertExpectations(t)
	})

	t.Run("Failed : Error When Insert", func(t *testing.T) {
		helper.On("GenerateHash", inputStaffForm.Password).Return("randomgeneratedhash").Once()
		repository.On("InsertNewLibrarian", librarian).Return(nil, errors.New("error when insert")).Once()

		result, message := service.RegisterAStaff(inputStaffForm)
		assert.Nil(t, result)
		assert.NotEmpty(t, message)
		helper.AssertExpectations(t)
		repository.AssertExpectations(t)
	})

	t.Run("Failed : Error When Hashing", func(t *testing.T) {
		helper.On("GenerateHash", inputStaffForm.Password).Return("").Once()

		result, message := service.RegisterAStaff(inputStaffForm)
		assert.Nil(t, result)
		assert.NotEmpty(t, message)
		helper.AssertExpectations(t)
		repository.AssertExpectations(t)
	})
}

func TestRefreshToken(t *testing.T) {
	var repository = mocks.NewRepository(t)
	var helper = helperMocks.NewHelper(t)
	var service = New(repository, helper)

	var newToken = helpers.JSONWebToken{
		AccessToken: "random-access-token",
		RefreshToken: "random-refresh-token",
	}

	var userID = 1
	var role = "librarian"

	t.Run("Success", func(t *testing.T) {
		helper.On("GenerateToken", userID, role).Return(&newToken).Once()
		
		result, message := service.RefreshToken(userID, role)
		assert.Empty(t, message)
		assert.Equal(t, newToken.AccessToken, result.AccessToken)
		helper.AssertExpectations(t)
	})

	t.Run("Failed", func(t *testing.T) {
		helper.On("GenerateToken", userID, role).Return(nil).Once()
		
		result, message := service.RefreshToken(userID, role)
		assert.Nil(t, result)
		assert.NotEmpty(t, message)
		helper.AssertExpectations(t)
	})

}