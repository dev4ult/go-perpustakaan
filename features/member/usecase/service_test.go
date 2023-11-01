package usecase

import (
	"errors"
	"perpustakaan/features/member"
	"perpustakaan/features/member/dtos"
	"perpustakaan/features/member/mocks"
	helperMocks "perpustakaan/helpers/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindAll(t *testing.T) {
	var repository = mocks.NewRepository(t)
	var helper = helperMocks.NewHelper(t)
	var service = New(repository, helper)

	var members = []member.Member{
		{
			ID: 1,
			FullName: "Sarbin Sisarbin",
			CredentialNumber: "2107411026",
			Email: "sarbin@example.com",
			Password: "aspojdapowkdpoawwwdwdk", 
			PhoneNumber: "080000000000",
			Address: "Jalan Hj. Raisin",
		},
	}

	var page = 1
	var size = 10
	var email = ""
	var credNumber = ""

	t.Run("Success", func(t *testing.T) {
		repository.On("Paginate", page, size, email, credNumber).Return(members, nil).Once()

		result, message := service.FindAll(page, size, email, credNumber)
		assert.Equal(t, members[0].FullName, result[0].FullName)
		assert.Empty(t, message)
		repository.AssertExpectations(t)
	})

	t.Run("Failed", func(t *testing.T) {
		repository.On("Paginate", page, size, email, credNumber).Return(nil, errors.New("record not found")).Once()

		result, message := service.FindAll(page, size, email, credNumber)
		assert.Nil(t, result)
		assert.NotEmpty(t, message)
		repository.AssertExpectations(t)
	})
}

func TestFindByID(t *testing.T) {
	var repository = mocks.NewRepository(t)
	var helper = helperMocks.NewHelper(t)
	var service = New(repository, helper)

	var member = member.Member{
		ID: 1,
		FullName: "Sarbin Sisarbin",
		CredentialNumber: "2107411026",
		Email: "sarbin@example.com",
		Password: "aspojdapowkdpoawwwdwdk", 
		PhoneNumber: "080000000000",
		Address: "Jalan Hj. Raisin",
	}

	var memberID = 1

	t.Run("Success", func(t *testing.T) {
		repository.On("SelectByID", memberID).Return(&member, nil).Once()

		result, message := service.FindByID(memberID)
		assert.Equal(t, member.FullName, result.FullName)
		assert.Empty(t, message)
		repository.AssertExpectations(t)
	})

	t.Run("Failed", func(t *testing.T) {
		repository.On("SelectByID", 999).Return(nil, errors.New("record not found")).Once()

		result, message := service.FindByID(999)
		assert.Nil(t, result)
		assert.NotEmpty(t, message)
		repository.AssertExpectations(t)
	})
}

func TestCreate(t *testing.T) {
	var repository = mocks.NewRepository(t)
	var helper = helperMocks.NewHelper(t)
	var service = New(repository, helper)

	var validMember = member.Member{
		FullName: "Sarbin Sisarbin",
		CredentialNumber: "2107411026",
		Email: "sarbin@example.com",
		Password: "randomgeneratedhash", 
		PhoneNumber: "080000000000",
		Address: "Jalan Hj. Raisin",
	}

	var invalidMember = member.Member{}

	var input = dtos.InputMember{
		FullName: "Sarbin Sisarbin",
		CredentialNumber: "2107411026",
		Email: "sarbin@example.com",
		Password: "sarbin123", 
		PhoneNumber: "080000000000",
		Address: "Jalan Hj. Raisin",
	}

	var emptyInput = dtos.InputMember{}

	t.Run("Success", func(t *testing.T) {
		repository.On("SelectByEmail", input.Email).Return(nil, errors.New("record not found")).Once()
		repository.On("SelectByCredentialNumber", input.CredentialNumber).Return(nil, errors.New("record not found")).Once()
		helper.On("GenerateHash", input.Password).Return("randomgeneratedhash").Once()
		repository.On("Insert", validMember).Return(1, nil).Once()

		result, message := service.Create(input)
		assert.Equal(t, validMember.FullName, result.FullName)
		assert.Empty(t, message)
		repository.AssertExpectations(t)
	})

	t.Run("Failed : Email Has Already Registered", func(t *testing.T) {
		repository.On("SelectByEmail", input.Email).Return(&validMember, nil).Once()

		result, message := service.Create(input)
		assert.Nil(t, result)
		assert.NotEmpty(t, message)
		assert.Equal(t, "Email Has Already Registered!", message)
		repository.AssertExpectations(t)
	})

	t.Run("Failed : Credential Number Has Already Registered", func(t *testing.T) {
		repository.On("SelectByEmail", input.Email).Return(nil, errors.New("record not found")).Once()
		repository.On("SelectByCredentialNumber", input.CredentialNumber).Return(&validMember, nil).Once()

		result, message := service.Create(input)
		assert.Nil(t, result)
		assert.NotEmpty(t, message)
		assert.Equal(t, "Credential Number Has Already Registered!", message)
		repository.AssertExpectations(t)
	})

	t.Run("Failed : Error When Insert", func(t *testing.T) {
		invalidMember.Password = "randomgeneratedhash"
		repository.On("SelectByEmail", emptyInput.Email).Return(nil, errors.New("record not found")).Once()
		repository.On("SelectByCredentialNumber", emptyInput.CredentialNumber).Return(nil, errors.New("record not found")).Once()
		helper.On("GenerateHash", emptyInput.Password).Return("randomgeneratedhash").Once()
		repository.On("Insert", invalidMember).Return(0, errors.New("error when insert")).Once()

		result, message := service.Create(emptyInput)
		assert.Nil(t, result)
		assert.NotEmpty(t, message)
		repository.AssertExpectations(t)
	})

	t.Run("Failed : Error When Hashing Password", func(t *testing.T) {
		repository.On("SelectByEmail", emptyInput.Email).Return(nil, errors.New("record not found")).Once()
		repository.On("SelectByCredentialNumber", emptyInput.CredentialNumber).Return(nil, errors.New("record not found")).Once()
		helper.On("GenerateHash", emptyInput.Password).Return("").Once()

		result, message := service.Create(emptyInput)
		assert.Nil(t, result)
		assert.NotEmpty(t, message)
		repository.AssertExpectations(t)
	})
}

func TestModify(t *testing.T) {
	var repository = mocks.NewRepository(t)
	var helper = helperMocks.NewHelper(t)
	var service = New(repository, helper)

	var validMember = member.Member{
		FullName: "Sarbin Sisarbin",
		CredentialNumber: "2107411026",
		Email: "sarbin@example.com",
		Password: "randomgeneratedhash", 
		PhoneNumber: "080000000000",
		Address: "Jalan Hj. Raisin",
	}

	var invalidMember = member.Member{}

	var input = dtos.InputMember{
		FullName: "Sarbin Sisarbin",
		CredentialNumber: "2107411026",
		Email: "sarbin@example.com",
		Password: "sarbin123", 
		PhoneNumber: "080000000000",
		Address: "Jalan Hj. Raisin",
	}

	var emptyInput = dtos.InputMember{}

	var memberID = 1

	t.Run("Success", func(t *testing.T) {
		validMember.ID = memberID
		helper.On("GenerateHash", input.Password).Return("randomgeneratedhash").Once()
		repository.On("Update", validMember).Return(1, nil).Once()

		result, message := service.Modify(input, memberID)
		assert.Equal(t, true, result)
		assert.Empty(t, message)
		repository.AssertExpectations(t)
	})

	t.Run("Failed : Error When Update", func(t *testing.T) {
		invalidMember.Password = "randomgeneratedhash"
		helper.On("GenerateHash", emptyInput.Password).Return("randomgeneratedhash").Once()
		repository.On("Update", invalidMember).Return(0, errors.New("error when update")).Once()

		result, message := service.Modify(emptyInput, 0)
		assert.Equal(t, false, result)
		assert.NotEmpty(t, message)
		repository.AssertExpectations(t)
	})

	t.Run("Failed : Error When Hashing Password", func(t *testing.T) {
		helper.On("GenerateHash", emptyInput.Password).Return("").Once()

		result, message := service.Modify(emptyInput, 0)
		assert.Equal(t, false, result)
		assert.NotEmpty(t, message)
		repository.AssertExpectations(t)
	})
}

func TestRemove(t *testing.T) {
	var repository = mocks.NewRepository(t)
	var helper = helperMocks.NewHelper(t)
	var service = New(repository, helper)

	var memberID = 1

	t.Run("Success", func(t *testing.T) {
		repository.On("DeleteByID", memberID).Return(1, nil).Once()

		result, message := service.Remove(memberID)
		assert.Equal(t, true, result)
		assert.Empty(t, message)
		repository.AssertExpectations(t)
	})

	t.Run("Failed", func(t *testing.T) {
		repository.On("DeleteByID", 999).Return(0, errors.New("record not found")).Once()

		result, message := service.Remove(999)
		assert.Equal(t, false, result)
		assert.NotEmpty(t, message)
		repository.AssertExpectations(t)
	})
}