package usecase

import (
	"perpustakaan/features/publisher"
	"perpustakaan/features/publisher/dtos"

	"github.com/labstack/gommon/log"
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

func (svc *service) FindAll(page, size int) []dtos.ResPublisher {
	var publishers []dtos.ResPublisher

	publishersEnt := svc.model.Paginate(page, size)

	for _, publisher := range publishersEnt {
		var data dtos.ResPublisher

		if err := smapping.FillStruct(&data, smapping.MapFields(publisher)); err != nil {
			log.Error(err.Error())
		} 
		
		publishers = append(publishers, data)
	}

	return publishers
}

func (svc *service) FindByID(publisherID int) *dtos.ResPublisher {
	res := dtos.ResPublisher{}
	publisher := svc.model.SelectByID(publisherID)

	if publisher == nil {
		return nil
	}

	err := smapping.FillStruct(&res, smapping.MapFields(publisher))
	if err != nil {
		log.Error(err)
		return nil
	}

	return &res
}

func (svc *service) Create(newPublisher dtos.InputPublisher) *dtos.ResPublisher {
	publisher := publisher.Publisher{}
	
	err := smapping.FillStruct(&publisher, smapping.MapFields(newPublisher))
	if err != nil {
		log.Error(err)
		return nil
	}

	publisherID := svc.model.Insert(publisher)

	if publisherID == -1 {
		return nil
	}

	resPublisher := dtos.ResPublisher{}
	errRes := smapping.FillStruct(&resPublisher, smapping.MapFields(newPublisher))
	if errRes != nil {
		log.Error(errRes)
		return nil
	}

	return &resPublisher
}

func (svc *service) Modify(publisherData dtos.InputPublisher, publisherID int) bool {
	newPublisher := publisher.Publisher{}

	err := smapping.FillStruct(&newPublisher, smapping.MapFields(publisherData))
	if err != nil {
		log.Error(err)
		return false
	}

	newPublisher.ID = publisherID
	rowsAffected := svc.model.Update(newPublisher)

	if rowsAffected <= 0 {
		log.Error("There is No Publisher Updated!")
		return false
	}
	
	return true
}

func (svc *service) Remove(publisherID int) bool {
	rowsAffected := svc.model.DeleteByID(publisherID)

	if rowsAffected <= 0 {
		log.Error("There is No Publisher Deleted!")
		return false
	}

	return true
}