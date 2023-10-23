package auth

import (
	"perpustakaan/features/auth/dtos"
	"perpustakaan/features/member"

	"github.com/labstack/echo/v4"
)

type Repository interface {
	SelectLibrarianByStaffID(staffID string) *Librarian
	SelectMemberByCredentialNumber(credentialNumber string) *member.Member
	InsertNewLibrarian(newLibrarian Librarian) *Librarian
}

type Usecase interface {
	VerifyLogin(credential string, password string, isStaff bool) *dtos.ResAuthorization
	FindLibrarianByStaffID(staffID string) *dtos.ResLibrarian
	RegisterAStaff(newLibrarian dtos.InputStaffRegistration) *dtos.ResLibrarian
	RefreshToken(accessToken, refreshToken string) *dtos.Token
}

type Handler interface {
	Login() echo.HandlerFunc
	StaffRegistration() echo.HandlerFunc
	Refresh() echo.HandlerFunc
}
