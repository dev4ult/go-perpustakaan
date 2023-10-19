package member

import (
	"perpustakaan/features/member/dtos"

	"github.com/labstack/echo/v4"
)

type Repository interface {
	Paginate(page, size int) []Member
	Insert(newMember Member) int64
	SelectByID(memberID int) *Member
	Update(member Member) int64
	DeleteByID(memberID int) int64
}

type Usecase interface {
	FindAll(page, size int) []dtos.ResMember
	FindByID(memberID int) *dtos.ResMember
	Create(newMember dtos.InputMember) *dtos.ResMember
	Modify(memberData dtos.InputMember, memberID int) bool
	Remove(memberID int) bool
}

type Handler interface {
	GetMembers() echo.HandlerFunc
	MemberDetails() echo.HandlerFunc
	CreateMember() echo.HandlerFunc
	UpdateMember() echo.HandlerFunc
	DeleteMember() echo.HandlerFunc
}
