package usecase

import (
	"perpustakaan/features/author"
	"perpustakaan/features/author/dtos"
	"time"

	"github.com/mashingan/smapping"
)

type service struct {
	model author.Repository
}

func New(model author.Repository) author.Usecase {
	return &service {
		model: model,
	}
}

func (svc *service) FindAll(page int, size int, searchKey string) ([]dtos.ResAuthor, string) {
	var res []dtos.ResAuthor

	authors, err := svc.model.Paginate(page, size, searchKey)

	if err != nil {
		return nil, err.Error()
	}

	for _, author := range authors {
		var data dtos.ResAuthor

		if err := smapping.FillStruct(&data, smapping.MapFields(author)); err != nil {
			return nil, err.Error()
		}

		parseDate, err := time.Parse(time.RFC3339, author.DOB)
		if err != nil {
			return nil, err.Error()
		}

		formatDate := parseDate.Format("02/01/2006")

		data.DOB = formatDate

		res = append(res, data)
	}

	return res, ""
}

func (svc *service) FindByID(authorID int) (*dtos.ResAuthor, string) {
	var res dtos.ResAuthor
	author, err := svc.model.SelectByID(authorID)

	if err != nil {
		return nil, err.Error()
	}
	
	if err := smapping.FillStruct(&res, smapping.MapFields(author)); err != nil {
		return nil, err.Error()
	}

	parseDate, err := time.Parse(time.RFC3339, author.DOB)
	if err != nil {
		return nil, err.Error()
	}

	formatDate := parseDate.Format("02/01/2006")

	res.DOB = formatDate

	return &res, ""
}

func (svc *service) Create(newAuthor dtos.InputAuthor) (*dtos.ResAuthor, string) {
	var author author.Author
	
	if err := smapping.FillStruct(&author, smapping.MapFields(newAuthor)); err != nil {
		return nil, err.Error()
	}

	_, err := svc.model.Insert(author)

	if err != nil {
		return nil, err.Error()
	}

	var res dtos.ResAuthor
	
	if err := smapping.FillStruct(&res, smapping.MapFields(newAuthor)); err != nil {
		return nil, err.Error()
	}

	return &res, ""
}

func (svc *service) Modify(authorData dtos.InputAuthor, authorID int) (bool, string) {
	var newAuthor author.Author
	
	if err := smapping.FillStruct(&newAuthor, smapping.MapFields(authorData)); err != nil {
		return false, err.Error()
	}

	newAuthor.ID = authorID
	_, err := svc.model.Update(newAuthor)

	if err != nil {
		return false, err.Error()
	}
	
	return true, ""
}

func (svc *service) Remove(authorID int) (bool, string) {
	_, err := svc.model.DeleteByID(authorID)

	if err != nil {
		return false, err.Error()
	}

	return true, ""
}

func (svc *service) IsAuthorshipExistByID(authorshipID int) (bool, string) {
	_, err := svc.model.SelectAuthorshipByID(authorshipID)

	if err != nil {
		return false, err.Error()
	}

	return true, ""
}

func (svc *service) SetupAuthorship(IDS dtos.InputAuthorshipIDS) (*dtos.BookAuthors, string) {
	if exist, _ := svc.model.IsAuthorshipExist(IDS.BookID, IDS.AuthorID); exist {
		return nil, "An Authorship is Already Exist!"
	}

	bookAuthors, err := svc.model.InsertAuthorship(IDS)

	if err != nil {
		return nil, err.Error()
	}

	return bookAuthors, ""
}

func (svc *service) RemoveAuthorship(authorshipID int) (bool, string) {
	_, err := svc.model.DeleteAuthorshipByID(authorshipID)

	if err != nil {
		return false, err.Error()
	}

	return true, ""
}