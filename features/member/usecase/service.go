package usecase

import (
	"perpustakaan/features/member"
	"perpustakaan/features/member/dtos"

	"github.com/labstack/gommon/log"
	"github.com/mashingan/smapping"
)

type service struct {
	model member.Repository
}

func New(model member.Repository) member.Usecase {
	return &service {
		model: model,
	}
}

func (svc *service) FindAll(page, size int) []dtos.ResMember {
	var members []dtos.ResMember

	membersEnt := svc.model.Paginate(page, size)

	for _, member := range membersEnt {
		var data dtos.ResMember

		if err := smapping.FillStruct(&data, smapping.MapFields(member)); err != nil {
			log.Error(err.Error())
		} 
		
		members = append(members, data)
	}

	return members
}

func (svc *service) FindByID(memberID int) *dtos.ResMember {
	res := dtos.ResMember{}
	member := svc.model.SelectByID(memberID)

	if member == nil {
		return nil
	}

	err := smapping.FillStruct(&res, smapping.MapFields(member))
	if err != nil {
		log.Error(err)
		return nil
	}

	return &res
}

func (svc *service) Create(newMember dtos.InputMember) *dtos.ResMember {
	member := member.Member{}
	
	err := smapping.FillStruct(&member, smapping.MapFields(newMember))
	if err != nil {
		log.Error(err)
		return nil
	}

	memberID := svc.model.Insert(member)

	if memberID == -1 {
		return nil
	}

	resMember := dtos.ResMember{}
	errRes := smapping.FillStruct(&resMember, smapping.MapFields(newMember))
	if errRes != nil {
		log.Error(errRes)
		return nil
	}

	return &resMember
}

func (svc *service) Modify(memberData dtos.InputMember, memberID int) bool {
	newMember := member.Member{}

	err := smapping.FillStruct(&newMember, smapping.MapFields(memberData))
	if err != nil {
		log.Error(err)
		return false
	}

	newMember.ID = memberID
	rowsAffected := svc.model.Update(newMember)

	if rowsAffected <= 0 {
		log.Error("There is No Member Updated!")
		return false
	}
	
	return true
}

func (svc *service) Remove(memberID int) bool {
	rowsAffected := svc.model.DeleteByID(memberID)

	if rowsAffected <= 0 {
		log.Error("There is No Member Deleted!")
		return false
	}

	return true
}