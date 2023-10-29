package usecase

import (
	"errors"
	"perpustakaan/features/author"
	"perpustakaan/features/author/dtos"
	"perpustakaan/features/author/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindAll(t *testing.T) {
	var repository = mocks.NewRepository(t)
	var service = New(repository)

	var authors = []author.Author{
		{
			ID: 1,
			FullName: "Sirayuki", 
			DOB: "1965-04-03T00:00:00+07:00", 
			Biography: "Lorem Ipsum", 
		},
	}

	var page = 1 
	var size = 10
	var searchKey = ""

	t.Run("Success", func(t *testing.T) {
		repository.On("Paginate", page, size, searchKey).Return(authors, nil).Once()

		result, message := service.FindAll(page, size, searchKey)
		assert.Equal(t, 1, len(result))
		assert.Equal(t, authors[0].FullName, result[0].FullName)
		assert.Empty(t, message)
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

	var author = author.Author{
		ID: 1,
		FullName: "Sirayuki", 
		DOB: "1965-04-03T00:00:00+07:00", 
		Biography: "Lorem Ipsum", 
	}

	var authorID = 1

	t.Run("Success", func(t *testing.T) {
		repository.On("SelectByID", authorID).Return(&author, nil).Once()

		result, message := service.FindByID(authorID)
		assert.Equal(t, author.FullName, result.FullName)
		assert.Empty(t, message)
		repository.AssertExpectations(t)
	})

	t.Run("Failed", func(t *testing.T) {
		repository.On("SelectByID", authorID).Return(nil, errors.New("record not found")).Once()

		result, message := service.FindByID(authorID)
		assert.Nil(t, result)
		assert.NotEmpty(t, message)
		repository.AssertExpectations(t)
	})
}

func TestCreat(t *testing.T) {
	var repository = mocks.NewRepository(t)
	var service = New(repository)

	var validAuthor = author.Author{
		FullName: "Sirayuki", 
		DOB: "1965-04-03T00:00:00+07:00", 
		Biography: "Lorem Ipsum", 
	}

	var invalidAuthor = author.Author{}

	var input = dtos.InputAuthor{
		FullName: "Sirayuki", 
		DOB: "1965-04-03T00:00:00+07:00", 
		Biography: "Lorem Ipsum", 
	}

	var emptyInput = dtos.InputAuthor{}

	t.Run("Success", func(t *testing.T) {
		repository.On("Insert", validAuthor).Return(1, nil).Once()

		result, message := service.Create(input)
		assert.Equal(t, validAuthor.FullName, result.FullName)
		assert.Empty(t, message)
		repository.AssertExpectations(t)
	})

	t.Run("Failed", func(t *testing.T) {
		repository.On("Insert", invalidAuthor).Return(0, errors.New("record not found")).Once()

		result, message := service.Create(emptyInput)
		assert.Nil(t, result)
		assert.NotEmpty(t, message)
		repository.AssertExpectations(t)
	})
}

func TestModify(t *testing.T) {
	var repository = mocks.NewRepository(t)
	var service = New(repository)

	var validAuthor = author.Author{
		FullName: "Sirayuki", 
		DOB: "1965-04-03T00:00:00+07:00", 
		Biography: "Lorem Ipsum", 
	}

	var invalidAuthor = author.Author{}

	var input = dtos.InputAuthor{
		FullName: "Sirayuki", 
		DOB: "1965-04-03T00:00:00+07:00", 
		Biography: "Lorem Ipsum", 
	}

	var emptyInput = dtos.InputAuthor{}

	var authorID = 1

	t.Run("Success", func(t *testing.T) {
		validAuthor.ID = authorID
		repository.On("Update", validAuthor).Return(1, nil).Once()

		result, message := service.Modify(input, authorID)
		assert.Equal(t, true, result)
		assert.Empty(t, message)
		repository.AssertExpectations(t)
	})

	t.Run("Failed", func(t *testing.T) {
		repository.On("Update", invalidAuthor).Return(0, errors.New("record not found")).Once()

		result, message := service.Modify(emptyInput, 0)
		assert.Equal(t, false, result)
		assert.NotEmpty(t, message)
		repository.AssertExpectations(t)
	})
}

func TestRemove(t *testing.T) {
	var repository = mocks.NewRepository(t)
	var service = New(repository)

	var authorID = 1

	t.Run("Success", func(t *testing.T) {
		repository.On("DeleteByID", authorID).Return(1, nil).Once()

		result, message := service.Remove(authorID)
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

func TestIsAuthorshipExistByID(t *testing.T) {
	var repository = mocks.NewRepository(t)
	var service = New(repository)

	var authorship = author.Authorship{
		ID: 1,
		BookID: 2,  
		AuthorID: 3, 
	}

	var authorshipID = 1

	t.Run("Success", func(t *testing.T) {
		repository.On("SelectAuthorshipByID", authorshipID).Return(&authorship, nil).Once()

		result, message := service.IsAuthorshipExistByID(authorshipID)
		assert.Equal(t, true, result)
		assert.Empty(t, message)
		repository.AssertExpectations(t)
	})

	t.Run("Failed", func(t *testing.T) {
		repository.On("SelectAuthorshipByID", 0).Return(nil, errors.New("record not found")).Once()

		result, message := service.IsAuthorshipExistByID(0)
		assert.Equal(t, false, result)
		assert.NotEmpty(t, message)
		repository.AssertExpectations(t)
	})
}

func TestSetupAuthorship(t *testing.T) {
	var repository = mocks.NewRepository(t)
	var service = New(repository)

	var input = dtos.InputAuthorshipIDS{
		BookID: 1,
		AuthorID: 2,
	}
	
	var invalidInput = dtos.InputAuthorshipIDS{
		BookID: 2,
		AuthorID: 3,
	}

	var bookAuthors = dtos.BookAuthors{
		Title: "Dark Gathering",
		CoverImage: "",
		Summary: "lorem ipsum",
		PublicationYear: 2023,
		Quantity: 11,
		Language : "English",
		NumberOfPages: 200,

		Authors: []dtos.ResAuthor{
			{
				FullName: "Sirayuki", 
				DOB: "1965-04-03T00:00:00+07:00", 
				Biography: "Lorem Ipsum", 
			},
		},
	}

	var bookID = 2
	var authorID = 3

	t.Run("Success", func(t *testing.T) {
		repository.On("IsAuthorshipExist", 1, 2).Return(false, errors.New("record not found")).Once()
		repository.On("InsertAuthorship", input).Return(&bookAuthors, nil).Once()

		result, message := service.SetupAuthorship(input)
		assert.NotNil(t, result)
		assert.Equal(t, len(bookAuthors.Authors), len(result.Authors))
		assert.Equal(t, bookAuthors.Title, result.Title)
		assert.Empty(t, message)
		repository.AssertExpectations(t)
	})

	t.Run("Failed : An Authorship Is Already Exist", func(t *testing.T) {
		repository.On("IsAuthorshipExist", bookID, authorID).Return(true, nil).Once()

		result, message := service.SetupAuthorship(invalidInput)
		assert.Nil(t, result)
		assert.NotEmpty(t, message)
		assert.Equal(t, "An Authorship is Already Exist!", message)
		repository.AssertExpectations(t)
	})

	t.Run("Failed : Error When Insert An Authorship", func(t *testing.T) {
		repository.On("IsAuthorshipExist", 0, 3).Return(false, errors.New("record not found")).Once()
		repository.On("InsertAuthorship", dtos.InputAuthorshipIDS{BookID: 0, AuthorID: 3}).Return(nil, errors.New("error when insert")).Once()
	
		result, message := service.SetupAuthorship(dtos.InputAuthorshipIDS{BookID: 0, AuthorID: 3})
		assert.Nil(t, result)
		assert.NotEmpty(t, message)
		repository.AssertExpectations(t)
	})
}

func TestRemoveAuthorship(t *testing.T) {
	var repository = mocks.NewRepository(t)
	var service = New(repository)

	var authorshipID = 1

	t.Run("Success", func(t *testing.T) {
		repository.On("DeleteAuthorshipByID", authorshipID).Return(1, nil).Once()

		result, message := service.RemoveAuthorship(authorshipID)
		assert.Equal(t, true, result)
		assert.Empty(t, message)
		repository.AssertExpectations(t)
	})

	t.Run("Failed", func(t *testing.T) {
		repository.On("DeleteAuthorshipByID", 0).Return(0, errors.New("record not found")).Once()

		result, message := service.RemoveAuthorship(0)
		assert.Equal(t, false, result)
		assert.NotEmpty(t, message)
		repository.AssertExpectations(t)
	})
}