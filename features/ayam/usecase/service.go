package usecase

import (
	"perpustakaan/features/ayam"
	"perpustakaan/features/ayam/dtos"

	"github.com/labstack/gommon/log"
	"github.com/mashingan/smapping"
)

type service struct {
	model ayam.Repository
}

func New(model ayam.Repository) ayam.Usecase {
	return &service {
		model: model,
	}
}

func (svc *service) FindAll(page, size int) []dtos.ResAyam {
	var ayams []dtos.ResAyam

	ayamsEnt := svc.model.Paginate(page, size)

	for _, ayam := range ayamsEnt {
		var data dtos.ResAyam

		if err := smapping.FillStruct(&data, smapping.MapFields(ayam)); err != nil {
			log.Error(err.Error())
		} 
		
		ayams = append(ayams, data)
	}

	return ayams
}

func (svc *service) FindByID(ayamID int) *dtos.ResAyam {
	res := dtos.ResAyam{}
	ayam := svc.model.SelectByID(ayamID)

	if ayam == nil {
		return nil
	}

	err := smapping.FillStruct(&res, smapping.MapFields(ayam))
	if err != nil {
		log.Error(err)
		return nil
	}

	return &res
}

func (svc *service) Create(newAyam dtos.InputAyam) *dtos.ResAyam {
	ayam := ayam.Ayam{}
	
	err := smapping.FillStruct(&ayam, smapping.MapFields(newAyam))
	if err != nil {
		log.Error(err)
		return nil
	}

	ayamID := svc.model.Insert(ayam)

	if ayamID == -1 {
		return nil
	}

	resAyam := dtos.ResAyam{}
	errRes := smapping.FillStruct(&resAyam, smapping.MapFields(newAyam))
	if errRes != nil {
		log.Error(errRes)
		return nil
	}

	return &resAyam
}

func (svc *service) Modify(ayamData dtos.InputAyam, ayamID int) bool {
	newAyam := ayam.Ayam{}

	err := smapping.FillStruct(&newAyam, smapping.MapFields(ayamData))
	if err != nil {
		log.Error(err)
		return false
	}

	newAyam.ID = ayamID
	rowsAffected := svc.model.Update(newAyam)

	if rowsAffected <= 0 {
		log.Error("There is No Ayam Updated!")
		return false
	}
	
	return true
}

func (svc *service) Remove(ayamID int) bool {
	rowsAffected := svc.model.DeleteByID(ayamID)

	if rowsAffected <= 0 {
		log.Error("There is No Ayam Deleted!")
		return false
	}

	return true
}