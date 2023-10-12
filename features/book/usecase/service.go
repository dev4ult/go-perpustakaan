package usecase

import (
	"perpustakaan/features/book"
	"perpustakaan/features/book/dtos"
)

type service struct {
	model book.Repository
}

func (svc *service) GetAll(page, size int) []dtos.ResBook {
	return nil
}

func (svc *service) GetDetail(bookId int) *dtos.ResBook {
	return nil
}

func (svc *service) Create(newBook dtos.InputBook) bool {
	return true
}

func (svc *service) Modify(book dtos.InputBook) bool {
	return true
}

func (svc *service) Remove(bookId int) bool {
	return true
}