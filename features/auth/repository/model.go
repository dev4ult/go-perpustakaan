package repository

import (
	"perpustakaan/features/auth"
	"perpustakaan/features/member"

	"github.com/labstack/gommon/log"
	"gorm.io/gorm"
)

type model struct {
	db *gorm.DB
}

func New(db *gorm.DB) auth.Repository {
	return &model {
		db: db,
	}
}

func (mdl *model) SelectLibrarianByStaffID(staffID string) *auth.Librarian {
	var user auth.Librarian
	result := mdl.db.Where("staff_id = ?", staffID).First(&user)

	if result.Error != nil {
		log.Error(result.Error)
		return nil
	}

	return &user
}

func (mdl *model) SelectMemberByCredentialNumber(credentialNumber string) *member.Member {
	var user member.Member
	result := mdl.db.Where("credential_number = ?", credentialNumber).First(&user)

	if result.Error != nil {
		log.Error(result.Error)
		return nil
	}

	return &user
}

func (mdl *model) InsertNewLibrarian(newLibrarian auth.Librarian) *auth.Librarian {
	result := mdl.db.Create(&newLibrarian)

	if result.Error != nil {
		log.Error(result.Error.Error())
		return nil
	}

	return &newLibrarian
}

// func (mdl *model) Paginate(page, size int) []auth.Auth {
// 	var auths []auth.Auth

// 	offset := (page - 1) * size

// 	result := mdl.db.Offset(offset).Limit(size).Find(&auths)
	
// 	if result.Error != nil {
// 		log.Error(result.Error)
// 		return nil
// 	}

// 	return auths
// }

// func (mdl *model) Insert(newAuth auth.Auth) int64 {
// 	result := mdl.db.Create(&newAuth)

// 	if result.Error != nil {
// 		log.Error(result.Error)
// 		return -1
// 	}

// 	return int64(newAuth.ID)
// }

// func (mdl *model) SelectByID(authID int) *auth.Auth {
// 	var auth auth.Auth
// 	result := mdl.db.First(&auth, authID)

// 	if result.Error != nil {
// 		log.Error(result.Error)
// 		return nil
// 	}

// 	return &auth
// }

// func (mdl *model) Update(auth auth.Auth) int64 {
// 	result := mdl.db.Save(&auth)

// 	if result.Error != nil {
// 		log.Error(result.Error)
// 	}

// 	return result.RowsAffected
// }

// func (mdl *model) DeleteByID(authID int) int64 {
// 	result := mdl.db.Delete(&auth.Auth{}, authID)
	
// 	if result.Error != nil {
// 		log.Error(result.Error)
// 		return 0
// 	}

// 	return result.RowsAffected
// }