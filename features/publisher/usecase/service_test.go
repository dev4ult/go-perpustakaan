package usecase

import (
	"errors"
	"perpustakaan/features/publisher"
	"perpustakaan/features/publisher/dtos"
	"perpustakaan/features/publisher/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindAll(t *testing.T) {
	var repository = mocks.NewRepository(t)
	var service = New(repository)

	var publishers = []publisher.Publisher{
		{
			ID: 3,
			Name: "PT. Gremadia",
			Country: "Indonesia",
			City: "Jakarta Barat",
			Address: "Jalan HJ. Raisun",
			PostalCode: 12330,
			PhoneNumber: "080000000000", 
			Email: "gremadia@example.com",
		},
	}

	// var res = []dtos.ResPublisher{
	// 	{
	// 		Name: "PT. Gremadia",
	// 		Country: "Indonesia",
	// 		City: "Jakarta Barat",
	// 		Address: "Jalan HJ. Raisun",
	// 		PostalCode: 12330,
	// 		PhoneNumber: "080000000000", 
	// 		Email: "gremadia@example.com",
	// 	},
	// }

	var page = 1
	var size = 10
	var searchKey = ""

	t.Run("Success", func(t *testing.T) {
		repository.On("Paginate", page, size, searchKey).Return(publishers, nil).Once()

		result, message := service.FindAll(page, size, searchKey)
		assert.Equal(t, result[0].Name, publishers[0].Name)
		assert.Empty(t, message)
		assert.NotZero(t, len(result))
		repository.AssertExpectations(t)
	})

	t.Run("Failed", func(t *testing.T) {
		repository.On("Paginate", page, size, searchKey).Return(nil, errors.New("record not found")).Once()

		result, message := service.FindAll(page, size, searchKey)
		assert.Nil(t, result)
		assert.NotEmpty(t, message)
		repository.AssertExpectations(t)
	})
}

func TestFindByID(t *testing.T) {
	var repository = mocks.NewRepository(t)
	var service = New(repository)

	var publisher = publisher.Publisher{
		ID: 3,
		Name: "PT. Gremadia",
		Country: "Indonesia",
		City: "Jakarta Barat",
		Address: "Jalan HJ. Raisun",
		PostalCode: 12330,
		PhoneNumber: "080000000000", 
		Email: "gremadia@example.com",
	}

	var publisherID = 3

	t.Run("Success", func(t *testing.T) {
		repository.On("SelectByID", publisherID).Return(&publisher, nil).Once()

		result, message := service.FindByID(publisherID)
		assert.NotNil(t, result)
		assert.Equal(t, result.Name, publisher.Name)
		assert.Empty(t, message)
		repository.AssertExpectations(t)
	})

	t.Run("Failed", func(t *testing.T) {
		repository.On("SelectByID", publisherID).Return(nil, errors.New("record not found")).Once()

		result, message := service.FindByID(publisherID)
		assert.Nil(t, result)
		assert.NotEmpty(t, message)
		repository.AssertExpectations(t)
	})
}

func TestCreate(t *testing.T) {
	var repository = mocks.NewRepository(t)
	var service = New(repository)

	var validPublisher = publisher.Publisher{
		Name: "PT. Gremadia",
		Country: "Indonesia",
		City: "Jakarta Barat",
		Address: "Jalan HJ. Raisun",
		PostalCode: 12330,
		PhoneNumber: "080000000000", 
		Email: "gremadia@example.com",
	}

	var invalidPublisher = publisher.Publisher{}

	var input = dtos.InputPublisher{
		Name: "PT. Gremadia",
		Country: "Indonesia",
		City: "Jakarta Barat",
		Address: "Jalan HJ. Raisun",
		PostalCode: 12330,
		PhoneNumber: "080000000000", 
		Email: "gremadia@example.com",
	}

	var emptyInput = dtos.InputPublisher{}

	t.Run("Success", func(t *testing.T) {
		repository.On("Insert", validPublisher).Return(3, nil).Once()

		result, message := service.Create(input)
		assert.NotNil(t, result)
		assert.Equal(t, result.Name, validPublisher.Name)
		assert.Empty(t, message)
		repository.AssertExpectations(t)
	})

	t.Run("Failed", func(t *testing.T) {
		repository.On("Insert", invalidPublisher).Return(0, errors.New("record not found")).Once()

		result, message := service.Create(emptyInput)
		assert.Nil(t, result)
		assert.NotEmpty(t, message)
		repository.AssertExpectations(t)
	})
}

func TestModify(t *testing.T) {
	var repository = mocks.NewRepository(t)
	var service = New(repository)

	var validPublisher = publisher.Publisher{
		Name: "PT. Gremadia",
		Country: "Indonesia",
		City: "Jakarta Barat",
		Address: "Jalan HJ. Raisun",
		PostalCode: 12330,
		PhoneNumber: "080000000000", 
		Email: "gremadia@example.com",
	}

	var invalidPublisher = publisher.Publisher{}

	var input = dtos.InputPublisher{
		Name: "PT. Gremadia",
		Country: "Indonesia",
		City: "Jakarta Barat",
		Address: "Jalan HJ. Raisun",
		PostalCode: 12330,
		PhoneNumber: "080000000000", 
		Email: "gremadia@example.com",
	}

	var emptyInput = dtos.InputPublisher{}

	var publisherID = 3

	t.Run("Success", func(t *testing.T) {
		validPublisher.ID = publisherID
		repository.On("Update", validPublisher).Return(1, nil).Once()

		result, message := service.Modify(input, publisherID)
		assert.NotNil(t, result)
		assert.Equal(t, true, result)
		assert.Empty(t, message)
		repository.AssertExpectations(t)
	})

	t.Run("Failed", func(t *testing.T) {
		repository.On("Update", invalidPublisher).Return(0, errors.New("record not found")).Once()

		result, message := service.Modify(emptyInput, 0)
		assert.Equal(t, false, result)
		assert.NotEmpty(t, message)
		repository.AssertExpectations(t)
	})
}

func TestRemove(t *testing.T) {
	var repository = mocks.NewRepository(t)
	var service = New(repository)

	var publisherID = 3
	t.Run("Success", func(t *testing.T) {
		repository.On("DeleteByID", publisherID).Return(1, nil).Once()

		result, message := service.Remove(publisherID)
		assert.NotNil(t, result)
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