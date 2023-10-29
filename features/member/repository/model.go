package repository

import (
	"errors"
	"perpustakaan/features/member"
	"perpustakaan/helpers"

	"gorm.io/gorm"
)

type model struct {
	db *gorm.DB
}

func New(db *gorm.DB) member.Repository {
	return &model {
		db: db,
	}
}

func (mdl *model) Paginate(page int, size int, email string, credentialNumber string) ([]member.Member, error) {
	var members []member.Member

	offset := (page - 1) * size
	emailKey := "%" + email + "%" 
	credNumberKey := "%" + credentialNumber + "%" 

	if err := mdl.db.Where("email LIKE ?", emailKey).Where("credential_number LIKE ?", credNumberKey).Offset(offset).Limit(size).Find(&members).Error; err != nil {
		return nil, err
	}

	return members, nil
}

func (mdl *model) Insert(newMember member.Member) (int, error) {
	hashPassword := helpers.GenerateHash(newMember.Password)
	if hashPassword == "" {
		return 0, errors.New("Error When Hashing Password")
	}

	newMember.Password = hashPassword
	
	if err := mdl.db.Create(&newMember).Error; err != nil {
		return 0, err
	}

	return newMember.ID, nil
}

func (mdl *model) SelectByID(memberID int) (*member.Member, error) {
	var member member.Member
	
	if err := mdl.db.First(&member, memberID).Error; err != nil {
		return nil, err
	}

	return &member, nil
}

func (mdl *model) SelectByEmail(email string) (*member.Member, error) {
	var member member.Member
	
	if err := mdl.db.Where("email = ?", email).First(&member).Error; err != nil {
		return nil, err
	}

	return &member, nil
}


func (mdl *model) SelectByCredentialNumber(credentialNumber string) (*member.Member, error) {
	var member member.Member
	
	if err := mdl.db.Where("credential_number = ?", credentialNumber).First(&member).Error; err != nil {
		return nil, err
	}

	return &member, nil
}

func (mdl *model) Update(member member.Member) (int, error) {
	hashPassword := helpers.GenerateHash(member.Password)
	if hashPassword == "" {
		return 0, errors.New("Error When Hashing Password")
	}

	member.Password = hashPassword

	result := mdl.db.Save(&member)
	if result.Error != nil {
		return 0, result.Error
	}

	return int(result.RowsAffected), nil
}

func (mdl *model) DeleteByID(memberID int) (int, error) {
	result := mdl.db.Delete(&member.Member{}, memberID)
	
	if result.Error != nil {
		return 0, result.Error
	}

	return int(result.RowsAffected), nil
}