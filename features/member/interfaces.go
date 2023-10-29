package member

import (
	"perpustakaan/features/member/dtos"

	"github.com/labstack/echo/v4"
)

type Repository interface {
	Paginate(page int, size int, email string, credentialNumber string) ([]Member, error)
	Insert(newMember Member) (int, error)
	SelectByID(memberID int) (*Member, error)
	SelectByEmail(email string) (*Member, error)
	SelectByCredentialNumber(credentialNumber string) (*Member, error)
	Update(member Member) (int, error)
	DeleteByID(memberID int) (int, error)
}

type Usecase interface {
	FindAll(page int, size int, email string, credentialNumber string) ([]dtos.ResMember, string)
	FindByID(memberID int) (*dtos.ResMember, string)
	Create(newMember dtos.InputMember) (*dtos.ResMember, string)
	Modify(memberData dtos.InputMember, memberID int) (bool, string)
	Remove(memberID int) (bool, string)
}

type Handler interface {
	GetMembers() echo.HandlerFunc
	MemberDetails() echo.HandlerFunc
	CreateMember() echo.HandlerFunc
	UpdateMember() echo.HandlerFunc
	DeleteMember() echo.HandlerFunc
}
