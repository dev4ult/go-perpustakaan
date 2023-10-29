package author

import (
	"perpustakaan/features/author/dtos"

	"github.com/labstack/echo/v4"
)

type Repository interface {
	Paginate(page int, size int, searchKey string) ([]Author, error)
	Insert(newAuthor Author) (int, error)
	SelectByID(authorID int) (*Author, error)
	Update(author Author) (int, error)
	DeleteByID(authorID int) (int, error)
	IsAuthorshipExist(bookID, authorID int) (bool, error)
	SelectAuthorshipByID(authorshipID int) (*Authorship, error)
	InsertAuthorship(authorship dtos.InputAuthorshipIDS) (*dtos.BookAuthors, error)
	DeleteAuthorshipByID(authorshipID int) (int, error)
}

type Usecase interface {
	FindAll(page int, size int, searchKey string) ([]dtos.ResAuthor, string)
	FindByID(authorID int) (*dtos.ResAuthor, string)
	Create(newAuthor dtos.InputAuthor) (*dtos.ResAuthor, string)
	Modify(authorData dtos.InputAuthor, authorID int) (bool, string)
	Remove(authorID int) (bool, string)
	IsAuthorshipExistByID(authorshipID int) (bool, string)
	SetupAuthorship(anAuthorShipIDS dtos.InputAuthorshipIDS) (*dtos.BookAuthors, string)
	RemoveAuthorship(authorshipID int) (bool, string)
}

type Handler interface {
	GetAuthors() echo.HandlerFunc
	AuthorDetails() echo.HandlerFunc
	CreateAuthor() echo.HandlerFunc
	UpdateAuthor() echo.HandlerFunc
	DeleteAuthor() echo.HandlerFunc
	CreateAnAuthorship() echo.HandlerFunc
	DeleteAnAuthorship() echo.HandlerFunc
}
