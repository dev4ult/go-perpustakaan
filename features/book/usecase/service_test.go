package usecase

import (
	"errors"
	"mime/multipart"
	"perpustakaan/features/book"
	"perpustakaan/features/book/dtos"
	"perpustakaan/features/book/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindAll(t *testing.T) {
	var repository = mocks.NewRepository(t)
	var service = New(repository)

	var books = []dtos.ResBook{
		{
			Title: "Dark Gathering",        
			CoverImage: "https://res.cloudinary.com/dlmkeu9hg/image/upload/v1697793758/book-cover/a2rsx6qrvzg126yidvaj.jpg",      
			Summary: "Lorem Ipsum dolor sit amet.",         
			PublicationYear: 2023, 
			Quantity: 10,        
			Language: "English",    
			NumberOfPages: 200,  
			Category: "Fiction",  
			Publisher: "PT. Gremadia", 
		},
	}

	page := 1
	size := 10
	searchKey := ""

	t.Run("Success", func(t *testing.T) {
		repository.On("Paginate", page, size, searchKey).Return(books, nil).Once()

		result, message := service.FindAll(page, size, searchKey)
		assert.Empty(t, message)
		assert.Equal(t, len(books),len(result))
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

	var book = dtos.ResBook{
		Title: "Dark Gathering",        
		CoverImage: "https://res.cloudinary.com/dlmkeu9hg/image/upload/v1697793758/book-cover/a2rsx6qrvzg126yidvaj.jpg",      
		Summary: "Lorem Ipsum dolor sit amet.",         
		PublicationYear: 2023, 
		Quantity: 10,        
		Language: "English",    
		NumberOfPages: 200,  
		Category: "Fiction",  
		Publisher: "PT. Gremadia", 
	}

	var bookID = 1

	t.Run("Success", func(t *testing.T) {
		repository.On("SelectByID", bookID).Return(&book, nil).Once()

		result, message := service.FindByID(bookID)
		assert.Empty(t, message)
		assert.Equal(t, result.Title, book.Title)
		repository.AssertExpectations(t)
	})

	t.Run("Not Found", func(t *testing.T) {
		repository.On("SelectByID", bookID).Return(nil, errors.New("record not found")).Once()

		result, message := service.FindByID(bookID)
		assert.Nil(t, result)
		assert.NotEmpty(t, message)
		repository.AssertExpectations(t)
	})
}

func TestCreate(t *testing.T) {
	var repository = mocks.NewRepository(t)
	var service = New(repository)

	var validBook = book.Book{
		Title: "Dark Gathering",        
		CoverImage: "",      
		Summary: "Lorem Ipsum dolor sit amet.",         
		PublicationYear: 2023, 
		Quantity: 10,        
		Language: "English",    
		NumberOfPages: 200,  
		CategoryID: 1,  
		PublisherID: 1, 
	}

	var invalidBook = book.Book{}
	
	var inputBook = dtos.InputBook{
		Title: "Dark Gathering",        
		Summary: "Lorem Ipsum dolor sit amet.",         
		PublicationYear: 2023, 
		Quantity: 10,        
		Language: "English",    
		NumberOfPages: 200,  
		CategoryID: 1,  
		PublisherID: 1, 
	}
	
	var emptyInput = dtos.InputBook{}

	var bookCover multipart.File = nil

	t.Run("Success", func(t *testing.T) {
		repository.On("Insert", validBook).Return(1, nil).Once()

		result, message := service.Create(inputBook, bookCover)
		assert.Empty(t, message)
		assert.NotNil(t, result)
		assert.Equal(t, result.Title, validBook.Title)
		repository.AssertExpectations(t)
	})
	
	t.Run("Failed Mapping Request", func(t *testing.T) {
		repository.On("Insert", invalidBook).Return(0, errors.New("test")).Once()

		result, message := service.Create(emptyInput, bookCover)
		assert.NotEmpty(t, message)
		assert.Nil(t, result)
		repository.AssertExpectations(t)
	})
}

func TestModify(t *testing.T) {
	var repository = mocks.NewRepository(t)
	var service = New(repository)

	var validBook = book.Book{
		Title: "Dark Gathering",        
		CoverImage: "",      
		Summary: "Lorem Ipsum dolor sit amet.",         
		PublicationYear: 2023, 
		Quantity: 10,        
		Language: "English",    
		NumberOfPages: 200,  
		CategoryID: 1,  
		PublisherID: 1, 
	}

	var invalidBook = book.Book{}
	
	var inputBook = dtos.InputBook{
		Title: "Dark Gathering",        
		Summary: "Lorem Ipsum dolor sit amet.",         
		PublicationYear: 2023, 
		Quantity: 10,        
		Language: "English",    
		NumberOfPages: 200,  
		CategoryID: 1,  
		PublisherID: 1, 
	}

	var emptyInput = dtos.InputBook{}

	var bookID = 1
	t.Run("Success", func(t *testing.T) {
		validBook.ID = bookID
		repository.On("Update", validBook).Return(1, nil).Once()

		result, message := service.Modify(inputBook, bookID)
		assert.Empty(t, message)
		assert.Equal(t, true, result)
		repository.AssertExpectations(t)
	})

	t.Run("Failed", func(t *testing.T) {
		repository.On("Update", invalidBook).Return(0, errors.New("record not found")).Once()

		result, message := service.Modify(emptyInput, 0)
		assert.NotEmpty(t, message)
		assert.Equal(t, false, result)
		repository.AssertExpectations(t)
	})
}

func TestDelete(t *testing.T) {
	var repository = mocks.NewRepository(t)
	var service = New(repository)

	var bookID = 1
	t.Run("Success", func(t *testing.T) {
		repository.On("DeleteByID", bookID).Return(1, nil).Once()

		result, message := service.Remove(bookID)
		assert.Empty(t, message)
		assert.Equal(t, true, result)
		repository.AssertExpectations(t)
	})

	t.Run("Failed", func(t *testing.T) {
		repository.On("DeleteByID", 999).Return(0, errors.New("record not found")).Once()

		result, message := service.Remove(999)
		assert.NotEmpty(t, message)
		assert.Equal(t, false, result)
		repository.AssertExpectations(t)
	})
}