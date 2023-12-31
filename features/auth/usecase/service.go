package usecase

import (
	"perpustakaan/features/auth"
	"perpustakaan/features/auth/dtos"
	"perpustakaan/helpers"

	"github.com/mashingan/smapping"
)

type service struct {
	model auth.Repository
	helper helpers.Helper
}

func New(model auth.Repository, helper helpers.Helper) auth.Usecase {
	return &service {
		model: model,
		helper: helper,
	}
}

func (svc *service) VerifyLogin(credential string, password string, isLibrarian bool) (*dtos.ResAuthorization, string) {
	var userTemp dtos.User
	var err error
	var role string

	if isLibrarian {
		user, err := svc.model.SelectLibrarianByStaffID(credential)

		if err != nil {
			return nil, err.Error()
		}

		err = smapping.FillStruct(&userTemp, smapping.MapFields(user))
		role = "librarian"
	} else {
		user, err := svc.model.SelectMemberByCredentialNumber(credential)

		if err != nil {
			return nil, err.Error()
		}
		
		err = smapping.FillStruct(&userTemp, smapping.MapFields(user))
		role = "member"
	}

	if err != nil {
		return nil, err.Error()
	}

	if userTemp == (dtos.User{}) {
		return nil, "User Not Found"
	}
	
	if matchPassword := svc.helper.VerifyHash(password, userTemp.Password); matchPassword {
		token := svc.helper.GenerateToken(userTemp.ID, role)
		response := dtos.ResAuthorization{
			FullName: userTemp.FullName,
			AccessToken: token.AccessToken,
			RefreshToken: token.RefreshToken,
		}

		return &response, ""
	}

	return nil, "Password Does Not Match"
}

func (svc *service) FindLibrarianByStaffID(staffID string) (*dtos.ResLibrarian, string) {
	staff, err := svc.model.SelectLibrarianByStaffID(staffID)

	if err != nil {
		return nil, err.Error()
	}
	
	var resStaff dtos.ResLibrarian

	if err := smapping.FillStruct(&resStaff, smapping.MapFields(staff)); err != nil {
		return nil, err.Error()
	}

	return &resStaff, ""
}

func (svc *service) RegisterAStaff(newLibrarianInput dtos.InputStaffRegistration) (*dtos.ResLibrarian, string) {
	var librarian auth.Librarian

	if err := smapping.FillStruct(&librarian, smapping.MapFields(newLibrarianInput)); err != nil {
		return nil, err.Error()
	}

	hashPassword := svc.helper.GenerateHash(librarian.Password)
	if hashPassword == "" {
		return nil, "Error When Hashing The Password!"
	}
	librarian.Password = hashPassword
	
	_, err := svc.model.InsertNewLibrarian(librarian)

	if err != nil {
		return nil, err.Error()
	}

	var resLibrarian dtos.ResLibrarian

	if err := smapping.FillStruct(&resLibrarian, smapping.MapFields(librarian)); err != nil {
		return nil, err.Error()
	}
	
	return &resLibrarian, ""
}

func (svc *service) RefreshToken(userID int, role string) (*dtos.Token, string) {
	token := svc.helper.GenerateToken(userID, role)

	if token == nil {
		return nil, "Error When Generating Token!"
	}
	
	return &dtos.Token{
		AccessToken: token.AccessToken,
		RefreshToken: token.RefreshToken,
	}, ""
}