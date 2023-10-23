package usecase

import (
	"perpustakaan/features/auth"
	"perpustakaan/features/auth/dtos"
	"perpustakaan/helpers"

	"github.com/labstack/gommon/log"
	"github.com/mashingan/smapping"
)

type service struct {
	model auth.Repository
}

func New(model auth.Repository) auth.Usecase {
	return &service {
		model: model,
	}
}

func (svc *service) VerifyLogin(credential string, password string, isLibrarian bool) *dtos.ResAuthorization {
	var userTemp dtos.User
	var err error
	var role string

	if isLibrarian {
		user := svc.model.SelectLibrarianByStaffID(credential)

		if user == nil {
			return nil
		}

		err = smapping.FillStruct(&userTemp, smapping.MapFields(user))
		role = "librarian"
	} else {
		user := svc.model.SelectMemberByCredentialNumber(credential)

		if user == nil {
			return nil
		}
		
		err = smapping.FillStruct(&userTemp, smapping.MapFields(user))
		role = "member"
	}

	if err != nil {
		log.Error(err.Error())
		return nil
	}

	if userTemp == (dtos.User{}) {
		log.Error("User Not Found!")
		return nil
	}
	
	if matchPassword := helpers.VerifyHash(password, userTemp.Password); matchPassword {
		token := helpers.GenerateToken(userTemp.ID, role)
		response := dtos.ResAuthorization{
			FullName: userTemp.FullName,
			AccessToken: token.AccessToken,
			RefreshToken: token.RefreshToken,
		}

		return &response
	}

	log.Error("Password Does Not Match!")
	return nil
}

func (svc *service) FindLibrarianByStaffID(staffID string) *dtos.ResLibrarian {
	staff := svc.model.SelectLibrarianByStaffID(staffID)

	if staff == nil {
		log.Error("Librarian Not Found!")
		return nil
	}
	
	var resStaff dtos.ResLibrarian

	if err := smapping.FillStruct(&resStaff, smapping.MapFields(staff)); err != nil {
		log.Error(err.Error())
		return nil
	}

	return &resStaff
}

func (svc *service) RegisterAStaff(newLibrarianInput dtos.InputStaffRegistration) *dtos.ResLibrarian {
	var librarian auth.Librarian

	if err := smapping.FillStruct(&librarian, smapping.MapFields(newLibrarianInput)); err != nil {
		log.Error(err.Error())
		return nil
	}

	hashPassword := helpers.GenerateHash(librarian.Password)
	if hashPassword == "" {
		log.Error("Error when Hashing Password!")
		return nil
	}
	librarian.Password = hashPassword
	
	newLibrarian := svc.model.InsertNewLibrarian(librarian)

	if newLibrarian == nil {
		log.Error("New Librarian Not Created!")
		return nil
	}

	var resLibrarian dtos.ResLibrarian

	if err := smapping.FillStruct(&resLibrarian, smapping.MapFields(librarian)); err != nil {
		log.Error(err.Error())
		return nil
	}
	
	return &resLibrarian
}

func (svc *service) RefreshToken(accessToken, refreshToken string) *dtos.Token {
	claims := helpers.ExtractRefreshToken(refreshToken)
	
	if claims == nil {
		return nil
	}
	
	token := helpers.GenerateToken(claims["id"].(int), claims["role"].(string))

	if token == nil {
		log.Error("Token Error")
		return nil
	}
	
	return &dtos.Token{
		AccessToken: token.AccessToken,
		RefreshToken: token.RefreshToken,
	}
}

// func (svc *service) FindAll(page, size int) []dtos.ResAuth {
// 	var auths []dtos.ResAuth

// 	authsEnt := svc.model.Paginate(page, size)

// 	for _, auth := range authsEnt {
// 		var data dtos.ResAuth

// 		if err := smapping.FillStruct(&data, smapping.MapFields(auth)); err != nil {
// 			log.Error(err.Error())
// 		} 
		
// 		auths = append(auths, data)
// 	}

// 	return auths
// }

// func (svc *service) FindByID(authID int) *dtos.ResAuth {
// 	res := dtos.ResAuth{}
// 	auth := svc.model.SelectByID(authID)

// 	if auth == nil {
// 		return nil
// 	}

// 	err := smapping.FillStruct(&res, smapping.MapFields(auth))
// 	if err != nil {
// 		log.Error(err)
// 		return nil
// 	}

// 	return &res
// }

// func (svc *service) Create(newAuth dtos.InputAuth) *dtos.ResAuth {
// 	auth := auth.Auth{}
	
// 	err := smapping.FillStruct(&auth, smapping.MapFields(newAuth))
// 	if err != nil {
// 		log.Error(err)
// 		return nil
// 	}

// 	authID := svc.model.Insert(auth)

// 	if authID == -1 {
// 		return nil
// 	}

// 	resAuth := dtos.ResAuth{}
// 	errRes := smapping.FillStruct(&resAuth, smapping.MapFields(newAuth))
// 	if errRes != nil {
// 		log.Error(errRes)
// 		return nil
// 	}

// 	return &resAuth
// }

// func (svc *service) Modify(authData dtos.InputAuth, authID int) bool {
// 	newAuth := auth.Auth{}

// 	err := smapping.FillStruct(&newAuth, smapping.MapFields(authData))
// 	if err != nil {
// 		log.Error(err)
// 		return false
// 	}

// 	newAuth.ID = authID
// 	rowsAffected := svc.model.Update(newAuth)

// 	if rowsAffected <= 0 {
// 		log.Error("There is No Auth Updated!")
// 		return false
// 	}
	
// 	return true
// }

// func (svc *service) Remove(authID int) bool {
// 	rowsAffected := svc.model.DeleteByID(authID)

// 	if rowsAffected <= 0 {
// 		log.Error("There is No Auth Deleted!")
// 		return false
// 	}

// 	return true
// }