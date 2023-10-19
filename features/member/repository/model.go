package repository

import (
	"perpustakaan/features/member"

	"github.com/labstack/gommon/log"
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

func (mdl *model) Paginate(page, size int) []member.Member {
	var members []member.Member

	offset := (page - 1) * size

	result := mdl.db.Offset(offset).Limit(size).Find(&members)
	
	if result.Error != nil {
		log.Error(result.Error)
		return nil
	}

	return members
}

func (mdl *model) Insert(newMember member.Member) int64 {
	result := mdl.db.Create(&newMember)

	if result.Error != nil {
		log.Error(result.Error)
		return -1
	}

	return int64(newMember.ID)
}

func (mdl *model) SelectByID(memberID int) *member.Member {
	var member member.Member
	result := mdl.db.First(&member, memberID)

	if result.Error != nil {
		log.Error(result.Error)
		return nil
	}

	return &member
}

func (mdl *model) Update(member member.Member) int64 {
	result := mdl.db.Save(&member)

	if result.Error != nil {
		log.Error(result.Error)
	}

	return result.RowsAffected
}

func (mdl *model) DeleteByID(memberID int) int64 {
	result := mdl.db.Delete(&member.Member{}, memberID)
	
	if result.Error != nil {
		log.Error(result.Error)
		return 0
	}

	return result.RowsAffected
}