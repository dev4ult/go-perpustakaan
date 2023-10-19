package features

import (
	"perpustakaan/features/book"
	bh "perpustakaan/features/book/handler"
	br "perpustakaan/features/book/repository"
	bu "perpustakaan/features/book/usecase"

	"perpustakaan/features/member"
	mh "perpustakaan/features/member/handler"
	mr "perpustakaan/features/member/repository"
	mu "perpustakaan/features/member/usecase"
	"perpustakaan/utils"
)

func BookHandler() book.Handler {
	db := utils.InitDB()
	repo := br.New(db)
	uc := bu.New(repo)
	return bh.New(uc)
}

func MemberHandler() member.Handler {
	db := utils.InitDB()
	repo := mr.New(db)
	uc := mu.New(repo)
	return mh.New(uc)
}