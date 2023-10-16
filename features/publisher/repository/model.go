package repository

import (
	"perpustakaan/features/publisher"

	"github.com/labstack/gommon/log"
	"gorm.io/gorm"
)

type model struct {
	db *gorm.DB
}

func New(db *gorm.DB) publisher.Repository {
	return &model {
		db: db,
	}
}

func (mdl *model) Paginate(page, size int) []publisher.Publisher {
	var publishers []publisher.Publisher

	offset := (page - 1) * size

	result := mdl.db.Offset(offset).Limit(size).Find(&publishers)
	
	if result.Error != nil {
		log.Error(result.Error)
		return nil
	}

	return publishers
}

func (mdl *model) Insert(newPublisher publisher.Publisher) int64 {
	result := mdl.db.Create(&newPublisher)

	if result.Error != nil {
		log.Error(result.Error)
		return -1
	}

	return int64(newPublisher.ID)
}

func (mdl *model) SelectByID(publisherID int) *publisher.Publisher {
	var publisher publisher.Publisher
	result := mdl.db.First(&publisher, publisherID)

	if result.Error != nil {
		log.Error(result.Error)
		return nil
	}

	return &publisher
}

func (mdl *model) Update(publisher publisher.Publisher) int64 {
	result := mdl.db.Save(&publisher)

	if result.Error != nil {
		log.Error(result.Error)
	}

	return result.RowsAffected
}

func (mdl *model) DeleteByID(publisherID int) int64 {
	result := mdl.db.Delete(&publisher.Publisher{}, publisherID)
	
	if result.Error != nil {
		log.Error(result.Error)
		return 0
	}

	return result.RowsAffected
}