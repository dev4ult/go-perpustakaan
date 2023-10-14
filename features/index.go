package features

import (
	"perpustakaan/features/book"
	"perpustakaan/features/book/handler"
	"perpustakaan/features/book/repository"
	"perpustakaan/features/book/usecase"
	"perpustakaan/utils"
)

func BookHandler() book.Handler {
	db := utils.InitDB()
	repo := repository.New(db)
	uc := usecase.New(repo)
	return handler.New(uc)
}