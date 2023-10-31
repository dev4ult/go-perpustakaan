package repository

import (
	"perpustakaan/features/auth"
	"perpustakaan/features/member"

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

func (mdl *model) SelectLibrarianByStaffID(staffID string) (*auth.Librarian, error) {
	var user auth.Librarian

	if err := mdl.db.Where("staff_id = ?", staffID).First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (mdl *model) SelectMemberByCredentialNumber(credentialNumber string) (*member.Member, error) {
	var user member.Member

	if err := mdl.db.Where("credential_number = ?", credentialNumber).First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (mdl *model) InsertNewLibrarian(newLibrarian auth.Librarian) (*auth.Librarian, error) {
	if err := mdl.db.Create(&newLibrarian).Error; err != nil {
		return nil, err
	}

	return &newLibrarian, nil
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