package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DatabaseConfig struct {
	DB_USER string
	DB_PASS string
	DB_HOST	string
	DB_PORT	string
	DB_NAME string
}

type ServerConfig struct {
	SERVER_PORT string
}

func InitDB() *gorm.DB {
	config := loadDatabaseConfig()
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", config.DB_USER, config.DB_PASS, config.DB_HOST, config.DB_PORT, config.DB_NAME)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	migrate(db)
	
	return db
}

func loadDatabaseConfig() DatabaseConfig {
	godotenv.Load(".env")

	return DatabaseConfig {
		DB_USER: os.Getenv("DB_USER"),
		DB_PASS: os.Getenv("DB_PASS"),
		DB_HOST: os.Getenv("DB_HOST"),
		DB_PORT: os.Getenv("DB_PORT"),
		DB_NAME: os.Getenv("DB_NAME"),
	}
}

func migrate(db *gorm.DB) {
	db.AutoMigrate()
}