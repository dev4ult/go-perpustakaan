package utils

import (
	"perpustakaan/config"
	"perpustakaan/features/auth"
	"perpustakaan/features/author"
	"perpustakaan/features/book"
	"perpustakaan/features/feedback"
	"perpustakaan/features/loan_history"
	"perpustakaan/features/member"
	"perpustakaan/features/publisher"
	"perpustakaan/features/transaction"

	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {
	config := config.LoadDBConfig()
	
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", config.DB_USER, config.DB_PASS, config.DB_HOST, config.DB_PORT, config.DB_NAME)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	migrate(db)

	return db
}

func migrate(db *gorm.DB) {
	db.AutoMigrate(publisher.Publisher{}, author.Author{})
	db.AutoMigrate(book.Category{}, book.Book{}, author.Authorship{})
	db.AutoMigrate(loan_history.LoanHistory{}, loan_history.FineType{})
	db.AutoMigrate(member.Member{})
	db.AutoMigrate(auth.Librarian{}, feedback.Feedback{}, feedback.FeedbackReply{})
	db.AutoMigrate(transaction.Transaction{})
}
