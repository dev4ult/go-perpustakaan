package usecase

import (
	"perpustakaan/features/auth"
	"perpustakaan/features/auth/dtos"
	"perpustakaan/helpers"
)

type service struct {
	model auth.Repository
}

func New(model auth.Repository) auth.Usecase {
	return &service {
		model: model,
	}
}

func (svc *service) CheckAuthorization(credential string, password string, isStaff bool) *dtos.ResLogin {
	var response dtos.ResLogin


	if isStaff {
		user = svc.model.SelectLibrarianByStaffID(credential)
	} else {
		user = svc.model.SelectMemberByCredentialNumber(credential)
	}

	if user != nil && matchPassword := helpers.CompareHash(password, user.Password) ; matchPassword {
			
	}


	return nil
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