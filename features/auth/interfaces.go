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
	VerifyLogin(credential string, password string, isStaff bool) *dtos.ResAuthorization
	RefreshToken(accessToken, refreshToken string) *dtos.Token
}

type Handler interface {
	Login() echo.HandlerFunc
	Refresh() echo.HandlerFunc
}
