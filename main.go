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

	"github.com/labstack/echo/v4"
)

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

var (
	bookHandler = BookHandler()
	publisherHandler = PublisherHandler()
	authHandler = AuthHandler()
	memberHandler = MemberHandler()
)

func main() {
	cfg := config.LoadServerConfig()
	e := echo.New()

	routes.Auths(e, authHandler)
	routes.Books(e, bookHandler, cfg)
	routes.Publishers(e, publisherHandler)
	routes.Members(e, memberHandler)
	
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", cfg.SERVER_PORT)))
}