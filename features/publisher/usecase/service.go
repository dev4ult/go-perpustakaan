package usecase

import (
	"perpustakaan/features/publisher"
	"perpustakaan/features/publisher/dtos"

	"github.com/mashingan/smapping"
)

type service struct {
	model publisher.Repository
}

func New(model publisher.Repository) publisher.Usecase {
	return &service {
		model: model,
	}
}

func (svc *service) FindAll(page int, size int, searchKey string) ([]dtos.ResPublisher, string) {
	var res []dtos.ResPublisher

	publishers, err := svc.model.Paginate(page, size, searchKey)

	if err != nil {
		return nil, err.Error()
	}

	for _, publisher := range publishers {
		var data dtos.ResPublisher

		if err := smapping.FillStruct(&data, smapping.MapFields(publisher)); err != nil {
			return nil, err.Error()
		} 
		
		res = append(res, data)
	}

	return res, ""
}

func (svc *service) FindByID(publisherID int) (*dtos.ResPublisher, string) {
	var res dtos.ResPublisher

	publisher, err := svc.model.SelectByID(publisherID)

	if err != nil {
		return nil, err.Error()
	}
	
	if err := smapping.FillStruct(&res, smapping.MapFields(publisher)); err != nil {
		return nil, err.Error()
	}

	return &res, ""
}

func (svc *service) Create(input dtos.InputPublisher) (*dtos.ResPublisher, string) {
	var publisher publisher.Publisher
	
	if err := smapping.FillStruct(&publisher, smapping.MapFields(input)); err != nil {
		return nil, err.Error()
	}

	_, err := svc.model.Insert(publisher)

	if err != nil {
		return nil, err.Error()
	}

	var res dtos.ResPublisher
	
	if err := smapping.FillStruct(&res, smapping.MapFields(publisher)); err != nil {
		return nil, err.Error()
	}

	return &res, ""
}

func (svc *service) Modify(publisherData dtos.InputPublisher, publisherID int) (bool, string) {
	var publisher publisher.Publisher
	
	if err := smapping.FillStruct(&publisher, smapping.MapFields(publisherData)); err != nil {
		return false, err.Error()
	}

	publisher.ID = publisherID
	_, err := svc.model.Update(publisher)

	if err != nil {
		return false, err.Error()
	}
	
	return true, ""
}

func (svc *service) Remove(publisherID int) (bool, string) {
	_, err := svc.model.DeleteByID(publisherID)

	if err != nil {
		return false, err.Error()
	}

	return true, ""
}