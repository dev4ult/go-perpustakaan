package repository

import (
	"perpustakaan/features/publisher"

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

func (mdl *model) Paginate(page int, size int, searchKey string) ([]publisher.Publisher, error) {
	var publishers []publisher.Publisher

	offset := (page - 1) * size
	name := "%" + searchKey + "%"

	if err := mdl.db.Where("name LIKE ?", name).Offset(offset).Limit(size).Find(&publishers).Error; err != nil {
		return nil, err
	}

	return publishers, nil
}

func (mdl *model) Insert(newPublisher publisher.Publisher) (int, error) {
	if err := mdl.db.Create(&newPublisher).Error; err != nil {
		return 0, err
	}

	return newPublisher.ID, nil
}

func (mdl *model) SelectByID(publisherID int) (*publisher.Publisher, error) {
	var publisher publisher.Publisher

	if err := mdl.db.First(&publisher, publisherID).Error; err != nil {
		return nil, err
	}

	return &publisher, nil
}

func (mdl *model) Update(publisher publisher.Publisher) (int, error) {
	result := mdl.db.Save(&publisher)

	if result.Error != nil {
		return 0, result.Error
	}

	return int(result.RowsAffected), nil
}

func (mdl *model) DeleteByID(publisherID int) (int, error) {
	result := mdl.db.Delete(&publisher.Publisher{}, publisherID)
	
	if result.Error != nil {
		return 0, result.Error
	}

	return int(result.RowsAffected), nil
}