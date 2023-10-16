package repository

import (
	"perpustakaan/features/placeholder"

	"github.com/labstack/gommon/log"
	"gorm.io/gorm"
)

type model struct {
	db *gorm.DB
}

func New(db *gorm.DB) placeholder.Repository {
	return &model {
		db: db,
	}
}

func (mdl *model) Paginate(page, size int) []placeholder.Placeholder {
	var placeholders []placeholder.Placeholder

	offset := (page - 1) * size

	result := mdl.db.Offset(offset).Limit(size).Find(&placeholders)
	
	if result.Error != nil {
		log.Error(result.Error)
		return nil
	}

	return placeholders
}

func (mdl *model) Insert(newPlaceholder placeholder.Placeholder) int64 {
	result := mdl.db.Create(&newPlaceholder)

	if result.Error != nil {
		log.Error(result.Error)
		return -1
	}

	return int64(newPlaceholder.ID)
}

func (mdl *model) SelectByID(placeholderID int) *placeholder.Placeholder {
	var placeholder placeholder.Placeholder
	result := mdl.db.First(&placeholder, placeholderID)

	if result.Error != nil {
		log.Error(result.Error)
		return nil
	}

	return &placeholder
}

func (mdl *model) Update(placeholder placeholder.Placeholder) int64 {
	result := mdl.db.Save(&placeholder)

	if result.Error != nil {
		log.Error(result.Error)
	}

	return result.RowsAffected
}

func (mdl *model) DeleteByID(placeholderID int) int64 {
	result := mdl.db.Delete(&placeholder.Placeholder{}, placeholderID)
	
	if result.Error != nil {
		log.Error(result.Error)
		return 0
	}

	return result.RowsAffected
}