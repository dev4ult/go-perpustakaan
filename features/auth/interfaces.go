package auth

import (
	"perpustakaan/features/auth/dtos"
	"perpustakaan/features/member"

	"github.com/labstack/echo/v4"
)

type Repository interface {
	SelectLibrarianByStaffID(staffID string) *Librarian
	SelectMemberByCredentialNumber(credentialNumber string) *member.Member
}

type Usecase interface {
	CheckAuthorization(credential string, password string, isStaff bool) *dtos.ResLogin
}

type Handler interface {
	Login() echo.HandlerFunc
}
