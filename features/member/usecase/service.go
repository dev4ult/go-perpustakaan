package usecase

import (
	"perpustakaan/features/member"
	"perpustakaan/features/member/dtos"
	"perpustakaan/helpers"

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

func (svc *service) FindAll(page int, size int, email string, credentialNumber string) ([]dtos.ResMember, string) {
	var res []dtos.ResMember

	members, err := svc.model.Paginate(page, size, email, credentialNumber)

	if err != nil {
		return nil, err.Error()
	}

	for _, member := range members {
		var data dtos.ResMember

		if err := smapping.FillStruct(&data, smapping.MapFields(member)); err != nil {
			return nil, err.Error()
		} 
		
		res = append(res, data)
	}

	return res, ""
}

func (svc *service) FindByID(memberID int) (*dtos.ResMember, string) {
	var res dtos.ResMember
	member, err := svc.model.SelectByID(memberID)

	if err != nil {
		return nil, err.Error()
	}
	
	if err := smapping.FillStruct(&res, smapping.MapFields(member)); err != nil {
		return nil, err.Error()
	}

	return &res, ""
}

func (svc *service) Create(newMember dtos.InputMember) (*dtos.ResMember, string) {
	var member member.Member
	
	if err := smapping.FillStruct(&member, smapping.MapFields(newMember)); err != nil {
		return nil, err.Error()
	}

	hashPassword := helpers.GenerateHash(member.Password)
	if hashPassword == "" {
		return nil, "Error When Hashing Password"
	}

	member.Password = hashPassword

	_, err := svc.model.Insert(member)

	if err != nil {
		return nil, err.Error()
	}

	var res dtos.ResMember
	
	if err := smapping.FillStruct(&res, smapping.MapFields(newMember)); err != nil {
		return nil, err.Error()
	}

	return &res, ""
}

func (svc *service) Modify(memberData dtos.InputMember, memberID int) (bool, string) {
	var newMember member.Member
	
	if err := smapping.FillStruct(&newMember, smapping.MapFields(memberData)); err != nil {
		return false, err.Error()
	}

	newMember.ID = memberID
	_, err := svc.model.Update(newMember)

	if err != nil {
		return false, err.Error()
	}
	
	return true, ""
}

func (svc *service) Remove(memberID int) (bool, string) {
	_, err := svc.model.DeleteByID(memberID)

	if err != nil {
		return false, err.Error()
	}

	return true, ""
}