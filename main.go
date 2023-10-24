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

	"github.com/labstack/echo/v4"
)


func main() {
	cfg := config.LoadServerConfig()
	e := echo.New()

	routes.Auths(e, AuthHandler(), cfg)
	// routes.Books(e, BookHandler(), cfg)
	// routes.Publishers(e, PublisherHandler())
	// routes.Members(e, MemberHandler())
	// routes.Authors(e, AuthorHandler())
	// routes.Feedbacks(e, FeedbackHandler())
	
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", cfg.SERVER_PORT)))
}

func AuthHandler() auth.Handler {
	db := utils.InitDB()
	repo := ar.New(db)
	uc := au.New(repo)
	return ah.New(uc)
}

// func BookHandler() book.Handler {
// 	db := utils.InitDB()
// 	repo := br.New(db)
// 	uc := bu.New(repo)
// 	return bh.New(uc)
// }

// func PublisherHandler() publisher.Handler {
// 	db := utils.InitDB()
// 	repo := pr.New(db)
// 	uc := pu.New(repo)
// 	return ph.New(uc)
// }

// func MemberHandler() member.Handler {
// 	db := utils.InitDB()
// 	repo := mr.New(db)
// 	uc := mu.New(repo)
// 	return mh.New(uc)
// }

// func AuthorHandler() author.Handler {
// 	db := utils.InitDB()
// 	repo := aur.New(db)
// 	uc := auu.New(repo)
// 	return auh.New(uc)
// }

// func FeedbackHandler() feedback.Handler {
// 	db := utils.InitDB()
// 	repo := fr.New(db)
// 	uc := fu.New(repo)
// 	return fh.New(uc)
// }