package auth

import (
	"perpustakaan/features/auth/dtos"
	"perpustakaan/features/member"

	"github.com/labstack/echo/v4"
)

type Repository interface {
	SelectLibrarianByStaffID(staffID string) (*Librarian, error)
	SelectMemberByCredentialNumber(credentialNumber string) (*member.Member, error)
	InsertNewLibrarian(newLibrarian Librarian) (*Librarian, error)
}

type Usecase interface {
	VerifyLogin(credential string, password string, isStaff bool) (*dtos.ResAuthorization, string)
	FindLibrarianByStaffID(staffID string) (*dtos.ResLibrarian, string)
	RegisterAStaff(newLibrarian dtos.InputStaffRegistration) (*dtos.ResLibrarian, string)
	RefreshToken(accessToken, refreshToken string) (*dtos.Token, string)
}

type Handler interface {
	Login() echo.HandlerFunc
	StaffRegistration() echo.HandlerFunc
	Refresh() echo.HandlerFunc
}
