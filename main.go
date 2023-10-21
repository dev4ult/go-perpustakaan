package main

import (
	"fmt"
	"perpustakaan/config"
	"perpustakaan/routes"
	"perpustakaan/utils"

	"perpustakaan/features/auth"
	ah "perpustakaan/features/auth/handler"
	ar "perpustakaan/features/auth/repository"
	au "perpustakaan/features/auth/usecase"

	"perpustakaan/features/book"
	bh "perpustakaan/features/book/handler"
	br "perpustakaan/features/book/repository"
	bu "perpustakaan/features/book/usecase"

	"perpustakaan/features/publisher"
	ph "perpustakaan/features/publisher/handler"
	pr "perpustakaan/features/publisher/repository"
	pu "perpustakaan/features/publisher/usecase"

	"perpustakaan/features/member"
	mh "perpustakaan/features/member/handler"
	mr "perpustakaan/features/member/repository"
	mu "perpustakaan/features/member/usecase"

	"perpustakaan/features/author"
	auh "perpustakaan/features/author/handler"
	aur "perpustakaan/features/author/repository"
	auu "perpustakaan/features/author/usecase"

	"github.com/labstack/echo/v4"
)

var (
	bookHandler = BookHandler()
	publisherHandler = PublisherHandler()
	authHandler = AuthHandler()
	memberHandler = MemberHandler()
	authorHandler = AuthorHandler()
)

func main() {
	cfg := config.LoadServerConfig()
	e := echo.New()

	routes.Auths(e, authHandler, cfg)
	routes.Books(e, bookHandler, cfg)
	routes.Publishers(e, publisherHandler)
	routes.Members(e, memberHandler)
	routes.Authors(e, authorHandler)
	
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", cfg.SERVER_PORT)))
}

func AuthHandler() auth.Handler {
	db := utils.InitDB()
	repo := ar.New(db)
	uc := au.New(repo)
	return ah.New(uc)
}

func BookHandler() book.Handler {
	db := utils.InitDB()
	repo := br.New(db)
	uc := bu.New(repo)
	return bh.New(uc)
}

func PublisherHandler() publisher.Handler {
	db := utils.InitDB()
	repo := pr.New(db)
	uc := pu.New(repo)
	return ph.New(uc)
}

func MemberHandler() member.Handler {
	db := utils.InitDB()
	repo := mr.New(db)
	uc := mu.New(repo)
	return mh.New(uc)
}

func AuthorHandler() author.Handler {
	db := utils.InitDB()
	repo := aur.New(db)
	uc := auu.New(repo)
	return auh.New(uc)
}