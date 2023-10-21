package usecase

import (
	"perpustakaan/features/author"
	"perpustakaan/features/author/dtos"
	"time"

	"github.com/labstack/gommon/log"
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

func (svc *service) FindAll(page, size int) []dtos.ResAuthor {
	var authors []dtos.ResAuthor

	authorsEnt := svc.model.Paginate(page, size)

	for _, author := range authorsEnt {
		var data dtos.ResAuthor

		if err := smapping.FillStruct(&data, smapping.MapFields(author)); err != nil {
			log.Error(err.Error())
		}

		parseDate, err := time.Parse("2006-01-02 15:04:05 -0700 MST", author.DOB)
		if err != nil {
			log.Error(err.Error())
		}

		formatDate := parseDate.Format("02/01/2006")

		data.DOB = formatDate

		authors = append(authors, data)
	}

	return authors
}

func (svc *service) FindByID(authorID int) *dtos.ResAuthor {
	res := dtos.ResAuthor{}
	author := svc.model.SelectByID(authorID)

	if author == nil {
		return nil
	}

	err := smapping.FillStruct(&res, smapping.MapFields(author))
	if err != nil {
		log.Error(err.Error())
		return nil
	}

	parseDate, err := time.Parse(time.RFC3339, author.DOB)
	if err != nil {
		log.Error(err.Error())
	}

	formatDate := parseDate.Format("02/01/2006")

	res.DOB = formatDate

	// formatDate := parseDate.Format("02/01/2006")
	// res.DOB = formatDate

	return &res
}

func (svc *service) Create(newAuthor dtos.InputAuthor) *dtos.ResAuthor {
	author := author.Author{}
	
	err := smapping.FillStruct(&author, smapping.MapFields(newAuthor))
	if err != nil {
		log.Error(err)
		return nil
	}

	authorID := svc.model.Insert(author)

	if authorID == -1 {
		return nil
	}

	resAuthor := dtos.ResAuthor{}
	errRes := smapping.FillStruct(&resAuthor, smapping.MapFields(newAuthor))
	if errRes != nil {
		log.Error(errRes)
		return nil
	}

	return &resAuthor
}

func (svc *service) Modify(authorData dtos.InputAuthor, authorID int) bool {
	newAuthor := author.Author{}

	err := smapping.FillStruct(&newAuthor, smapping.MapFields(authorData))
	if err != nil {
		log.Error(err)
		return false
	}

	newAuthor.ID = authorID
	rowsAffected := svc.model.Update(newAuthor)

	if rowsAffected <= 0 {
		log.Error("There is No Author Updated!")
		return false
	}
	
	return true
}

func (svc *service) Remove(authorID int) bool {
	rowsAffected := svc.model.DeleteByID(authorID)

	if rowsAffected <= 0 {
		log.Error("There is No Author Deleted!")
		return false
	}

	return true
}

func (svc *service) IsAuthorshipExistByID(authorshipID int) bool {
	authorship := svc.model.SelectAuthorshipByID(authorshipID)

	return authorship != nil
}

func (svc *service) SetupAuthorship(IDS dtos.InputAuthorshipIDS) (*dtos.BookAuthors, string) {

	if isExist := svc.model.IsAuthorshipExist(IDS.BookID, IDS.AuthorID); isExist {
		return nil, "An Authorship is Already Exist!"
	}

	bookAuthors, err := svc.model.InsertAuthorship(IDS)

	if err != nil {
		return nil, err.Error()
	}

	return bookAuthors, ""
}

func (svc *service) RemoveAuthorship(authorshipID int) bool {
	rowsAffected := svc.model.DeleteAuthorshipByID(authorshipID)

	if rowsAffected <= 0 {
		log.Error("There is No Author Deleted!")
		return false
	}

	return true
}