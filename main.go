package main

import (
	"fmt"
	"perpustakaan/config"
	"perpustakaan/helpers"
	m "perpustakaan/middlewares"
	"perpustakaan/routes"
	"perpustakaan/utils"

	"perpustakaan/features/auth"
	ah "perpustakaan/features/auth/handler"
	ar "perpustakaan/features/auth/repository"
	au "perpustakaan/features/auth/usecase"

	"perpustakaan/features/author"
	auh "perpustakaan/features/author/handler"
	aur "perpustakaan/features/author/repository"
	auu "perpustakaan/features/author/usecase"

	"perpustakaan/features/book"
	bh "perpustakaan/features/book/handler"
	br "perpustakaan/features/book/repository"
	bu "perpustakaan/features/book/usecase"

	"perpustakaan/features/feedback"
	fh "perpustakaan/features/feedback/handler"
	fr "perpustakaan/features/feedback/repository"
	fu "perpustakaan/features/feedback/usecase"

	"perpustakaan/features/member"
	mh "perpustakaan/features/member/handler"
	mr "perpustakaan/features/member/repository"
	mu "perpustakaan/features/member/usecase"

	"perpustakaan/features/publisher"
	ph "perpustakaan/features/publisher/handler"
	pr "perpustakaan/features/publisher/repository"
	pu "perpustakaan/features/publisher/usecase"

	"perpustakaan/features/loan_history"
	lh "perpustakaan/features/loan_history/handler"
	lr "perpustakaan/features/loan_history/repository"
	lu "perpustakaan/features/loan_history/usecase"

	"perpustakaan/features/transaction"
	th "perpustakaan/features/transaction/handler"
	tr "perpustakaan/features/transaction/repository"
	tu "perpustakaan/features/transaction/usecase"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)


func main() {
	cfg := config.LoadServerConfig()
	e := echo.New()

	e.Use(m.Logger())
	e.Use(middleware.CORS())
	e.Pre(middleware.RemoveTrailingSlash())

	routes.Auths(e, AuthHandler())
	routes.Books(e, BookHandler())
	routes.Publishers(e, PublisherHandler())
	routes.Members(e, MemberHandler())
	routes.Authors(e, AuthorHandler())
	routes.Feedbacks(e, FeedbackHandler())
	routes.LoanHistories(e, LoanHistoryHandler())
	routes.Transactions(e, TransactionHandler())
	
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", cfg.SERVER_PORT)))
}

var helper = helpers.New()

func AuthHandler() auth.Handler {
	db := utils.InitDB()
	repo := ar.New(db)
	uc := au.New(repo, helper)
	return ah.New(uc)
}

func BookHandler() book.Handler {
	db := utils.InitDB()
	repo := br.New(db)
	uc := bu.New(repo, helper)
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
	uc := mu.New(repo, helper)
	return mh.New(uc)
}

func AuthorHandler() author.Handler {
	db := utils.InitDB()
	repo := aur.New(db)
	uc := auu.New(repo)
	return auh.New(uc)
}

func FeedbackHandler() feedback.Handler {
	db := utils.InitDB()
	repo := fr.New(db)
	uc := fu.New(repo, helper)
	return fh.New(uc)
}

func LoanHistoryHandler() loan_history.Handler {
	db := utils.InitDB()
	repo := lr.New(db)
	uc := lu.New(repo)
	return lh.New(uc)
}

func TransactionHandler() transaction.Handler {
	db := utils.InitDB()
	repo := tr.New(db)
	uc := tu.New(repo, helper)
	return th.New(uc)
}