package author

import (
	"perpustakaan/features/author/dtos"

	"github.com/labstack/echo/v4"
)

type Repository interface {
	Paginate(page, size int) []Author
	Insert(newAuthor Author) int64
	SelectByID(authorID int) *Author
	Update(author Author) int64
	DeleteByID(authorID int) int64
	IsAuthorshipExist(bookID, authorID int) bool
	SelectAuthorshipByID(authorshipID int) *Authorship
	InsertAuthorship(authorship dtos.InputAuthorshipIDS) (*dtos.BookAuthors, error)
	DeleteAuthorshipByID(authorshipID int) int64
}

type Usecase interface {
	FindAll(page, size int) []dtos.ResAuthor
	FindByID(authorID int) *dtos.ResAuthor
	Create(newAuthor dtos.InputAuthor) *dtos.ResAuthor
	Modify(authorData dtos.InputAuthor, authorID int) bool
	Remove(authorID int) bool
	IsAuthorshipExistByID(authorshipID int) bool
	SetupAuthorship(anAuthorShipIDS dtos.InputAuthorshipIDS) (*dtos.BookAuthors, string)
	RemoveAuthorship(authorshipID int) bool
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
