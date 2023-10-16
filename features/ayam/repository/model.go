package repository

import (
	"perpustakaan/features/ayam"

	"github.com/labstack/gommon/log"
	"gorm.io/gorm"
)

type model struct {
	db *gorm.DB
}

func New(db *gorm.DB) ayam.Repository {
	return &model {
		db: db,
	}
}

func (mdl *model) Paginate(page, size int) []ayam.Ayam {
	var ayams []ayam.Ayam

	offset := (page - 1) * size

	result := mdl.db.Offset(offset).Limit(size).Find(&ayams)
	
	if result.Error != nil {
		log.Error(result.Error)
		return nil
	}

	return ayams
}

func (mdl *model) Insert(newAyam ayam.Ayam) int64 {
	result := mdl.db.Create(&newAyam)

	if result.Error != nil {
		log.Error(result.Error)
		return -1
	}

	return int64(newAyam.ID)
}

func (mdl *model) SelectByID(ayamID int) *ayam.Ayam {
	var ayam ayam.Ayam
	result := mdl.db.First(&ayam, ayamID)

	if result.Error != nil {
		log.Error(result.Error)
		return nil
	}

	return &ayam
}

func (mdl *model) Update(ayam ayam.Ayam) int64 {
	result := mdl.db.Save(&ayam)

	if result.Error != nil {
		log.Error(result.Error)
	}

	return result.RowsAffected
}

func (mdl *model) DeleteByID(ayamID int) int64 {
	result := mdl.db.Delete(&ayam.Ayam{}, ayamID)
	
	if result.Error != nil {
		log.Error(result.Error)
		return 0
	}

	return result.RowsAffected
}